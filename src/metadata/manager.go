package metadata

import (
	"time"

	"github.com/benaan/flyrics/src/model"
)

type Manager struct {
	Delay     time.Duration
	Output    chan<- model.Metadata
	Stop      <-chan bool
	Listeners []Listener
}

type Listener interface {
	GetMetadata() (*model.Metadata, error)
}

func (manager *Manager) SetOutput(output chan<- model.Metadata) {
	manager.Output = output
}

func (manager *Manager) Listen() {
	if !manager.hasValidInput() {
		return
	}
	for {
		if !manager.iterateListeners() {
			return
		}
		time.Sleep(manager.Delay)
	}
}
func (manager *Manager) hasValidInput() bool {
	return len(manager.Listeners) > 0
}

func (manager *Manager) iterateListeners() bool {
	for _, listener := range manager.Listeners {
		metadata, err := listener.GetMetadata()
		if err == nil {
			select {
			case manager.Output <- *metadata:
				return true
			case <-manager.Stop:
				return false
			}

		}
	}
	return true
}
