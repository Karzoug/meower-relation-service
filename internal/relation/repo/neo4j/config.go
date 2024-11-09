package neo4j

type Config struct {
	DBName string `env:"DB_NAME" envDefault:"meower-realations"`
}
