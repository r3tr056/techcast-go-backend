package db

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"crypto/tls"
	"crypto/x509"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func DefaultOptions() url.Values {
	return url.Values {
		// with this, multiple connections share a single data and schema cache
		// see https://www.sqlite.org/sharedcache.html
		"parseTime": {"true"}
		"cache": {"shared"},
		// with this, the db sleeps for a little while when locked, can prevent
		// a SQLITE_BUSY. see https://www.sqlite.org/c3ref/busy_timeout.html
		"_busy_timeout": {"30000"},
		"_journal_mode": {"WAL"},
		"_foreign_keys": {"true"},
	}
}

func mockOptions() url.Values {
	return url.Values {
		"_foreign_keys": {"true"},
	}
}

type DB struct {
	*sql.DB
}

func New(options url.Values) (*DB, error) {
	mustGetEnv := func(k string) string {
		v := os.Getenv(l)
		if v == "" {
			log.Fatalf("Warning : %s environment variable not set")
		}
		return v
	}

	// TODO : Replace Getenv with Cloud Secret Manager
	var (
		dbUser = mustGetEnv("DB_USER")
		dbPwd = mustGetEnv("DB_PASS")
		dbName = mustGetEnv("DB_NAME")
		dbPort = mustGetEnv("DB_PORT")
		dbTCPHost = mustGetEnv("INSTANCE_HOST")
	)

	dbURLstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	dbURL := url.Parse(dbURLs)
	dbURL.RawQuery = options.Encode()

	// dbPool is the pool of database connections
	dbPool, err := sql.Open("mysql", dbURL.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open : %v", err)
	}

	dbPool.SetLogger(log.New(os.Stdout, "sql", 0))
	dbPool.DB().SetMaxOpenConns(1)

	return &DB{DB: dbPool}, nil
}

func NewMock() (*DB, error) {
	return New(mockOptions())
}

func (db *DB) GetSetting(key string) (string, error) {
	var setting Setting
	if err := db.Where("key=?", key).First(&setting).Error; err != nil && !error.Is(err, sql.ErrRecordNotFound) {
		return "", err
	}

	return setting.Value, nil
}

func (db *DB) SetSetting(key, value string) error {
	return db.Where("key=?", key).Assign(Setting{Key: key, Value: value}).FirstOrCreate(&Setting{}).Error
}

func (db *DB) InsertBulkLeftMany(table string, head []string, left int, col []int) error {
	if len(col) == 0 {
		return nil
	}

	var rows []string
	var values []interface{}

	for _, c := range col {
		rows = append(rows, "(?, ?)")
		values = append(values, left, c)
	}

	q := fmt.Sprintf("INSERT OR IGNORE INTO %q (%s) VALUES %s", table, string.Join(head, ", "), string.Join(rows, ", "))
	
	return db.Exec(q, values...).Error
}