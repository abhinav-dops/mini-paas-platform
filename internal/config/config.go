package config

import "os"

var SSHKeyPath = getEnv("SSH_KEY_PATH", "./mini-paas-key.pem")

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
