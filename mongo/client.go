package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ClientConfig represents the configuration for a client.
type ClientConfig struct {
	Hosts               []string
	Username            string
	Password            string
	NoDefaultConnect    bool
	TimeoutMilliseconds uint32
}

// Client represents a MongoDB client.
// This wraps the Mongo-Go-Driver client.
type Client struct {
	client    *mgo.Client
	connected bool
	timeout   uint32
}

// NewClient creates new client based on the ClientConfig provided.
func NewClient(config ClientConfig) (*Client, error) {
	connStr := fmt.Sprintf(
		"mongodb://%s:%s@%s",
		config.Username,
		config.Password,
		config.Hosts[0],
	)
	mgoClient, err := mgo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		err = errors.Wrap(err, "Error Creating MongoDB Client")
		return nil, err
	}
	if config.TimeoutMilliseconds == 0 {
		config.TimeoutMilliseconds = 1000
	}
	client := &Client{
		client:  mgoClient,
		timeout: config.TimeoutMilliseconds,
	}

	if config.NoDefaultConnect {
		client.connected = false
		return client, err
	}
	err = client.Connect()
	return client, err
}

// Connect connects the created client to Database. This is a no-op if
// the client is already connected.
// This is also run by default unless "NoDefaultConnect" is specified in ClientConfig.
func (c *Client) Connect() error {
	if !c.connected {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(c.timeout)*time.Millisecond,
		)
		defer cancel()

		err := c.client.Connect(ctx)
		if err != nil {
			err = errors.Wrap(err, "Error Connecting MongoDB Client to Database")
			return err
		}
		c.connected = true
	}
	return nil
}

// Disconnect disconnects the created client from Database. This is a no-op if
// the client is already disconnected.
func (c *Client) Disconnect() error {
	if c.connected {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(c.timeout)*time.Millisecond,
		)
		defer cancel()

		err := c.client.Disconnect(ctx)
		if err != nil {
			err = errors.Wrap(err, "Error Disconnecting MongoDB Client from Database")
			return err
		}
		c.connected = false
	}
	return nil
}

// Database returns a handle for a given database.
func (c *Client) Database(dbName string) *mgo.Database {
	return c.client.Database(dbName)
}

// DriverClient returns the wrapped Mongo-Go-Driver client.
func (c *Client) DriverClient() *mgo.Client {
	return c.client
}
