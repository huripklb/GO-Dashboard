package testfolder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	ut "utils"
)

// Dummy : struct type for dummy data
type Dummy struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    int    `json:"age"`
}

// Dummy2 : struct type for dummy data
type Dummy2 struct {
	Name         string `json:"employee_name"`
	Salary       string `json:"employee_salary"`
	Age          string `json:"employee_age"`
	ProfileImage string `json:"profile_image"`
}

func NewDummyData(name string, salary, age int) *Dummy {
	return &Dummy{
		Name:   name,
		Salary: salary,
		Age:    age,
	}
}

// FolderPage : Public function to test function call from external folder/package
func FolderPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sucessfulyy call function in a folder.")
}

// PostPage : Test API Call POST
func PostPage(w http.ResponseWriter, r *http.Request) {
	dummyData := NewDummyData("name test", 12345, 25)
	jsonDummy, err := json.Marshal(dummyData)
	if err != nil {
		log.Println(err)
	}
	log.Printf("JSON Sent : %s", jsonDummy)

	header := &ut.ReqHeader{ContentType: "Content-Type", Content: "application/json"}
	apiRes, status := ut.CallRequest("http://dummy.restapiexample.com/api/v1/create", jsonDummy, "POST", header)

	if status == false {
		fmt.Fprintf(w, "Request Not OK!")
		return
	}

	log.Printf("JSON Received : %s", apiRes)
	fmt.Fprintf(w, "%s", string(apiRes))
}

// GetPage : Test API Call GET
func GetPage(w http.ResponseWriter, r *http.Request) {
	apiRes, status := ut.CallRequest("http://dummy.restapiexample.com/api/v1/employees", nil, "GET", nil)

	if status == false {
		fmt.Fprintf(w, "Request Not OK!")
		return
	}

	var data []Dummy2
	var counter int
	json.Unmarshal(apiRes, &data)
	counter = 0
	fmt.Fprintf(w, "%30v | %v | %v\n", "Name", "Age", "Salary")
	for k := range data {
		fmt.Fprintf(w, "%30v | %v | %v\n", data[k].Name, data[k].Age, data[k].Salary)
		counter++
		if counter == 10 {
			return
		}
	}
	log.Printf("JSON Received : %s", apiRes)
}

// TestPutBeans : test put to beanstalk queue
func TestPutBeans(w http.ResponseWriter, r *http.Request) {
	message := []byte("hollaaa")
	tubename := "go_tube"

	log.Println("Beanstalk access start")
	id := ut.PutBeans(tubename, message)

	fmt.Fprintf(w, "Beans put id : %d", id)
}

// TestReserveBeans : test put to beanstalk queue
func TestReserveBeans(w http.ResponseWriter, r *http.Request) {
	tubename := "go_tube"

	log.Println("Beanstalk access start")
	body := ut.ReserveBeans(tubename)

	fmt.Fprintf(w, "Beans put body content : %s", body)
}
