package app

import "go-guice/config"

func ContainerCommand(use string, short, long string, writer HelpWriter) Command {
	return newCommand(use, short, long, true, nil, writer, []config.EnvVar{})
}
