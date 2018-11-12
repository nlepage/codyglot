package executor

import (
	"context"
	"net"
	"strconv"
	"time"

	executor_config "github.com/nlepage/codyglot/executor"
	"github.com/nlepage/codyglot/router/server/config"
	"github.com/nlepage/codyglot/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type executor struct {
	hostport   string
	aliveCh    chan bool
	backoff    time.Duration
	_languages []string
}

func newStatic(hostport string) *executor {
	host, port, err := net.SplitHostPort(hostport)
	if err != nil {
		host = hostport
		port = strconv.Itoa(executor_config.DefaultPort)
		log.Warnf("Invalid executor address \"%s\" (%s), using value as host and default port %s", hostport, err.Error(), port)
	}

	return &executor{
		hostport: net.JoinHostPort(host, port),
		aliveCh:  make(chan bool),
	}
}

func (exec *executor) start() {
	go exec.runPing()
	go exec.runLanguages()
}

func (exec *executor) runPing() {
	alive := false

	for {
		if alive {
			time.Sleep(config.PingInterval)
		} else {
			time.Sleep(exec.backoff * 100 * time.Millisecond)
		}

		if pong := exec.ping(); pong {
			if !alive {
				alive = true
				log.Infof("Executor %s is alive", exec.hostport)
				exec.aliveCh <- true
			}

			if exec.backoff != 0 {
				exec.backoff = 0
			}
		} else {
			if alive {
				alive = false
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

func (exec *executor) runLanguages() {
	alive := false

	for {
		select {
		case alive = <-exec.aliveCh:
			if alive {
				time.Sleep(100 * time.Millisecond)
				exec.languages()
			} else {
				languagesCh <- execLanguages{exec, exec._languages, true}
			}
		case <-time.After(time.Minute): //FIXME config
			if alive {
				exec.languages()
			}
		}
	}
}

func (exec *executor) languages() {
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

	exec._languages = res.Languages
	languagesCh <- execLanguages{exec, exec._languages, false}
}

func (exec *executor) execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(exec.hostport, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Router: could not create executor dial context")
	}
	defer conn.Close()

	client := service.NewCodyglotClient(conn)

	res, err := client.Execute(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Router: an error occured while calling executor")
	}

	return res, nil
}
