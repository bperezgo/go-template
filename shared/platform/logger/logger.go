package logger

import (
	"github.com/rs/zerolog/log"
)

type LogHttpInput struct {
	Request *struct {
		Body   interface{}
		Params interface{}
		Query  interface{}
	}
	Response *struct {
		Data interface{}
	}
}

type LogInput struct {
	Action  string
	State   string
	Message string
	Http    *LogHttpInput
}

type Logger struct{}

var logger *Logger

func InitLogger() {
	if logger == nil {
		logger = &Logger{}
	}
}

func GetLogger() *Logger {
	InitLogger()
	return logger
}

func (Logger) Info(input LogInput) {
	if e := log.Info(); e.Enabled() {
		e.Str("action", input.Action).Str("action", input.State).Msg(input.Message)
	}
}

func (Logger) Error(input LogInput) {
	if e := log.Error(); e.Enabled() {
		e.Str("action", input.Action).Str("action", input.State).Msg(input.Message)
	}
}

func (Logger) Warn(input LogInput) {
	if e := log.Warn(); e.Enabled() {
		e.Str("action", input.Action).Str("action", input.State).Msg(input.Message)
	}
}
