package metadata

import (
	"testing"
	"time"

	"github.com/benaan/flyrics/src/model"
)

type listenerMock struct {
	metadata *model.Metadata
	error    error
}

func (listener *listenerMock) GetMetadata() (*model.Metadata, error) {
	return listener.metadata, listener.error
}

func TestReturnsWhenNoListeners(t *testing.T) {
	manager := Manager{}
	manager.Listen()
}

func TestIteratesListeners(t *testing.T) {

	metadata := &model.Metadata{Status: model.PLAYING}
	listener := &listenerMock{metadata, nil}

	output := make(chan model.Metadata)
	stop := make(chan bool)
	manager := Manager{
		Delay:     0,
		Output:    output,
		Stop:      stop,
		Listeners: []Listener{listener},
	}
	go manager.Listen()

	if result := <-output; result != *metadata {
		t.Errorf("Expected metadata to be filled, received: %s", result)
	}

	if result := <-output; result != *metadata {
		t.Errorf("Expected metadata to be filled, received: %s", result)
	}

	stop <- true
	select {
	case <-output:
		t.Error("Shouldnt receive results after closing the channel")
	case <-time.After(time.Millisecond):
	}
}
