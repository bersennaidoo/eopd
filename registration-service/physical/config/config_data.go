package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration Data for config object.
type CFGData struct {
	Port     string
	dsn      string
	dbDriver string
	ndsn     string
}

func NewCFGData() CFGData {

	viper.SetConfigName("config")

	viper.SetConfigType("yaml")

	viper.AddConfigPath("./registration-service")
	viper.AddConfigPath(".")

	viper.SetDefault("environments.dev.port", ":3000")
	viper.SetDefault("environments.dev.dsn", "root:bersen@/eopd")
	viper.SetDefault("environments.dev.dbdriver", "mysql")
	viper.SetDefault("environments.dev.ndsn", "0.0.0.0:4222")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("err: %w", err))
	}

	return CFGData{
		Port:     viper.GetString("environments.dev.port"),
		dsn:      viper.GetString("environments.dev.dsn"),
		dbDriver: viper.GetString("environments.dev.dbdriver"),
		ndsn:     viper.GetString("environments.dev.ndsn"),
	}
}
