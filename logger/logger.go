package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(env string) {
	var err error
	if env == "dev" || env == "" {
		Log, err = zap.NewDevelopment()
	} else {
		Log, err = zap.NewProduction()
	}
	if err != nil {
		panic("cannot initialize zap logger: " + err.Error())
	}
}
