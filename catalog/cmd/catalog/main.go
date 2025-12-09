package main

import (
	"log"
	"time"

	"github.com/Cutshadows/microservices-go/catalog"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://catalog:catalog@localhost:5432/catalog?sslmode=disable"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var r catalog.Repository
	for {
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(2 * time.Second)
	}

	defer r.Close()
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, 8080))
}
