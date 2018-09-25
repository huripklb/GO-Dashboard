package main

    import (
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
        "time"
        "encoding/json"
        "sync"
    )

    func main() {
        var wg sync.WaitGroup
    	db, err := sql.Open("mysql", "root:huripdb@/newkalbe?charset=utf8")
    	checkErr(err)

    	// get topten orders
        go topTenOrder(&wg, db, "pending")
        go topTenOrder(&wg, db, "proceed")
        go topTenOrder(&wg, db, "shipped")

        // get top ten PO
        go topTenPo(&wg, db, "draft")
        go topTenPo(&wg, db, "shipped")
        
        // make sure the channel doesn't close while executing go routine
        // other method available, will be done in next research
        time.Sleep(3 * time.Second)
    }

    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }

    func topTenOrder(wg *sync.WaitGroup, db *sql.DB, status string) {
        query := `SELECT khdex_order_no, khdex_code, date_pending, datediff(date_pending, now()) as date_diff
FROM khdex_order_go 
WHERE status = '%s' 
AND date_pending IS NOT NULL 
ORDER BY date_pending ASC
LIMIT 0,10`

        rows, err := db.Query(fmt.Sprintf(query, status))
        checkErr(err)

        i := 0
        result := make(map[string]string)
        for rows.Next() {
            var khdex_order_no string
            var khdex_code string
            var date_pending string
            var date_diff string
            err = rows.Scan(&khdex_order_no, &khdex_code, &date_pending, &date_diff)
            checkErr(err)
            result["khdex_order_no"] = khdex_order_no
            result["khdex_code"] = khdex_code
            result["date_pending"] = date_pending
            result["date_diff"] = date_diff
            slcB, _ := json.Marshal(result)
            //fmt.Println(string(slcB))
            dashboard_key := fmt.Sprintf(`topten_order_%s`, status)
            insert(db, dashboard_key, i, slcB)
            //fmt.Printf("%d. %s : %d %d\n", i, status, total_orders, (date_diff * -1))
            i++
        }
    }

    func topTenPo(wg *sync.WaitGroup, db *sql.DB, status string) {
        query := `SELECT stockist_po_no, khdex_code, stockist_code, created_at, datediff(created_at, now()) as date_diff
FROM stockist_po 
WHERE status = '%s' 
ORDER BY created_at ASC
LIMIT 0,10`

        rows, err := db.Query(fmt.Sprintf(query, status))
        checkErr(err)

        i := 0
        result := make(map[string]string)
        for rows.Next() {
            var stockist_po_no string
            var khdex_code string
            var stockist_code string
            var created_at string
            var date_diff string
            err = rows.Scan(&stockist_po_no, &khdex_code, &stockist_code, &created_at, &date_diff)
            checkErr(err)
            result["stockist_po_no"] = stockist_po_no
            result["khdex_code"] = khdex_code
            result["stockist_code"] = stockist_code
            result["created_at"] = created_at
            result["date_diff"] = date_diff
            slcB, _ := json.Marshal(result)
            //fmt.Println(string(slcB))
            dashboard_key := fmt.Sprintf(`topten_po_%s`, status)
            insert(db, dashboard_key, i, slcB)
            //fmt.Printf("%d. %s : %d %d\n", i, status, total_orders, (date_diff * -1))
            i++
        }
    }

    func insert(db *sql.DB, key string, total_orders int, json_data []byte) {
    	//db, err := sql.Open("mysql", "root:huripdb@/newkalbe?charset=utf8")
        //checkErr(err)

        // insert
        stmt, err := db.Prepare("INSERT go_dashboard SET go_dashboard_key=?,value_1=?,value_2=?,created_at=?")
        checkErr(err)

        created_at := time.Now()
        stmt.Exec(key, total_orders, json_data, created_at)
    }