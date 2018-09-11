package main

    import (
        "fmt"
        "time"
        "github.com/kr/beanstalk"
    )

    var conn, _ = beanstalk.Dial("tcp", "localhost:11300")

    func main() {
        /*id, err := conn.Put([]byte("myjob2"), 1, 0, time.Minute)
        if err != nil {
            panic(err)
        }
        fmt.Println("job", id)*/


        id, body, err := conn.Reserve(5 * time.Second)
        if err != nil {
            panic(err)
        }
        fmt.Println("job", id)
        fmt.Println(string(body))
    }