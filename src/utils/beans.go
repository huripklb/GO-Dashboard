package utils

import (
	"fmt"
	"log"
	"time"

	rc "readconf"

	"github.com/kr/beanstalk"
)

type BeanstalkConfig struct {
	Beansurl string `mapstructure:"beansurl"`
}

type Config struct {
	Bu BeanstalkConfig `mapstructure:"beanstalk"`
}

// PutBeans : Put bytes of message to Beanstalk queue
func PutBeans(tubename string, message []byte) uint64 {
	v := rc.ReadConfig()
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}
	var conn, _ = beanstalk.Dial("tcp", c.Bu.Beansurl)
	tube := &beanstalk.Tube{conn, tubename}
	id, err := tube.Put(message, 1, 0, time.Minute)
	if err != nil {
		log.Println(err)
	}

	return id
}

// ReserveBeans : Reserve beanstalk queue data
func ReserveBeans(tubename string) []byte {
	v := rc.ReadConfig()
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}
	var conn, _ = beanstalk.Dial("tcp", c.Bu.Beansurl)
	tubeSet := beanstalk.NewTubeSet(conn, tubename)

	id, body, err := tubeSet.Reserve(5 * time.Hour)
	if err != nil {
		log.Println(err)
	}

	conn.Delete(id)
	return body
}
