package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type dbEnv struct {
	host     string
	port     string
	dbName   string
	user     string
	password string
}

func (d dbEnv) getDbUrl() string {
	return fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v",
		d.user,
		d.password,
		d.host,
		d.port,
		d.dbName,
	)
}

func (d dbEnv) getDbMigrationUrl() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		d.user,
		d.password,
		d.host,
		d.port,
		d.dbName,
	)
}

var connPool *pgxpool.Pool

func InitDBPool() {
	var err error
	connPool, err = pgxpool.NewWithConfig(context.Background(), dbConfig())
	if err != nil {
		log.Fatalf("Error. Couldn't crate connection pool to db. %v\n", err)
	}
}

func CloseDBPool() {
	if connPool == nil {
		log.Fatalln("Error. Connection pool was not initialized")
	}
	connPool.Close()
}

func RunMigrations() {
	dbEnvInstance := getDbEnv()
	m, err := migrate.New(
		"file://db/migrations", dbEnvInstance.getDbMigrationUrl())
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Println("Migration status: ", err)
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}
}

func dbConfig() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	dbEnvInstance := getDbEnv()
	config, err := pgxpool.ParseConfig(dbEnvInstance.getDbUrl())
	if err != nil {
		log.Fatalf("Error. Failed to create db config: ", err)
	}

	config.MaxConns = defaultMaxConns
	config.MinConns = defaultMinConns
	config.MaxConnLifetime = defaultMaxConnLifetime
	config.MaxConnIdleTime = defaultMaxConnIdleTime
	config.HealthCheckPeriod = defaultHealthCheckPeriod
	config.ConnConfig.ConnectTimeout = defaultConnectTimeout

	config.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		pgxuuid.Register(c.TypeMap())
		return nil
	}

	config.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	config.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	config.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	log.Println("connection string: ", dbEnvInstance.getDbUrl())

	return config
}

func getDbEnv() dbEnv {
	dbEnvInstance := dbEnv{}

	host, ok := os.LookupEnv("DATABASE_HOST")
	if ok {
		dbEnvInstance.host = host
	} else {
		log.Fatalln("ERROR. no DATABASE_HOST env value is set")
	}

	port, ok := os.LookupEnv("DATABASE_PORT")
	if ok {
		dbEnvInstance.port = port
	} else {
		log.Fatalln("ERROR. no DATABASE_PORT env value is set")
	}

	dbName, ok := os.LookupEnv("DATABASE_NAME")
	if ok {
		dbEnvInstance.dbName = dbName
	} else {
		log.Fatalln("ERROR. no DATABASE_NAME env value is set")
	}

	user, ok := os.LookupEnv("DATABASE_USER")
	if ok {
		dbEnvInstance.user = user
	} else {
		log.Fatalln("ERROR. no DATABASE_USER env value is set")
	}

	password, ok := os.LookupEnv("DATABASE_PASSWORD")
	if ok {
		dbEnvInstance.password = password
	} else {
		log.Fatalln("ERROR. no DATABASE_PASSWORD env value is set")
	}

	return dbEnvInstance
}
