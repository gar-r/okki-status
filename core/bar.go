package core

type Bar interface {
	Observer
	Renderer
}

type Renderer interface {
	Render()
}
