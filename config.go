package nilan

import "os"

// Config holds network settings of Nilan heatpump.
type Config struct {
	// NilanAdress is IP address and port of Nilan heatpump. Factory
	// default address is "/dev/ttyUSB1".
	NilanAddress string
}

// StandardConfig returns factory-default adress of Nilan heatpump
func StandardConfig() Config {
	return Config{NilanAddress: "/dev/ttyUSB1"}
}

// CurrentConfig reads NILAN_ADDRESS environment variable and returns
// configuration. If environment variable is not present, function returns
// standard config.
func CurrentConfig() Config {
	envConf := os.Getenv("NILAN_ADDRESS")
	if len(envConf) > 0 {
		return Config{NilanAddress: envConf}
	}
	return StandardConfig()
}
