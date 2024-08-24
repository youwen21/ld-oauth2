package conf

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

func init() {
	// init by config file
	initByConfig()

	// init default
	initDefault()

	initDbs()
}

func initByConfig() {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		fmt.Println(filename)
	}
	viper.SetConfigName(GetConfigNameByRunMode()) // name of config file (without extension)
	viper.AddConfigPath(".")                      // optionally look for config in the working directory
	//viper.AddConfigPath("./conf")                 // optionally look for config in the working directory
	if ok {
		viper.AddConfigPath(path.Dir(path.Dir(filename))) // optionally look for config in the working directory
		viper.AddConfigPath(path.Dir(filename))           // optionally look for config in the working directory
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Printf("Fatal error config file: %v \n", err)
		fmt.Printf("read config from embeg data")
		viper.SetConfigType("toml")
		viper.ReadConfig(bytes.NewReader(GetEmbedConfigByRunMode()))
	} else {
		fmt.Println("viper use embed file:", viper.ConfigFileUsed())
	}

	err = viper.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("viper Unmarshal error: %w \n", err))
	}
}

func initDefault() {

}

func initDbs() {
	if Config.MysqlDefault.HOST != "" {
		Config.MysqlDefault.initConnPool()
	}

	if Config.RedisDefault.Host != "" {
		Config.RedisDefault.initConnPool()
	}
}
