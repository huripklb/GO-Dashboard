package main

    import (
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
        //"strings"
        "time"
    )

    func main() {
    	//ch := make(chan int, 4)
        //go getOrderAge("pending", ch)
        //go getOrderAge("proceed", ch)
        //go getOrderAge("shipped", ch)
        //go getOrderAge("delivered", ch)
        go getOrderAge("pending")
        go getOrderAge("proceed")
        go getOrderAge("shipped")
        go getOrderAge("delivered")

        time.Sleep(2 * time.Second)
    }

    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }

    func getOrderAge(status string/*, ch chan int*/) {
    	db, err := sql.Open("mysql", "root:huripdb@/newkalbe?charset=utf8")
        checkErr(err)

    	// query
    	query := `select count(khdex_order_id) as total_orders, 
datediff(date_pending, now()) as date_diff
FROM khdex_order_go 
WHERE status = '%s' 
AND date_pending IS NOT NULL 
GROUP BY date_diff`

        rows, err := db.Query(fmt.Sprintf(query, status))
        checkErr(err)

        i := 0
        for rows.Next() {
            var total_orders int
            var date_diff int
            err = rows.Scan(&total_orders, &date_diff)
            checkErr(err)
            go insert(status, total_orders, (date_diff * -1))
            //fmt.Printf("%d. pending : %d %d\n", i, total_orders, (date_diff * -1))
            i++
        }

        //time.Sleep(2 * time.Second)

        //close(ch)
    }

    func insert(key string, total_orders int, date_diff int) {
    	db, err := sql.Open("mysql", "root:huripdb@/newkalbe?charset=utf8")
        checkErr(err)

        // insert
        stmt, err := db.Prepare("INSERT go_dashboard SET go_dashboard_key=?,value_1=?,value_2=?,created_at=?")
        checkErr(err)

        created_at := time.Now()
        stmt.Exec(key, total_orders, date_diff, created_at)

        //time.Sleep(2 * time.Second)
    }