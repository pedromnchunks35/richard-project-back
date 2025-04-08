package repositories

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"richard-project-back/configs/connector/postgresql"
	"testing"
)

var username string = ""
var database string = ""
var password string = ""
var ip string = ""
var port string = ""

var pgContainer *postgres.PostgresContainer
var dbVault *postgresql.PostgresDbVault

func before(t *testing.T) {
	context := context.Background()
	pgContainerLocal, err := postgres.Run(
		context,
		"postgres",
		postgres.WithDatabase(database),
		postgres.WithUsername(username),
		postgres.WithPassword(password),
	)
	if err != nil {
		t.Fatalf("error trying to create the postgres container: %s", err.Error())
	}
	pgContainer = pgContainerLocal
}

func after(t *testing.T) {
	err := pgContainer.Terminate(context.Background())
	if err != nil {
		t.Fatalf("error terminating the container %s", err)
	}
}

func Test_db_connection(t *testing.T) {
	before(t)
	ctx := context.Background()

	containerIp, err := pgContainer.ContainerIP(ctx)
	assert.Nil(t, err)

	ports, err := pgContainer.Ports(ctx)
	assert.Nil(t, err)

	port, ok := ports["5432/tcp"]
	assert.True(t, ok)

	mappedPort := port[0].HostPort

	dbVault.InitDbConnection(&postgresql.DbProfile{
		Username: username,
		Password: password,
		Ip:       containerIp,
		Port:     mappedPort,
		DbName:   database,
	})
	after(t)
}
