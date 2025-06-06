package main

import (
	"github.com/funtimecoding/go-podman/pkg/compose"
	"github.com/funtimecoding/go-podman/pkg/container"
	"time"
)

const PostgresPassword = "postgres"

var Postgres = container.New(
	"docker.io/library/postgres:16.6",
	"gp-postgres",
	5432,
	map[string]string{"POSTGRES_PASSWORD": PostgresPassword},
)

func main() {
	c := compose.New(Postgres)
	c.Create()
	defer c.Destroy()

	time.Sleep(30 * time.Second)
}
