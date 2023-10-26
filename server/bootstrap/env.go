package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
    Host       string `mapstructure:"HOST"`
    User       string `mapstructure:"USER"`
    Pass       string `mapstructure:"PASS"`
    Port       string `mapstructure:"PORT"`
}

func NewEnv() *Env {
    env := Env{}
    viper.SetConfigFile(".env")

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal("Cannot find find env file.", err)
    }

    err = viper.Unmarshal(&env)
    if err != nil {
        log.Fatal("Environment failed to load.", err)
    }

    return &env
}
