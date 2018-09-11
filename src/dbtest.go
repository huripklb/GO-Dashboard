package main

    import (
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
        "time"
    )

    func main() {
        db, err := sql.Open("mysql", "root:huripdb@/golang?charset=utf8")
        checkErr(err)

        // insert
        stmt, err := db.Prepare("INSERT golang_test SET golang_test_text=?,golang_test_int=?,created_at=?")
        checkErr(err)

        created_at := time.Now()
        /*created_at := created_at.Format("2006-01-02 15:04:05")*/

        res, err := stmt.Exec("huripto sugandi", "12", created_at)
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        fmt.Println(id)
    }

    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }