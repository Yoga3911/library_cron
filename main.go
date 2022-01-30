package main

import (
	"context"
	"library_cron/config"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jasonlvhit/gocron"
)

var db *pgxpool.Pool = config.DatabaseConnection()

const up = `CALL updateCoin()`

func main() {
	log.Println("CRON JOB")
	gocron.Every(1).Hour().Do(func()  {
		_, err := db.Exec(context.Background(), up)
		if err != nil {
			log.Println(err)
		}

		log.Println("Cron job done!")
	})

	<-gocron.Start()
}