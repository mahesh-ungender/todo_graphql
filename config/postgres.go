package config

import (
	"todo_graphql/logger"
)

// postgres is the config for the db vars
type postgres struct {
	host         string
	port         string
	user         string
	userPassword string
	database     string
	connURL      string
}

// load loads the config for the postgresdb
func (postgresConfig *postgres) load() {
	logger.Log.Info("Reading postgres config...")

	postgresConfig.host = "localhost"
	postgresConfig.port = "5432"
	postgresConfig.user = "postgres"
	postgresConfig.userPassword = "123"
	postgresConfig.database = "graph_beego"
	postgresConfig.connURL = "postgres://postgres:123@localhost:5432/todo_beego?sslmode=disable"
}

// Host returns the postgres host
func (postgresConfig *postgres) Host() string {
	return postgresConfig.host
}

// Port returns the postgres Port
func (postgresConfig *postgres) Port() string {
	return postgresConfig.port
}

// User returns the postgres user
func (postgresConfig *postgres) User() string {
	return postgresConfig.user
}

// UserPassword returns the postgres userPassword
func (postgresConfig *postgres) UserPassword() string {
	return postgresConfig.userPassword
}

// Database returns the postgres Database
func (postgresConfig *postgres) Database() string {
	return postgresConfig.database
}

// ConnURL returns the connection url
func (postgresConfig *postgres) ConnURL() string {
	return postgresConfig.connURL
}
