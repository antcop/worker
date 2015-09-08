package entity

type Config struct {
	JobConcurrency int
	MaxWorker int
	MemoryLimit int
	RedisHost string
	RedisPort string
	StorageFile string
}

func LoadConfigFile(filePath string) (Config, error) {
	config := Config {
	}
	return config, nil
}