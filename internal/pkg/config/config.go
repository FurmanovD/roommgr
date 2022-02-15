package config

import (
	"flag"

	"github.com/peterbourgon/ff"
	"github.com/peterbourgon/ff/fftoml"

	"github.com/FurmanovD/roommgr/internal/app/service"
	"github.com/FurmanovD/roommgr/pkg/commoncfg"
)

const (
	// ConfigPrefix used in reading the base constants from ini file
	ConfigPrefix = "roommgr."

	// Default INI file to use
	DefaultConfigIni = "/app/config.ini"

	// default port to use
	DefaultAPIPort = 8080
	// default log level to use
	DefaultLogLevel = "info"

	// default DB maxConnections to use
	DefaultDBMaxConnections = 10
)

// Config stuct for the Station API
type Config struct {
	APIPort  int
	LogLevel string

	SQLConfig commoncfg.SqlDBConfig
	Service   service.ServiceConfig
}

// ParseConfig parses the configuration file
func ParseConfig(args []string) (*Config, error) {

	fs := flag.NewFlagSet("roommgr", flag.ContinueOnError)

	cfg := &Config{}

	// initialize command-line parameters:
	fs.IntVar(&cfg.APIPort, "port", DefaultAPIPort, "this is the port for Station API to serve(8080 by default)")
	fs.StringVar(
		&cfg.LogLevel,
		"loglevel",
		"info",
		"Use this option to set the log level for the application. "+
			"Possible Values: trace, debug, info, warn, error, fatal, panic, nolog",
	)
	// end of command-line parameters

	initServiceConfig(fs, &cfg.Service)

	initDBConfig(fs, &cfg.SQLConfig)

	_ = fs.String("config", DefaultConfigIni, "config ini")
	err := ff.Parse(fs, args,
		ff.WithEnvVarNoPrefix(),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(fftoml.Parser),
		ff.WithIgnoreUndefined(true),
	)

	return cfg, err
}

// initDBConfig parses the DB configuration parameters
func initDBConfig(fs *flag.FlagSet, cfg *commoncfg.SqlDBConfig) {

	//  MySQL config
	fs.StringVar(&cfg.Host, ConfigPrefix+"DB_HOST", "", "host name of mysql db instance")
	fs.IntVar(&cfg.Port, ConfigPrefix+"DB_PORT", commoncfg.DefaultMySQLPort, "port number name of mysql db instance")
	fs.StringVar(&cfg.Database, ConfigPrefix+"DB_NAME", "", "write instance database name")
	fs.StringVar(&cfg.User, ConfigPrefix+"DB_USER", "", "mysql user name for db")
	fs.StringVar(&cfg.Password, ConfigPrefix+"DB_PASS", "", "password for mysql user for db")
	fs.IntVar(&cfg.MaxConnections, ConfigPrefix+"DB_MAX_CONNECTIONS", DefaultDBMaxConnections, "max conenctions to a DB")
}

// initServiceConfig parses the StaionService configuration parameters
func initServiceConfig(fs *flag.FlagSet, cfg *service.ServiceConfig) {
	// TODO!!!
}
