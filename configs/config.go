package configs

import (
	"log"
)

// Configs .
type Configs struct {
	IsInitialized bool
	DBConfig      DatabaseConfig
	TwConfig      TwitterConfig
}

// Cfg var accessed from global
var Cfg Configs

// InitConfig set the config var t
func InitConfig() (err error) {

	log.Println("Initializing Config Started")

	dbConfig, err := loadDatabaseConfig()
	if err != nil {
		log.Fatalln("Unable to load database config:", err)
		return err
	}

	twConfig, err := loadTwitterConfig()
	if err != nil {
		log.Fatalln("Unable to load database config:", err)
		return err
	}

	Cfg = Configs{
		IsInitialized: true,
		DBConfig:      dbConfig,
		TwConfig:      twConfig,
	}

	log.Println("Config Initialized")

	return
}

// GetConfig .
func GetConfig() Configs {

	if !Cfg.IsInitialized {
		InitConfig()
	}

	return Cfg

}
