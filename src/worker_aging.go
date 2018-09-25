package main

    import (
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
        "time"
        "sync"
    )

    func main() {
    	var wg sync.WaitGroup
    	db, err := sql.Open("mysql", "root:huripdb@/newkalbe?charset=utf8")
    	checkErr(err)

    	// clear the table
    	clear(db)
    	// don't use go routine. this process must finish first

    	wg.Add(6)

    	// get orders age
        go getOrderAge(&wg, db, "pending")
        go getOrderAge(&wg, db, "proceed")
        go getOrderAge(&wg, db, "shipped")
        go getOrderAge(&wg, db, "delivered")

        // get PO age
        go getPoAge(&wg, db, "draft")
        go getPoAge(&wg, db, "shipped")
        //go getPoAge(db, "completed")

        // make sure the channel doesn't close while executing go routine
        // other method available, will be done in next research
        //time.Sleep(3 * time.Second)
        wg.Wait()
    }

    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }

    func clear(db *sql.DB) {
    	stmt, err := db.Prepare("TRUNCATE go_dashboard")
        checkErr(err)

        stmt.Exec()
    }

    func getOrderAge(wg *sync.WaitGroup, db *sql.DB, status string/*, ch chan int*/) {
    	// query to get order age by pending date, based on order status
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
            insert(db, status, total_orders, (date_diff * -1))
            //fmt.Printf("%d. %s : %d %d\n", i, status, total_orders, (date_diff * -1))
            i++
        }

        defer wg.Done()
    }

    func getPoAge(wg *sync.WaitGroup, db *sql.DB, status string) {
    	// query to get order age by pending date, based on order status
    	query := `SELECT COUNT(stockist_po_id) as total_po, datediff(created_at, now()) as date_diff 
FROM stockist_po WHERE status = '%s'
GROUP BY date_diff`

        rows, err := db.Query(fmt.Sprintf(query, status))
        checkErr(err)

        i := 0
        for rows.Next() {
            var total_po int
            var date_diff int
            err = rows.Scan(&total_po, &date_diff)
            checkErr(err)
            dashboard_key := fmt.Sprintf(`po_%s`, status)
            insert(db, dashboard_key, total_po, (date_diff * -1))
            //fmt.Printf("%d. %s : %d %d\n", i, dashboard_key, total_po, (date_diff * -1))
            i++
        }

        defer wg.Done()
    }

    func insert(db *sql.DB, key string, total_orders int, date_diff int) {
    	//db, err := sql.Open("mysql", "root:huripdb@/newkalbe?charset=utf8")
        //checkErr(err)

        // insert
        stmt, err := db.Prepare("INSERT go_dashboard SET go_dashboard_key=?,value_1=?,value_2=?,created_at=?")
        checkErr(err)

        created_at := time.Now()
        stmt.Exec(key, total_orders, date_diff, created_at)

        //defer wg.Done()
    }