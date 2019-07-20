package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"github.com/spatocode/minercave/app"
	"github.com/spatocode/minercave/pool"
)

var cfg pool.Config

func main() {
	app.Exec()
	//readConfig(&cfg)
}


func readConfig(cfg *pool.Config) {
	log.Printf("%v", cfg)
	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&cfg); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
	log.Printf("%v", cfg)
}