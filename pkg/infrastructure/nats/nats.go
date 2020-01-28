package nats

import (
	stan "github.com/nats-io/go-nats-streaming"
	"log"
)

// Configurator is an dependency used for config retrieving
type Configurator func(key string) string

type config struct {
	url       string
	clusterID string
	clientID  string
}

func newConfig(configurator Configurator, paramsPrefix string) (config, error) {
	cfg := config{
		url:       configurator("url"),
		clusterID: configurator("clusterId"),
		clientID:  configurator("clientId"),
	}

	return cfg, nil
}

// Streaming wraps the NATS Streaming library
type Streaming struct {
	conn stan.Conn
}

// NewStreaming creates NATS Streaming client instance connected to the specified host
func NewStreaming(configurator Configurator, paramsPrefix string) (*Streaming, error) {
	cfg, err := newConfig(configurator, paramsPrefix)
	if err != nil {
		return nil, err
	}

	conn, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return &Streaming{
		conn: conn,
	}, nil
}

// ReadString implements StringReader interface
func (s *Streaming) ReadString() string {
	log.Fatal("implement me")
	return ""
}

// WriteString implements StringWriter interface
func (s *Streaming) WriteString(str string) {
	log.Fatal("implement me")
}

func connect(cfg config) (stan.Conn, error) {
	conn, err := stan.Connect(cfg.clusterID, cfg.clientID, func(options *stan.Options) error {
		options.NatsURL = cfg.url
		return nil
	})

	if err != nil {
		return nil, err
	}

	return conn, nil
}
