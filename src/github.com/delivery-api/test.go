package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	rc "readconf"

	"github.com/delivery-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kr/beanstalk"
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

type BeanstalkConfig struct {
	Beansurl string `mapstructure:"beansurl"`
}

type Config struct {
	Db DatabaseConfig  `mapstructure:"postgredb"`
	Bu BeanstalkConfig `mapstructure:"beanstalk"`
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Yuuhuu.. this is the homepage.")
}

func AnotherPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Yuuhhuu... This is another page.")
}

func TestCall() {
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

	data := []byte("huripto-test")
	password := fmt.Sprintf("%x", md5.Sum(data))
	username := "huripto-local"

	checkUserAPI(db, username, password)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkUserAPI(db *gorm.DB, username string, password string) {

	//db.Create(&model.Apiuser{Username: "huripto-local", Password: "a69f39ced2429be6de7b53afb1700019", Active: true})
	//db.Create(&model.Apiuser{Username: "huripto-testing", Password: "a69f39ced2429be6de7b53afb1700019", Active: true})

	var users []model.Apiuser
	db.Where("username = ? AND password = ?", username, password).First(&users).RecordNotFound()

	if len(users) == 0 {
		fmt.Println("Invalid password or user not found")
		return
	}

	fmt.Println("api_user_id | username             | status")
	for k := range users {
		fmt.Printf("%11v | %20v | %v\n", users[k].ID, users[k].Username, users[k].Active)
	}
	jsonUsers, _ := json.Marshal(users)
	fmt.Println(string(jsonUsers))
}

func TestPutBeans(tubename string, message []byte) {
	v := rc.ReadConfig()
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}
	var conn, _ = beanstalk.Dial("tcp", c.Bu.Beansurl)
	tube := &beanstalk.Tube{conn, tubename}
	id, err := tube.Put([]byte("hollaa"), 1, 0, time.Minute)
	if err != nil {
		panic(err)
	}
	fmt.Println("job", id)
}
