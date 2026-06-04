package server

type ServerConfig struct {
	Host string
	Port uint16

	LogRequests      bool
	DefaultRateLimit uint16 // request per minute
}
