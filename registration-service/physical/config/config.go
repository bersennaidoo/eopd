package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nuid"
)

// Config contains state and logic for setting up nats and db client
// connections.
type Config struct {
	cmu sync.Mutex

	id string

	nc *nats.Conn

	db *sql.DB

	kind string
}

func New(kind string) *Config {

	id := nuid.Next()
	return &Config{
		id:   id,
		kind: kind,
	}
}

// SetupConnectionToDB creates a connection to the database and stores connection in Config.
func (c *Config) SetupConnectionToDB(dbDriver string, connectionString string) error {

	c.cmu.Lock()
	defer c.cmu.Unlock()
	db, err := sql.Open(dbDriver, connectionString)
	if err != nil {
		panic(err.Error())
	}
	c.db = db

	return err
}

// Config returns the current Database connection.
func (c *Config) DB() *sql.DB {

	c.cmu.Lock()
	defer c.cmu.Unlock()

	return c.db
}

// SetupConnectionToNATS setup a connection to nats, registers callbacks and stores
// connection in Config.
func (c *Config) SetupConnectionToNATS(servers string, options ...nats.Option) error {

	options = append(options, nats.Name(c.Name()))

	c.cmu.Lock()
	defer c.cmu.Unlock()

	// Connect to NATS with customized options.
	nc, err := nats.Connect(servers, options...)
	if err != nil {
		return err
	}

	c.nc = nc

	// Setup NATS event callbacks
	// Handle protocol errors and slow consumers cases.
	nc.SetErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, err error) {
		log.Printf("NATS error: %s\n", err)
	})
	nc.SetReconnectHandler(func(_ *nats.Conn) {
		log.Println("Reconnected to NATS!")
	})
	nc.SetDisconnectHandler(func(_ *nats.Conn) {
		log.Println("Disconnected from NATS!")
	})
	nc.SetClosedHandler(func(_ *nats.Conn) {
		panic("Connection to NATS is closed!")
	})

	return err
}

// NATS returns the current NATS connection.
func (c *Config) NATS() *nats.Conn {

	c.cmu.Lock()
	defer c.cmu.Unlock()

	return c.nc
}

// ID returns the ID from the Config
func (c *Config) ID() string {

	c.cmu.Lock()
	defer c.cmu.Unlock()

	return c.id
}

// Name is the label used to identify the NATS connection.
func (c *Config) Name() string {

	c.cmu.Lock()
	defer c.cmu.Unlock()

	return fmt.Sprintf("%s:%s", c.kind, c.id)
}

// Shutdown closes connections to DB and NATS.
func (c *Config) Shutdown() error {

	c.NATS().Close()
	defer c.DB().Close()

	return nil
}
