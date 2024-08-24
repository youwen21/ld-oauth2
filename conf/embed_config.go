package conf

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed config.toml config_*.toml
var embedConfig embed.FS

func GetEmbedConfigByRunMode() []byte {
	configFile := fmt.Sprintf("%s.toml", GetConfigNameByRunMode())
	data, _ := embedConfig.ReadFile(configFile)
	return data
}
