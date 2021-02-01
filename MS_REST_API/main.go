package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/iavealokin/microservices/MS_REST_API/app/apiserver"
	_ "github.com/lib/pq"
)

var configPath string

func init(){
	flag.StringVar(&configPath,"config-path","configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()
	config:=apiserver.NewConfig()
	_,err := toml.DecodeFile(configPath,config)
	if err!= nil{
		log.Fatal(err)
	}
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}


