package bootstrap

import "github.com/spf13/viper"

const (
	ModeLocal      = "local"
	ModeDevelop    = "dev"
	ModeTest       = "test"
	ModeProduction = "prod"
)

// GetMode ...
func GetMode() string {
	return viper.GetString("mode")
}

// IsDevelopment ...
func IsDevelopment() bool {
	var isDev bool

	switch GetMode() {
	case ModeLocal, ModeDevelop, ModeTest:
		isDev = true
	}

	return isDev
}
