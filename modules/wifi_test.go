package modules

import (
	"testing"
)

func TestWifi_Status(t *testing.T) {

	var mock string
	wifiFn = func(device string) []byte {
		return []byte(mock)
	}

	t.Run("wifi name and signal", func(t *testing.T) {
		mock = `Connected to 70:4f:57:61:7e:5e (on wlp59s0)
			SSID: okki
			signal: -64 dBm`

		wifi := &Wifi{Device: "wlp59s0"}

		actual := wifi.Status()
		expected := "okki (-64 dBm)"

		if expected != actual {
			t.Errorf("expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("wifi name without sighnal", func(t *testing.T) {
		mock = `Connected to 70:4f:57:61:7e:5e (on wlp59s0)
			SSID: okki`

		wifi := &Wifi{Device: "wlp59s0"}

		actual := wifi.Status()
		expected := "okki"

		if expected != actual {
			t.Errorf("expected '%s', got '%s'", expected, actual)
		}
	})

	t.Run("no wifi no signal", func(t *testing.T) {
		mock = "something"

		wifi := &Wifi{Device: "wlp59s0"}

		actual := wifi.Status()
		expected := "?"

		if expected != actual {
			t.Errorf("expected '%s', got '%s'", expected, actual)
		}
	})

}
