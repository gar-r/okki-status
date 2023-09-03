package core

type Observable interface {
	Attach(Observer)
	Notify()
}

type Observer interface {
	Update(Event)
}

type Event struct {
	Status string
}
