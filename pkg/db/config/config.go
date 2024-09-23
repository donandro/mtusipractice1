package config

import pg "pillsreminder/pkg/db/postgres/config"

type DatalayerConfig struct {
	Postgres pg.PqConfig
}
