package sql

import (
	"fmt"
	"github.com/jatis/oms/lib/log"
	"github.com/jmoiron/sqlx"
	"time"

	"github.com/XSAM/otelsql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//Db object
var (
	Master   *DB
	Slave    *DB
	dbTicker *time.Ticker
)

type (
	DBConfig struct {
		Driver          string `yaml:"driver"`
		MasterDSN       string `yaml:"masterDsn"`
		RetryInterval   string `yaml:"retryInterval"`
		MaxIdleConn     int    `yaml:"maxIdleConn"`
		MaxConn         int    `yaml:"maxConn"`
		ConnMaxLifetime string `yaml:"connMaxLifetime"`
	}

	DB struct {
		DBConnection    *sqlx.DB
		DBString        string
		RetryInterval   string
		MaxIdleConn     int
		MaxConn         int
		ConnMaxLifetime string
		doneChannel     chan bool
	}

	Store struct {
		master *sqlx.DB
	}

	Options struct {
		dbTx *sqlx.Tx
	}
)

func (s *Store) GetMaster() *sqlx.DB {
	return s.master
}

func New(cfg DBConfig) *Store {
	Master = &DB{
		DBString:        cfg.MasterDSN,
		RetryInterval:   cfg.RetryInterval,
		MaxIdleConn:     cfg.MaxIdleConn,
		MaxConn:         cfg.MaxConn,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
		doneChannel:     make(chan bool),
	}

	err := Master.ConnectAndMonitor(cfg.Driver)
	if err != nil {
		log.Fatal("Could not initiate Master DB connection: " + err.Error())
		return &Store{}
	}

	dbTicker = time.NewTicker(time.Second * 2)

	return &Store{master: Master.DBConnection}
}

func (d *DB) Connect(driverName string) error {
	db, err := otelsql.Open(driverName, d.DBString)
	if err != nil {
		return fmt.Errorf("failed to open DB connection: %w", err)
	}

	db.SetMaxOpenConns(d.MaxConn)
	db.SetMaxIdleConns(d.MaxIdleConn)

	if d.ConnMaxLifetime != "" {
		t, err := time.ParseDuration(d.ConnMaxLifetime)
		if err != nil {
			return err
		}
		db.SetConnMaxLifetime(t)
	}

	d.DBConnection = sqlx.NewDb(db, driverName)

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}

	return nil
}

// ConnectAndMonitor to database
func (d *DB) ConnectAndMonitor(driver string) error {
	err := d.Connect(driver)

	if err != nil {
		log.Printf("Not connected to database %v, trying", d.DBString)
		return err
	}

	timeTicker, err := time.ParseDuration(d.RetryInterval)
	if err != nil {
		return fmt.Errorf("failed to parse duration retry interval: %v", err)
	}

	ticker := time.NewTicker(timeTicker)
	go func() error {
		for {
			select {
			case <-ticker.C:
				if d.DBConnection == nil {
					d.Connect(driver)
				} else {
					err := d.DBConnection.Ping()
					if err != nil {
						log.Error("[Error]: DB reconnect error", err.Error())
						return err
					}
				}
			case <-d.doneChannel:
				return nil
			}
		}
	}()
	return nil
}
