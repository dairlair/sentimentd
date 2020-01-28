package nats

import (
	stan "github.com/nats-io/go-nats-streaming"
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
		url:       configurator(paramsPrefix + ".url"),
		clusterID: configurator(paramsPrefix +".clusterId"),
		clientID:  configurator(paramsPrefix +".clientId"),
	}

	return cfg, nil
}

// NewStreaming creates NATS Streaming client instance connected to the specified host
func NewStreaming(configurator Configurator, paramsPrefix string) (stan.Conn, error) {
	cfg, err := newConfig(configurator, paramsPrefix)
	if err != nil {
		return nil, err
	}

	conn, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
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
