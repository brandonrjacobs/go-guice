package config

import (
	_ "embed"
	"github.com/spf13/viper"
	"os"
)

type Loader interface {
	LoadConfig(cfg *viper.Viper)
}

type simpleFileLoader struct {
	fileType FileType
	filesPaths []string
}

func (s simpleFileLoader) LoadConfig(cfg *viper.Viper){
	cfg.SetConfigType(s.fileType.ToViper())
	for _, f := range s.filesPaths {
		if fd, err := os.Open(f); err != nil {
			panic(any(err))
		} else {
			if err = cfg.ReadConfig(fd); err != nil {
				panic(any(err))
			}
		}
	}
}
