package core

import "os/exec"

type Command interface {
	Exec() ([]byte, error)
}

type OsCommand struct {
	Name string
	Args []string
}

func (c *OsCommand) Exec() ([]byte, error) {
	return exec.Command(c.Name, c.Args...).Output()
}

type ErrCommand struct {
	Err error
}

func (c *ErrCommand) Exec() ([]byte, error) {
	return nil, c.Err
}

type ConstCommand struct {
	Output string
}

func (c *ConstCommand) Exec() ([]byte, error) {
	return []byte(c.Output), nil
}
