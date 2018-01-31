package codyglot

import (
	log "github.com/sirupsen/logrus"
)

type logLevelValue struct {
	log.Level
}

func (lv *logLevelValue) Set(v string) error {
	level, err := log.ParseLevel(v)
	if err != nil {
		return err
	}
	lv.Level = level
	return nil
}

func (lv *logLevelValue) String() string {
	return lv.Level.String()
}

func (lv *logLevelValue) Type() string {
	return "string"
}
