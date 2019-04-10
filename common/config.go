package common

import (
	"github.com/kelseyhightower/envconfig"
)

type CoreConfig struct {
	owner int64
	b_ID int64
	
	cl_ID string
	cl_secret string
	b_token string
	host string
	email string
}

func LoadConfig() (c *CoreConfig, err error) {
	c = &CoreConfig {}
	err = envconfig.Process("RoomBot", c)
	return 
}