package provider

import (
	"okki-status/core"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Updates struct {
	CheckCommand      string   `yaml:"check_command"`
	UpdateCommand     string   `yaml:"update_command"`
	CheckCommandArgs  []string `yaml:"check_command_args"`
	UpdateCommandArgs []string `yaml:"update_command_args"`
	RefreshMinutes    int      `yaml:"refresh_minutes"`
	IgnoreExitError   bool     `yaml:"ignore_exit_error"`
}

func (u *Updates) Run(ch chan<- core.Update, events <-chan core.Event) {
	ch <- u.getUpdateCount()

	if u.RefreshMinutes == 0 {
		u.RefreshMinutes = 60
	}

	t := time.NewTicker(time.Duration(u.RefreshMinutes) * time.Minute)
	for {
		select {
		case <-t.C:
			ch <- u.getUpdateCount()
		case e := <-events:
			if _, ok := e.(*core.Click); ok {
				u.execUpdate()
			}
		}
	}
}

func (u *Updates) getUpdateCount() core.Update {
	out, err := exec.Command(u.CheckCommand, u.CheckCommandArgs...).Output()
	if err != nil {
		_, isExitError := err.(*exec.ExitError)
		if !u.IgnoreExitError || u.IgnoreExitError && !isExitError {
			return &core.ErrorUpdate{P: u}
		}
	}
	count := countNonEmptyLines(string(out))
	return &core.SimpleUpdate{
		P: u,
		T: strconv.Itoa(count),
	}
}

func (u *Updates) execUpdate() {
	_ = exec.Command(u.UpdateCommand, u.UpdateCommandArgs...).Start()
}

func countNonEmptyLines(s string) int {
	lines := strings.Split(string(s), "\n")
	count := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			count++
		}
	}
	return count
}
