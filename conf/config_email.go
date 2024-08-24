package conf

import "gofly/lib/libutils"

type SmtpCfg struct {
	Host     string `toml:"host" mapstructure:"host"`
	Port     int    `toml:"port" mapstructure:"port"`
	User     string `toml:"user" mapstructure:"user"`
	Password string `toml:"password" mapstructure:"password"`
}

func (cfg *SmtpCfg) GetClient() *libutils.Mail {
	return libutils.NewMail(cfg.User, cfg.Password, cfg.Host, cfg.Port)
}
