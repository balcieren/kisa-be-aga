package logger

import "go.uber.org/zap"

type Logger = zap.SugaredLogger

func New() (*Logger, error) {
	l, err := zap.NewProduction()
	sugar := l.Sugar()
	return sugar, err
}
