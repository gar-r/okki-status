package modules

import (
	"testing"
)

func TestWifi_Status(t *testing.T) {

	const output = `Connected to 70:4f:57:61:7e:5e (on wlp59s0)
		SSID: okki
		freq: 5180
		RX: 1175482 bytes (4697 packets)
		TX: 133033 bytes (712 packets)
		signal: -64 dBm
		rx bitrate: 260.0 MBit/s VHT-MCS 3 80MHz short GI VHT-NSS 2
		tx bitrate: 6.0 MBit/s

		bss flags:      short-slot-time
		dtim period:    1
		beacon int:     100`

	wifiFn = func(device string) []byte {
		return []byte(output)
	}

	wifi := &Wifi{Device: "wlp59s0"}

	actual := wifi.Status()
	expected := "okki (-64 dBm)"

	if expected != actual {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}

}
