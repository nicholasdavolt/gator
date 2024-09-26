package main

import (
	"fmt"
	"log"

	"github.com/nicholasdavolt/gator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Print(err.Error())
	}

	err = cfg.SetUser("nick")

	if err != nil {
		log.Print(err)
	}

	cfg, err = config.Read()

	if err != nil {
		log.Print(err)
	}

	fmt.Printf("DB: %s User: %s", cfg.Db_url, cfg.Current_user_name)

}
