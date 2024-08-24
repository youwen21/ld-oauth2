package conf

type AppConfig struct {
	Smtp SmtpCfg `mapstructure:"smtp"` //,squash

	MysqlDefault MysqlCfg `toml:"mysql_default" mapstructure:"mysql_default"` //,squash

	MysqlJourney MysqlCfg `toml:"mysql_journey" mapstructure:"mysql_journey"` //,squash

	RedisDefault RedisCfg `toml:"redis_default" mapstructure:"redis_default"` //,
}

var Config *AppConfig
