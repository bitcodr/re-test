package config

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	// MEMORY connection name.
	MEMORY = "memory"
)

// DB structure - it is possible in projects that has multiple
// source of data for persistent, caching and events
// in DB structure we can have other flags like default connection, etc. besides the other connections.
type DB struct {
	Connections map[string]*Connection
}

// Connection structure - it is possible in projects that has multiple
// source of data for persistent, caching and events
type Connection struct {
	Name string
}

// Service - is a general model for our service that can contains port, read, write, idle timeout
type Service struct {
	PORT         string
	Name         string
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

// Config - complete config structure that will be mapped from config.yml
type Config struct {
	DB
	Service
}

// Load used in the main file to load config file from ./config.d/config.yml to Config structure
func Load(ctx context.Context) (*Config, error) {
	v := viper.New()

	setDefaultServiceConfig(v)

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("/app/config.d")
	v.AddConfigPath("config.d")
	v.AddConfigPath(".\\config.d")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error in reading configs from file: %+v \n\n", err)
	}

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var config Config

	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	watchConfigChanges(ctx, v)

	return &config, nil
}

// watchConfigChanges, when a project is running with this function, we are able to modify the config
// without recompile it
func watchConfigChanges(_ context.Context, v *viper.Viper) {
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		var config Config
		err := v.Unmarshal(&config)
		if err != nil {
			log.Fatalln("Fatal error when unmarshal config:", err)
		}
	})
}

// setDefaultServiceConfig add default value for service in case the config.yml is empty
func setDefaultServiceConfig(v *viper.Viper) {
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.name", "re")
	v.SetDefault("server.read_timeout", 30*time.Second)
	v.SetDefault("server.write_timeout", 30*time.Second)
	v.SetDefault("server.idle_timeout", 30*time.Second)
}
