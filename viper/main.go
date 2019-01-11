package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type config struct {
	Environment string
	Dur         time.Duration
}

func main() {

	conf, err := newConfig("./test.toml")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(conf)

}

func newConfig(path string) (*config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &config{
		Environment: viper.GetString("environment"),
		Dur:         viper.GetDuration("dur"),
	}, nil
}
