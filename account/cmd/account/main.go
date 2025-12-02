package main

import (
	"log"
	"time"

	"github.com/Cutshadows/microservices-go/account"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	for {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(2 * time.Second)
	}

	defer r.Close()
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
