package account

import (
	"log"
	"time"

	"github.com/Cutshadows/microservices-go/account"
	"github.com/kelseyhightower/envconfig"
	"github.com/ydb-platform/ydb-go-sdk/v3/retry"
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
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})

	defer r.Close()
	log.Println("listening on port 8080 ...")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
