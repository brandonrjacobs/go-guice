package logging

import (
	"github.com/spf13/viper"
	go_guice "go-guice"
	"go.uber.org/zap"
)

func New(cfg *viper.Viper) (*zap.Logger, error) {
	zconfig := zapConfig(cfg)
	zlog, _ := zconfig.Build()
	return zlog, nil
}

func NewStartUp(cfg *viper.Viper) go_guice.StartupLog {
	zlog, _ := New(cfg)
	zlog = zlog.Named("startup")
	zWrapper := zapLog{
		logger:  zlog,
		sugared: zlog.Sugar(),
	}
	return &zWrapper
}

func NewBackgroundLog(cfg *viper.Viper) go_guice.BackgroundLog {
	zlog, _ := New(cfg)
	zlog = zlog.Named("background")
	zWrapper := zapLog{
		logger:  zlog,
		sugared: zlog.Sugar(),
	}
	return &background{zLog: &zWrapper}
}

func NewRequestLog(cfg *viper.Viper) go_guice.RequestLog {
	zlog, _ := New(cfg)
	zlog = zlog.Named("request")
	zWrapper := &zapLog{
		logger:  zlog,
		sugared: zlog.Sugar(),
	}
	return &request{zLog: zWrapper}
}
