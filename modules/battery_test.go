package modules

import "testing"

func TestBattery_Status(t *testing.T) {

	b := &Battery{
		Battery:  "TEST",
		Charging: "(+)",
	}

	var status, capacity string

	queryFn = func(battery, attrib string) string {
		switch attrib {
		case "status":
			return status
		case "capacity":
			return capacity
		default:
			return ""
		}
	}

	tests := []struct {
		status, capacity string
		expected         string
	}{
		{"Charging", "99", "(+)99%"},
		{"Discharging", "23", "23%"},
	}

	for _, test := range tests {
		t.Run("test battery status", func(t *testing.T) {
			status = test.status
			capacity = test.capacity

			actual := b.Status()

			if test.expected != actual {
				t.Errorf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}
