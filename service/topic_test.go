package service

import (
	pb "github.com/cs244b-2020-spring-pubsub/pubsub/proto"
	"testing"
)

func TestPersistTopics(t *testing.T) {
	topicToServersMap = make(map[string][]pb.PubSub_SubscribeServer)
	messageHistoryQueue = nil

	var stream1, stream2, stream3 pb.PubSub_SubscribeServer
	PersistTopics(stream1, []*pb.Topic{{Name: "Topic1"}})
	PersistTopics(stream2, []*pb.Topic{{Name: "Topic1"}, {Name: "Topic2"}})
	PersistTopics(stream3, []*pb.Topic{{Name: "Topic3"}})

	if len(GetServersForTopic("Topic1")) != 2 ||
		GetServersForTopic("Topic1")[0] != stream1 ||
		GetServersForTopic("Topic1")[1] != stream2 {
		t.Errorf("Topic1 persist error %v", len(GetServersForTopic("Topic1")))
	}

	if len(GetServersForTopic("Topic2")) != 1 ||
		GetServersForTopic("Topic2")[0] != stream2 {
		t.Errorf("Topic2 persist error %v", len(GetServersForTopic("Topic2")))
	}

	if len(GetServersForTopic("Topic3")) != 1 ||
		GetServersForTopic("Topic3")[0] != stream3 {
		t.Error("Topic3 persist error")
	}
}
