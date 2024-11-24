package config

type Config struct {
	MaxRequest  string `env:"MAX_REQUEST" envDefault:"2"`
	BucketToken string `env:"TOKEN_BUCKET_RATE_PER_SECOND" envDefault:"1"`
	HttpPort    string `env:"PORT" envDefault:"8080"`
}

func InitConfig() *Config {
	return &Config{}
}
