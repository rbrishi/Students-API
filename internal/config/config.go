package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)
type Config struct{
	Env string `yaml:"env" env: "ENV" env-required:"true"` //struct tags
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server`
}

type HTTPServer struct{
	Addr string
}


//go get -u github.com/ilyakaznacheev/cleanenv used



func MustLoad *Config{
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == ""{
		flags := flag.String("config","","path to config file")
		flag.Parse()
		configPath = *flags

		if configPath == ""{
			log.Fatal("config path is not set") //fatal error that means program will run further
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file does not exist: %s", configPath)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil{
		log.Fatalf("cannot read config: %s", err.Error())
	}
	return &cfg
	
}