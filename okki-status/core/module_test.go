package core

import (
	"testing"
	"time"
)

func Test_Module_Info(t *testing.T) {

	m := makeTestModule("mod", "status")
	m.Icon = &StaticIcon{"I"}

	t.Run("icon before status", func(t *testing.T) {
		m.BlockOrder = IconFirst
		info := m.Info()
		expected := "|Istatus|"
		if info != expected {
			t.Errorf("expected %v, got %v", expected, info)
		}
	})

	t.Run("icon after status", func(t *testing.T) {
		m.BlockOrder = TextFirst
		info := m.Info()
		expected := "|statusI|"
		if info != expected {
			t.Errorf("expected %v, got %v", expected, info)
		}
	})
}

func Test_Schedule(t *testing.T) {
	ch := make(chan Module, 1)
	timeout := make(chan bool, 1)

	module := makeTestModule("mod", "status")
	module.Refresh = 50 * time.Millisecond

	module.Schedule(ch)
	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()

	select {
	case actual := <-ch:
		if actual != module {
			t.Errorf("expected %v, got %v", module, actual)
		}
	case <-timeout:
		t.Errorf("event not received in time")
	}
}

func makeTestModule(name, status string) Module {
	return Module{
		Name:       name,
		Status:     &MockStatusProvider{Status: status},
		Icon:       &StaticIcon{Icon: ""},
		Gap:        Gap{Before: "|", After: "|"},
		BlockOrder: IconFirst,
		Refresh:    1 * time.Second,
	}
}

type MockStatusProvider struct {
	Status string
}

func (s *MockStatusProvider) GetStatus() string {
	return s.Status
}
