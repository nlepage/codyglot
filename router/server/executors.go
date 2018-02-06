package server

import (
	"context"
	"net"
	"strconv"
	"time"

	executor_config "github.com/nlepage/codyglot/executor/config"
	"github.com/nlepage/codyglot/router/server/config"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type executor struct {
	host      string
	port      string
	alive     bool
	backoff   time.Duration
	languages []string
}

var (
	executors []*executor
)

func initExecutorsStatic() error {
	for _, hostport := range config.Executors {
		host, port, err := net.SplitHostPort(hostport)
		if err != nil {
			host = hostport
			port = strconv.Itoa(executor_config.DefaultPort)
			log.Warnf("Invalid executor address \"%s\" (%s), using value as host and default port %s", hostport, err.Error(), port)
		}

		executors = append(executors, &executor{
			host: host,
			port: port,
		})
	}

	return nil
}

func startPinging() {
	for _, exec := range executors {
		go func(exec *executor) {
			for {
				if exec.alive {
					time.Sleep(config.PingInterval)
				} else {
					time.Sleep(exec.backoff * 100 * time.Millisecond)
				}

				if alive := pingExecutor(exec); alive {
					if !exec.alive {
						exec.alive = true
						log.Infof("Executor %s is alive", exec.host)
						// FIXME send on chan
					}
					if exec.backoff != 0 {
						exec.backoff = 0
					}
				} else {
					if exec.alive {
						exec.alive = false
						log.Infof("Executor %s is dead", exec.host)
						// FIXME send on chan
					}
					if exec.backoff == 0 {
						exec.backoff = 1
					} else if exec.backoff*100 < config.MaxBackoff {
						exec.backoff <<= 1
					}
				}

			}
		}(exec)
	}
}

func pingExecutor(exec *executor) bool {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(net.JoinHostPort(exec.host, exec.port), opts...)
	if err != nil {
		log.WithError(err).Error("Router: could not create executor dial context")
		return false
	}
	defer conn.Close()

	client := service.NewCodyglotClient(conn)

	_, err = client.Ping(context.Background(), &service.Ping{})
	if err != nil {
		log.WithError(err).Errorf("Router: an error occured while pinging executor %s", exec.host)
		return false
	}

	return true
}

func getExecutorLanguages(executor string) ([]string, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(executor, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Router: could not create executor dial context")
	}
	defer conn.Close()

	client := service.NewCodyglotClient(conn)

	res, err := client.Languages(context.Background(), &service.LanguagesRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Router: an error occured while calling executor")
	}

	return res.Languages, nil
}
