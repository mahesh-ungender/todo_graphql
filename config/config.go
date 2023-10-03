package config

type config struct {
	postgresConfig      postgres
}

var configuration = &config{}

// Load loads the config into the configuration object
func Load() {
	configuration.postgresConfig.load()
}

// Postgres returns the postgres config
func Postgres() postgres {
	return configuration.postgresConfig
}

