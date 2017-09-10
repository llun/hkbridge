package accessories

import (
	"github.com/brutella/hc/log"

	"container/list"
	"time"
)

type Action interface {
	Run()
	Name() string
	RemoveDuplicateCommand() bool
}

type Worker struct {
	actions  *list.List
	tickerCh <-chan time.Time
}

func NewWorker() *Worker {
	return &Worker{
		actions:  list.New(),
		tickerCh: time.Tick(1 * time.Second),
	}
}

func (w *Worker) Run() {
	for range w.tickerCh {
		if w.actions.Len() == 0 {
			continue
		}

		log.Debug.Printf("Workers queue %v", w.actions.Len())
		head := w.actions.Front()
		action, ok := w.actions.Remove(head).(Action)
		if !ok {
			log.Debug.Printf("Worker skip action")
			continue
		}

		log.Debug.Printf("Worker run %v", action.Name())
		action.Run()
	}
}

func (w *Worker) AddAction(action Action) {
	if w.actions.Len() > 0 {
		tail := w.actions.Back()
		lastAction, ok := tail.Value.(Action)
		if !ok {
			return
		}

		if lastAction.RemoveDuplicateCommand() &&
			lastAction.Name() == action.Name() {
			w.actions.Remove(tail)
		}
	}
	w.actions.PushBack(action)
}
