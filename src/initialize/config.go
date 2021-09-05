package initialize

import (
	"MSC2021/src/global"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() error {
	if err := InitViper(); err != nil {
		log.Println("Could not init viper.")
		return err
	} else if err := SetConfigWithViper(); err != nil {
		log.Println("Could not read config file, maybe you lost some configurations or the file is broken.")
		return err
	}
	return nil
}

func InitViper() error {
	v := viper.New()
	v.SetConfigName("config")              // name of config file (without extension)
	v.SetConfigType("yaml")                // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("/etc/MessageBoard/")  // path to look for the config file in
	v.AddConfigPath("$HOME/.MessageBoard") // call multiple times to add many search paths
	v.AddConfigPath(".")                   // optionally look for config in the working directory
	if err := v.ReadInConfig(); err != nil {
		log.Println("Could not access config file. maybe it is not exist?")
		return err
	}
	log.Println("Initialized Viper.")
	global.VIPER = v
	return nil
}

func SetConfigWithViper() error {
	if err := global.VIPER.Unmarshal(&global.CONFIG); err != nil {
		return err
	}
	global.LOGGER = Zap()
	global.LOGGER.Info("Initialized Zap logger.")
	return nil
}
