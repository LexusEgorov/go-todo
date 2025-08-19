package config

import (
	"errors"
	"flag"
)

var ErrMigrationsNotProvided = errors.New("migrations path didn't provide")

type MigratorConfig struct {
	DBConfig
	MigrationsPath string
}

func NewMigratorConfig() (*MigratorConfig, error) {
	var (
		migrationsPath string
		user           string
		password       string
		name           string
	)

	flag.StringVar(&migrationsPath, "m", "", "path to migrations")
	flag.StringVar(&password, "p", "", "db password")
	flag.StringVar(&name, "n", "", "db name")
	flag.StringVar(&user, "u", "", "db username")
	flag.Parse()

	if migrationsPath == "" {
		return nil, ErrMigrationsNotProvided
	}

	if user == "" {
		return nil, ErrBadUserName
	}

	if password == "" {
		return nil, ErrBadPassword
	}

	if name == "" {
		return nil, ErrBadDBName
	}

	return &MigratorConfig{
		MigrationsPath: migrationsPath,
		DBConfig: DBConfig{
			User:     user,
			Password: password,
			Name:     name,
		},
	}, nil
}
