package main

type Subject interface {
	Subscribe(o Observer) (bool, error)
	Unsubscribe(o Observer) (bool, error)
	Notify(o Observer) (bool, error)
}
