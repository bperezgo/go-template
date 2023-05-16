package logger

import (
	"encoding/json"

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

type Error struct {
	Message string
}

type LogInput struct {
	Action  string
	State   string
	Message string
	Http    *LogHttpInput
	Error   *Error
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
		b, err := json.Marshal(input.Http)
		if err != nil {
			e.Str("action", input.Action).Str("action", input.State).Msg(err.Error())
			return
		}
		e.Str("action", input.Action).
			Str("action", input.State).
			RawJSON("http", b).
			Msg(input.Message)
	}
}

func (Logger) Error(input LogInput) {
	if e := log.Error(); e.Enabled() {
		e.Str("action", input.Action).Str("action", input.State).Msg(input.Message)
		e.Any("http", input.Http)
	}
}

func (Logger) Warn(input LogInput) {
	if e := log.Warn(); e.Enabled() {
		e.Str("action", input.Action).Str("action", input.State).Msg(input.Message)
		e.Any("http", input.Http)
	}
}
