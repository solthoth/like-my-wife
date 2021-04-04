package configuration

import (
	"fmt"
	"time"

	jww "github.com/spf13/jwalterweatherman"
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

type LoggingConfigurations struct {
	LogLevel string
}

func Init(appName string) *Configurations {
	LoadConfig(appName, ".")
	LoadConfig(fmt.Sprintf(".%s",appName), "$HOME")
	return Build()
}

func LoadConfig(appName string, path string) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(appName)
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		jww.ERROR.Printf("Unable to decode into struct, %v\n", err)
	}
	viper.MergeConfigMap(v.AllSettings())
}

func Build() *Configurations {
	var config Configurations
	viper.Unmarshal(&config)
	return &config
}