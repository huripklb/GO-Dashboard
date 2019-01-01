package dbconn

import (
	"fmt"

	rc "readconf"

	"github.com/jinzhu/gorm"
	// use gorm postgre wrapper
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DatabaseConfig : struct config to contain DB configuration
type DatabaseConfig struct {
	Dbtype string `mapstructure:"dbtype"`
	Host   string `mapstructure:"hostname"`
	Port   string `mapstructure:"port"`
	User   string `mapstructure:"username"`
	Pass   string `mapstructure:"password"`
	Dbname string `mapstructure:"dbname"`
}

// Config : Wrapper for DB Config
type Config struct {
	Db DatabaseConfig `mapstructure:"postgredb"`
}

// PostgreConn : Connect to postgre database server
// author : Huripto Sugandi
// created date : 4 Dec 2018
func PostgreConn() (*gorm.DB, error) {
	v := rc.ReadConfig()
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		c.Db.User, c.Db.Pass, c.Db.Dbname)
	//db, err := sql.Open(c.Db.Dbtype, dbinfo)
	db, err := gorm.Open("postgres", dbinfo)

	checkErr(err)

	return db, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
