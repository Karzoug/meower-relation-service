package neo4j

type Config struct {
	URI      string `env:"URI,notEmpty"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
}
