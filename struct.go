package RedisPull

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
	"log"
	"os"
)

type Connect struct {
	GoRedis *redis.Client
	Ruedis  rueidis.Client
}

type Connections struct {
	m ConnectMap
	d ConfigMap
}

func NewConnections() (*Connections, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	var con Connections
	err1 := con.LoadDomains()
	if err1 != nil {
		return &con, err1
	}
	con.initConnections()

	return &con, nil
}

func (c *Connections) getConnection(domain string) (Connect, bool) {
	conn, err := c.m[domain]
	return conn, err
}

func (c *Connections) LoadDomains() error {
	filename := os.Getenv("PATH_TO_CONFIG")
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var configs ConfigMap
	if err := json.Unmarshal(data, &configs); err != nil {
		return err
	}
	c.d = configs
	return nil
}

func (c *Connections) initConnections() {
	for domain, config := range c.d {
		connect, err := connect(config)
		if err != nil {
			fmt.Printf("Can't connect to redis server %#v\n", config)

		} else {
			c.m[domain] = connect
		}
	}
}

type ConfigData struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type ConfigMap map[string]ConfigData
type ConnectMap map[string]Connect

func connect(config ConfigData) (Connect, error) {
	var connect Connect
	var err error
	connect.Ruedis, err = ruedisConnect(config)
	if err != nil {
		fmt.Printf("Can't connect to redis server %#v\n", config)

	}

	connect.GoRedis, err = goRedisConnect(config)
	if err != nil {
		fmt.Printf("Can't connect to redis server %#v\n", config)

	}
	return connect, err
}

func ruedisConnect(config ConfigData) (rueidis.Client, error) {
	conn, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{config.Host},
		Password:    config.Password,
		SelectDB:    config.DB,
	})

	return conn, err

}

func goRedisConnect(config ConfigData) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := conn.Ping(context.Background()).Result()
	return conn, err
}
