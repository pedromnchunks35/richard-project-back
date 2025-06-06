package postgresql

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4"
)

type PostgresDbVault struct {
	initializer sync.Once
	Db          *pgx.Conn
	dbMutex     sync.Mutex
}

type DbProfile struct {
	Username string
	Password string
	Ip       string
	Port     string
	DbName   string
}

func (p *PostgresDbVault) InitDbConnection(connectionString string) (*pgx.Conn, error) {
	var err error
	p.initializer.Do(func() {
		conn, connErr := pgx.Connect(context.Background(), connectionString)
		if connErr != nil {
			fmt.Printf("error creating the initial connection %s\n", connErr.Error())
			return
		}
		p.Db = conn
	})
	if p.Db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	return p.Db, err
}

func (p *PostgresDbVault) CreateConnectionString(dbProfile *DbProfile) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbProfile.Username,
		dbProfile.Password,
		dbProfile.Ip,
		dbProfile.Port,
		dbProfile.DbName,
	)
}

func (p *PostgresDbVault) ExecuteQuery(query string) (*pgx.Rows, error) {
	p.dbMutex.Lock()
	defer p.dbMutex.Unlock()

	response, err := p.Db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("some error occurred querying: %s", err.Error())
	}

	return &response, nil
}

func (p *PostgresDbVault) ExecuteNonQuery(query string) (int64, error) {
	p.dbMutex.Lock()
	defer p.dbMutex.Unlock()

	rowsAffected, err := p.Db.Exec(context.Background(), query)
	if err != nil {
		return 0, fmt.Errorf("some error occured while doing non-query: %s", err.Error())
	}

	return rowsAffected.RowsAffected(), nil
}
