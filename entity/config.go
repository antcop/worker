package entity

type Config struct {
	User string
	Password string
	Name string
	Description string
	Bind string
	Port int
	JobConcurrency int
	MaxWorker int
	MemoryLimit int
	RedisHost string
	RedisPort int
	DatabasePath string
}

func LoadConfigFile(filePath string) (Config, error) {
	config := Config {
	}
	return config, nil
}