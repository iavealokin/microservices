package apiserver

//Config struct ..
type Config struct {
	BindAddr 	string `toml:"bind_addr"`
	LogLevel 	string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	RmqURL	 	string `toml:"rmq_url"`
	QueueName	string `toml:"rmq_queue"`
}
//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8082",
		LogLevel: "debug",
	}
}