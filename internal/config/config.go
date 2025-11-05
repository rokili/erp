package config

import (
	"github.com/spf13/viper"
)

type Config struct {
<<<<<<< HEAD
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Host  string `mapstructure:"host"`
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
=======
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
>>>>>>> eb04c5bc7c6a1998a4c109ada9e19202dab00b44
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 设置默认值
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8080")
<<<<<<< HEAD
	viper.SetDefault("server.debug", true)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "Aa123456")
	viper.SetDefault("database.name", "erp_system")
	viper.SetDefault("database.sslmode", "disable")

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		// 如果配置文件不存在，则使用默认配置
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
=======
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "erp_user")
	viper.SetDefault("database.name", "erp_system")
	viper.SetDefault("database.sslmode", "disable")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
>>>>>>> eb04c5bc7c6a1998a4c109ada9e19202dab00b44
}
