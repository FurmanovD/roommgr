package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/FurmanovD/roommgr/internal/app/service"
	"github.com/FurmanovD/roommgr/internal/app/webapp"
	"github.com/FurmanovD/roommgr/internal/pkg/config"
	"github.com/FurmanovD/roommgr/internal/pkg/db/apidbconvert/v1"
	"github.com/FurmanovD/roommgr/internal/pkg/db/repository"
	"github.com/FurmanovD/roommgr/pkg/commoncfg"
	"github.com/FurmanovD/roommgr/pkg/sqldb"
)

const (
	// program's exit codes
	errCodeConfigError       = 1
	errCodeDBConnectionError = 2
	// errCodeServiceInitError     = 3 // TODO!!!
	// errCodeSyncerInitError      = 4
	// errCodePrometheusStartError = 5
	errCodeWebSrvListenError = 9
)

// Build information
// The actual information will be stored when 'go build' is called from the Docker file
var (
	Version   = "local-dev"
	BuildTime = time.Now().Format(time.RFC3339)
	GitCommit = ""

	buildInfo = ""
)

func init() {
	buildInfo = fmt.Sprintf(
		"Version: %v BuildTime: %v GitCommit: %v",
		Version,
		BuildTime,
		GitCommit,
	)

	log.Println(buildInfo)
}

func main() {

	// Parse flags/config file to populate config
	cfg, err := config.ParseConfig(os.Args[1:])
	if err != nil {
		fmt.Printf("configuration read error: %+v", err)
		os.Exit(errCodeConfigError)
	}
	initLogging(cfg.LogLevel)
	log.Infof("Logger initialized with LogLevel: %v", cfg.LogLevel)

	log.Info(buildInfo)

	// create a DB connection
	log.Info("Creating a DB connection...")
	dbInstance, err := initDBConnection(sqldb.NewDB(), &cfg.SQLConfig)
	if err != nil {
		log.Errorf("Creating DB connection failed", err)
		os.Exit(errCodeDBConnectionError)
	}

	// Create a service instance that will do all required operations to DB, storages etc.
	log.Infof("Creating service instance")
	roommgrService := service.NewService(
		cfg.Service,
		repository.NewRepository(dbInstance),
		apidbconvert.NewAPIDBConverter(),
	)
	log.Infof("Starting a web server on port %d", cfg.APIPort)

	// create a web server instances to serve HTTP endpoints
	webServer := webapp.NewServer(roommgrService)
	webServer.RegisterRoutes()

	err = webServer.ListenAndServe(cfg.APIPort)
	if err != nil {
		log.Errorf("Web server start failed", err)
		os.Exit(errCodeWebSrvListenError)
	}
}

// initLogging establishes process logging level
func initLogging(logLevel string) {

	log.SetFormatter(&log.JSONFormatter{})
	// sets the logging level in app
	level, err := log.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

func initDBConnection(db sqldb.SqlDB, dbConfig *commoncfg.SqlDBConfig) (sqldb.SqlDB, error) {

	err := db.Connect(
		dbConfig.MySqlConnectionString(),
		dbConfig.MaxConnections,
	)
	if err != nil {
		return nil, fmt.Errorf("connecting read DB error: %w", err)
	}

	return db, nil
}
