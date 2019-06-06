package modules

import "testing"

func TestBattery_Status(t *testing.T) {

	b := &Battery{
		Battery:  "BAT0",
		Charging: "[C]", Discharging: "[D]",
		Full: "[F]", High: "[H]", Normal: "[N]", Low: "[L]", Critical: "[C]", Unknown: "[?]",
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

	tests := []struct {
		status, level, capacity string
		expected                string
	}{
		{"error", "error", "error", "[?] error%"},
		{"Charging", "Full", "99", "[C] [F] 99%"},
		{"Discharging", "High", "85", "[D] [H] 85%"},
		{"Something", "Normal", "50", "[N] 50%"},
		{"Something", "Low", "50", "[L] 50%"},
		{"Something", "Critical", "50", "[C] 50%"},
	}

	for _, test := range tests {
		t.Run("test battery status", func(t *testing.T) {
			status = test.status
			level = test.level
			capacity = test.capacity

			actual := b.Status()

			if test.expected != actual {
				t.Errorf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}
