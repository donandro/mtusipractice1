package db

import (
	"pillsreminder/pkg/db/config"
	db "pillsreminder/pkg/db/postgres"
)

func InitDB(config config.DatalayerConfig) {
	db.InitDB(config.Postgres)
}
