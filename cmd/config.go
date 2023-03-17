package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	SamplesNumber     int
	SampleOutFileName string
}

var cfg config

var vp *viper.Viper

func init() {
	vp = viper.New()
	vp.SetDefault("SamplesNumber", 100)
	vp.SetDefault("SampleOutFileName", "test.csv")
	if err := vp.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	log.Infof("Config: %#v", cfg)
}
