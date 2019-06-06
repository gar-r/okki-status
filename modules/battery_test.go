package modules

import "testing"

func TestBattery_Status(t *testing.T) {

	b := &Battery{
		Battery:  "BAT0",
		Charging: "[C]", Discharging: "[D]",
		Full: "[F]", High: "[H]", Normal: "N", Low: "[L]", Critical: "[C]", Unknown: "[?]",
	}

	var status, level, capacity string

	queryFn = func(battery, attrib string) string {
		switch attrib {
		case "status":
			return status
		case "capacity":
			return capacity
		case "capacity_level":
			return level
		default:
			return ""
		}
	}

	t.Run("test error", func(t *testing.T) {
		status = "error"
		level = "error"
		capacity = "error"

		expected := "[?] error%"
		actual := b.Status()

		if expected != actual {
			t.Errorf("expected '%s', got '%s'", expected, actual)
		}
	})

}
