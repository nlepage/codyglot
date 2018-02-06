package server

import (
	"context"
	"net"
	"strconv"
	"time"

	executor_config "github.com/nlepage/codyglot/executor/config"
	"github.com/nlepage/codyglot/router/server/config"
	"github.com/nlepage/codyglot/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type executor struct {
	hostport  string
	alive     bool
	aliveCh   chan bool
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
			hostport: net.JoinHostPort(host, port),
			aliveCh:  make(chan bool),
		})
	}

	return nil
}

func startExecutors() {
	for _, exec := range executors {
		go exec.runPing()
	}
}

func (exec *executor) runPing() {
	for {
		if exec.alive {
			time.Sleep(config.PingInterval)
		} else {
			time.Sleep(exec.backoff * 100 * time.Millisecond)
		}

		if alive := exec.ping(); alive {
			if !exec.alive {
				exec.alive = true
				log.Infof("Executor %s is alive", exec.hostport)
				exec.aliveCh <- true
			}

			if exec.backoff != 0 {
				exec.backoff = 0
			}
		} else {
			if exec.alive {
				exec.alive = false
				log.Infof("Executor %s is dead", exec.hostport)
				exec.aliveCh <- false
			}

			if exec.backoff == 0 {
				exec.backoff = 1
			} else if exec.backoff*100 < config.MaxBackoff {
				exec.backoff <<= 1
			}
		}
	}
}

func (exec *executor) ping() bool {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(exec.hostport, opts...)
	if err != nil {
		log.WithError(err).Error("Router: could not create executor dial context")
		return false
	}
	defer conn.Close()

	_, err = service.NewCodyglotClient(conn).Ping(context.Background(), &service.Ping{})
	if err != nil {
		log.WithError(err).Errorf("Router: an error occured while pinging executor %s", exec.hostport)
		return false
	}

	return true
}

func (exec *executor) runRefreshLanguages() {
	for {
		select {
		case alive := <-exec.aliveCh:
			if alive {
				exec.refreshLanguages()
			} else {

			}
		case <-time.After(time.Minute):
			// FIXME put a rwlock on alive ? on exec ?
			if exec.alive {
				exec.refreshLanguages()
			}
		}
	}
}

func (exec *executor) refreshLanguages() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(exec.hostport, opts...)
	if err != nil {
		log.WithError(err).Error("Router: could not create executor dial context")
		return
	}
	defer conn.Close()

	res, err := service.NewCodyglotClient(conn).Languages(context.Background(), &service.LanguagesRequest{})
	if err != nil {
		log.WithError(err).Error("Router: an error occured while calling executor")
		return
	}

	// FIXME put a rwlock on languages
	exec.languages = res.Languages

	// FIXME send notif to refresh executor map
}
