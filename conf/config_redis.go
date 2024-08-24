package conf

import (
	"github.com/go-redis/redis/v8"
)

type RedisCfg struct {
	Host     string `toml:"host" mapstructure:"host"`
	Port     string `toml:"port" mapstructure:"port"`
	Password string `toml:"password" mapstructure:"password"`
	Db       int    `toml:"db" mapstructure:"db"`

	instance *redis.Client
}

func (cfg *RedisCfg) initConnPool() {
	cfg.instance = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
	})
}

func (cfg *RedisCfg) GetClient() *redis.Client {
	return cfg.instance
}
