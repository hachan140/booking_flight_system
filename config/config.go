package config

type Configurations struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"secretKey"`
}
