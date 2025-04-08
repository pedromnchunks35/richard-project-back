package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"sync"
)

type PostgresDbVault struct {
	initializer sync.Once
	db          *pgx.Conn
	dbMutex     sync.Mutex
}

type DbProfile struct {
	Username string
	Password string
	Ip       string
	Port     string
	DbName   string
}

func (p PostgresDbVault) InitDbConnection(dbProfile *DbProfile) (*pgx.Conn, error) {
	var err error
	p.initializer.Do(func() {
		conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			dbProfile.Username,
			dbProfile.Password,
			dbProfile.Ip,
			dbProfile.Port,
			dbProfile.DbName,
		))
		if err != nil {
			fmt.Printf("error creating the initial connection %s", err.Error())
		}
		p.db = conn
	})
	return p.db, err
}

func (p PostgresDbVault) ExecuteQuery(query string) (*pgx.Rows, error) {
	p.dbMutex.Lock()
	defer p.dbMutex.Unlock()

	response, err := p.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("some error occurred querying: %s", err.Error())
	}

	return &response, nil
}

func (p PostgresDbVault) ExecuteNonQuery(query string) (int64, error) {
	p.dbMutex.Lock()
	defer p.dbMutex.Unlock()

	rowsAffected, err := p.db.Exec(context.Background(), query)
	if err != nil {
		return 0, fmt.Errorf("some error occured while doing non-query: %s", err.Error())
	}

	return rowsAffected.RowsAffected(), nil
}
