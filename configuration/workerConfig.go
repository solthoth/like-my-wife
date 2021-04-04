package configuration

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Configurations struct {
	WorkerContext WorkerConfigurations
	SocialMediaIdentifier SocialMediaIdentifierConfigurations
}

type WorkerConfigurations struct {
	WaitTime time.Duration
}

type SocialMediaIdentifierConfigurations struct {
	Instagram string
}

func Init(appName string) *Configurations {
	Configure(appName)
	return Build()
}

func Configure(appName string)  {
	log := log.Default()
	viper.SetConfigName(appName)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to decode into struct, %v\n", err)
	}
	
	LoadConfig(fmt.Sprintf("~/.%s", appName))
}

func LoadConfig(path string) {
	log := log.Default()
	viper.AddConfigPath(path)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Unable to decode into struct, %v\n", err)
	}
}

func Build() *Configurations {
	var config Configurations
	viper.Unmarshal(&config)
	return &config
}