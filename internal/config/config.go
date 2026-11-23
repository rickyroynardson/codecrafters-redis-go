package config

import "flag"

type Config struct {
	Port int
}

func Init() *Config {
	port := flag.Int("port", 6379, "TCP port to listen on")

	flag.Parse()

	return &Config{
		Port: *port,
	}
}
