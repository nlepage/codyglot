package executor

import (
	"context"
	"fmt"
	"sync"

	"github.com/nlepage/codyglot/router/server/config"
	"github.com/nlepage/codyglot/service"
)

type execLanguages struct {
	exec      *executor
	languages []string
	delete    bool
}

var (
	executors     []*executor
	languagesCh   = make(chan execLanguages)
	languages     = make(map[string]*executor)
	languagesLock sync.RWMutex
)

func Init() {
	initExecutorsStatic()
	startExecutors()
	go updateLanguages()
}

func Languages() []string {
	languagesLock.RLock()
	languagesList := make([]string, 0, len(languages))
	for language := range languages {
		languagesList = append(languagesList, language)
	}
	languagesLock.RUnlock()
	return languagesList
}

func Execute(ctx context.Context, req *service.ExecuteRequest) (*service.ExecuteResponse, error) {
	languagesLock.RLock()
	exec, ok := languages[req.Language]
	languagesLock.RUnlock()
	if !ok {
		return nil, fmt.Errorf("No executor for language %s", req.Language)
	}

	return exec.execute(ctx, req)
}

func initExecutorsStatic() {
	for _, hostport := range config.Executors {
		executors = append(executors, newStatic(hostport))
	}
}

func startExecutors() {
	for _, exec := range executors {
		exec.start()
	}
}

func updateLanguages() {
	for execLanguages := range languagesCh {
		languagesLock.Lock()
		if execLanguages.delete {
			for _, language := range execLanguages.languages {
				delete(languages, language)
			}
		} else {
			for _, language := range execLanguages.languages {
				languages[language] = execLanguages.exec
			}
		}
		languagesLock.Unlock()
	}
}
