package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"

	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"google.golang.org/grpc"
)

// RaftServer implementation for PubSub
type RaftServer struct {
	SingleMachineServer
	id                string // addr:port
	peers             *sync.Map
	peerNum           uint64
	role              pb.Role
	term              uint64
	leaderID          string
	heartbeatInterval time.Duration
	ticker            *time.Ticker
	timer             *time.Timer
	log               []*pb.LogEntry
	voteFor           string
	tickerChange      chan time.Duration
}

// Peer structure in Raft implementation.
type Peer struct {
	client   pb.PubSubClient
	sideCar  pb.RaftSidecarClient
	logIndex int
}

// NewRaftServer creates new raft server in raft mode.
func NewRaftServer(cfg *pb.ServerConfig) *RaftServer {
	// Set the id
	var id string
	if cfg.RaftConfig.Address != "" {
		id = fmt.Sprintf("%s:%d", cfg.RaftConfig.Address, cfg.Port)
	} else {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		id = fmt.Sprintf("%s:%d", conn.LocalAddr().(*net.UDPAddr).IP.String(), cfg.Port)
	}
	log.Printf("Server id: %s", id)

	// Parse peers
	peers := &sync.Map{}
	timeout, _ := time.ParseDuration("10m")
	peerStrs := strings.Split(cfg.RaftConfig.Peers, ";")
	for _, peer := range peerStrs {
		log.Printf("connect to peer: %s", peer)
		conn, err := grpc.Dial(
			peer,
			grpc.WithTimeout(timeout),
			grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		peers.Store(
			peer,
			&Peer{
				client:   pb.NewPubSubClient(conn),
				sideCar:  pb.NewRaftSidecarClient(conn),
				logIndex: 1,
			})
	}

	// Parse heart beat interval
	heartbeatInterval, _ := time.ParseDuration(cfg.RaftConfig.HeartbeatInterval)

	svr := &RaftServer{
		SingleMachineServer: SingleMachineServer{&sync.Map{}},
		id:                  id,
		peers:               peers,
		peerNum:             uint64(len(peerStrs)),
		role:                pb.Role_Candidate,
		term:                0,
		heartbeatInterval:   heartbeatInterval,
		log:                 []*pb.LogEntry{{Term: 0}}, // Append an empty log entry for prev log index and term
		tickerChange:        make(chan time.Duration),
	}

	// Set ticker
	svr.ticker = time.NewTicker(svr.getTimeOut())
	go svr.heartBeat(context.Background())
	return svr
}

// Publish a message
func (s *RaftServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	// For the initial version, let's block the message if the server is in candidate role.
	for s.role == pb.Role_Candidate {
	}
	log.Printf("publish %+v", req)
	if s.role == pb.Role_Leader {
		// Append to log
		logEntry := &pb.LogEntry{Term: s.term, Topic: req.Topic, Msg: req.Msg}
		log.Printf("append log: %+v", logEntry)
		s.log = append(s.log, logEntry)
		// Broadcast the message immediately
		s.broadcastPublishRequest(ctx, logEntry)
		return s.SingleMachineServer.Publish(ctx, req)
	}

	// Redirect to Leader
	log.Printf("redirect to leader: %s", s.leaderID)
	leader, _ := s.peers.Load(s.leaderID)
	return leader.(*Peer).client.Publish(ctx, req)
}

// AppendEntries appends log entries or deal with heart beat message.
func (s *RaftServer) AppendEntries(ctx context.Context, req *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	log.Printf("AppendEntries: %+v", req)
	s.role = pb.Role_Follower
	s.leaderID = req.LeaderId
	// Reset ticker
	s.tickerChange <- s.getTimeOut()
	if s.voteFor != "" {
		s.voteFor = ""
	}
	if s.term > req.Term {
		return &pb.AppendEntriesResponse{
			Term:   s.term,
			Status: pb.AppendEntriesResponse_FAILURE,
			Info:   "Leader's term need to be updated",
		}, nil
	}
	s.term = req.Term
	if len(req.Entries) > 0 {
		if len(s.log) <= int(req.PrevLogIndex) || s.log[req.PrevLogIndex].Term != req.PrevLogTerm {
			return &pb.AppendEntriesResponse{
				Term:   s.term,
				Status: pb.AppendEntriesResponse_FAILURE,
				Info:   "Prev index and term mismatch",
			}, nil
		}
		s.log = append(s.log[:req.PrevLogIndex+1], req.Entries...)
		for _, e := range req.Entries {
			s.SingleMachineServer.Publish(ctx, &pb.PublishRequest{Msg: e.Msg, Topic: e.Topic})
		}
	}
	return &pb.AppendEntriesResponse{
		Term:   s.term,
		Status: pb.AppendEntriesResponse_SUCCESS,
	}, nil
}

// RequestVote processes request vote message
func (s *RaftServer) RequestVote(ctx context.Context, req *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	log.Printf("RequestVote: %+v", req)
	voteGrant := s.role != pb.Role_Leader &&
		s.term <= req.Term &&
		(s.voteFor == "" || s.voteFor == req.CandidateId) &&
		len(s.log) <= int(req.LastLogIndex+1)
	if voteGrant {
		s.voteFor = req.CandidateId
	}
	return &pb.RequestVoteResponse{
		Term:        s.term,
		VoteGranted: voteGrant,
	}, nil
}

func (s *RaftServer) heartBeat(ctx context.Context) {
	for {
		select {
		case <-s.ticker.C:
			switch s.role {
			case pb.Role_Leader:
				log.Printf("Leader heartbeat")
				// Send heart beat request to all the followers
				s.sendHeartBeatRequest(ctx)
			case pb.Role_Follower, pb.Role_Candidate:
				log.Printf("Timeout heartbeat")
				// Sleep a random time to avoid conflict for leader election
				time.Sleep(time.Duration(float32(s.heartbeatInterval) * float32(rand.Intn(10)) / float32(10)))
				if s.role == pb.Role_Follower {
					s.term++
					s.role = pb.Role_Candidate
				}
				if s.voteFor != "" || s.role == pb.Role_Follower {
					continue
				}
				// Request vote
				voteChannel, voteNum, total := make(chan bool), uint64(1), uint64(0)
				s.peers.Range(func(k, v interface{}) bool {
					go func() {
						res, _ := v.(*Peer).sideCar.RequestVote(ctx, &pb.RequestVoteRequest{
							Term:         s.term,
							CandidateId:  s.id,
							LastLogIndex: uint64(len(s.log) - 1),
							LastLogTerm:  s.log[len(s.log)-1].Term,
						})
						log.Printf("Vote response %+v, received from %s", k)
						if res == nil {
							voteChannel <- false
						} else {
							voteChannel <- res.VoteGranted
						}
					}()
					return true
				})
				go func() {
					for s.role == pb.Role_Candidate && !s.isMajority(voteNum) && total < s.peerNum {
						select {
						case v := <-voteChannel:
							total++
							if v {
								voteNum++
							}
						default:
						}
					}
					if s.role == pb.Role_Candidate && s.isMajority(voteNum) &&
						s.voteFor == "" {
						// Upgrade to leader
						log.Printf("Upgrade to leader")
						s.role = pb.Role_Leader
						s.sendHeartBeatRequest(ctx)
						s.tickerChange <- s.heartbeatInterval
					}
				}()
			}
		case d := <-s.tickerChange:
			s.resetTicker(d)
		}
	}
}

func (s *RaftServer) broadcastPublishRequest(ctx context.Context, logEntry *pb.LogEntry) {
	s.peers.Range(func(k, v interface{}) bool {
		index := v.(*Peer).logIndex
		if index == len(s.log)-1 {
			res, err := v.(*Peer).sideCar.AppendEntries(ctx,
				&pb.AppendEntriesRequest{
					LeaderId:     s.id,
					Term:         s.term,
					Entries:      []*pb.LogEntry{logEntry},
					PrevLogTerm:  s.log[index-1].Term,
					PrevLogIndex: uint64(index - 1),
				})
			if err == nil && res.Status == pb.AppendEntriesResponse_SUCCESS {
				v.(*Peer).logIndex = len(s.log)
			}
		}
		return true
	})
}

func (s *RaftServer) sendHeartBeatRequest(ctx context.Context) {
	s.peers.Range(func(k, v interface{}) bool {
		go func() {
			req, index := &pb.AppendEntriesRequest{}, v.(*Peer).logIndex
			req.LeaderId = s.id
			req.Term = s.term
			if index <= len(s.log)-1 {
				req.Entries = s.log[index:]
				req.PrevLogTerm = s.log[index-1].Term
				req.PrevLogIndex = uint64(index - 1)
			}
			res, err := v.(*Peer).sideCar.AppendEntries(ctx, req)
			log.Printf("Respond for heart beat: %+v", res)
			if err == nil && res.Status == pb.AppendEntriesResponse_SUCCESS {
				v.(*Peer).logIndex = len(s.log)
			}
		}()
		return true
	})
}

func (s *RaftServer) isMajority(voteNum uint64) bool {
	return voteNum > (s.peerNum+1)/2
}

func (s *RaftServer) getTimeOut() time.Duration {
	return s.heartbeatInterval * 3
}

func (s *RaftServer) resetTicker(t time.Duration) {
	s.ticker.Stop()
	s.ticker = time.NewTicker(t)
}
