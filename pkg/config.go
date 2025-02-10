package pkg

import "fmt"

type GrpcServerConfig struct {
	Host string
	Port string
}

func (c GrpcServerConfig) Url() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
