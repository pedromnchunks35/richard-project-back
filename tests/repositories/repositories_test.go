package repositories

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"richard-project-back/configs/connector/postgresql"
	"testing"
	"time"
)

var username string = "root"
var database string = "postgres"
var password string = "12341234"

var pgContainer *postgres.PostgresContainer
var dbVault *postgresql.PostgresDbVault

func before(t *testing.T) {
	dbVault = &postgresql.PostgresDbVault{}
	context := context.Background()
	pgContainerLocal, err := postgres.Run(
		context,
		"postgres",
		postgres.WithDatabase(database),
		postgres.WithUsername(username),
		postgres.WithPassword(password),
		postgres.WithInitScripts("../script.sql"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		t.Fatalf("error trying to create the postgres container: %s", err.Error())
	}

	pgContainer = pgContainerLocal
}

func after() {
	err := pgContainer.Terminate(context.Background())
	if err != nil {
		log.Fatalf("error terminating the container %s", err.Error())
	}
}

func Test_db_connection(t *testing.T) {
	before(t)
	t.Cleanup(after)
	ctx := context.Background()

	connectionString, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.Nil(t, err)

	conn, err := dbVault.InitDbConnection(connectionString)
	assert.Nil(t, err)

	createTableResponse, err := dbVault.ExecuteNonQuery("CREATE TABLE USERS(ID SERIAL PRIMARY KEY,USERNAME TEXT NOT NULL);")
	assert.Nil(t, err)
	assert.Equal(t, createTableResponse, int64(0))

	insertDataResponse, err := dbVault.ExecuteNonQuery("INSERT INTO USERS (USERNAME) VALUES ('LOL');")
	assert.Nil(t, err)
	assert.Equal(t, insertDataResponse, int64(1))

	queryResponse, err := dbVault.ExecuteQuery("SELECT * FROM USERS;")
	assert.Nil(t, err)
	id := 0
	name := ""
	rows := (*queryResponse)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			t.Fatalf("failed to scan row: %v", err)
		}
	}
	assert.Equal(t, name, "LOL")
	assert.Equal(t, id, 1)

	defer conn.Close(context.Background())

	assert.Nil(t, err)
}
