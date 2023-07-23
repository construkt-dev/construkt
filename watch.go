package construkt

import "k8s.io/apimachinery/pkg/watch"

type Watch[T any] interface {
	ResultChan() <-chan Event[T]
	Stop()
}

type watchWrapper[T any] struct {
	watch   watch.Interface
	results chan Event[T]
}

func (w *watchWrapper[T]) ResultChan() <-chan Event[T] {
	return w.results
}

func (w *watchWrapper[T]) Stop() {
	w.watch.Stop()
}

func (w *watchWrapper[T]) run() {
	defer close(w.results)
	for {
		select {
		case event, ok := <-w.watch.ResultChan():
			if !ok {
				return
			}
			if event.Type == watch.Error {
				w.results <- Event[T]{
					Type:   watch.Error,
					Object: nil,
				}
				return
			}
			w.results <- Event[T]{
				Type:   event.Type,
				Object: event.Object.(*Object[T]),
			}
		}
	}
}

func newWatch[T any](actual watch.Interface) Watch[T] {
	w := &watchWrapper[T]{
		results: make(chan Event[T]),
		watch:   actual,
	}
	go w.run()
	return w
}

type Event[T any] struct {
	Type   watch.EventType
	Object *Object[T]
}
