package main

import (
	"fmt"
	"log"
	"time"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println("This is the first time I write a go program")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		taosDriverName := "taosSql"
		// demodb := "test"
		// demot := "demot"

		fmt.Printf("\n======== start demo test ========\n")
		// open connect to taos server
		db, err := sql.Open(taosDriverName, "root:taosdata@/tcp(127.0.0.1:0)/")
		if err != nil {
			log.Fatalf("Open database error: %s\n", err)
		}
		defer db.Close()

		select_data(db, "t154")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func select_data(db *sql.DB, demot string) {
    st := time.Now().Nanosecond()

    rows, err := db.Query("select * from ? " , demot)  // go text mode
    checkErr(err, "select db.Query")

    fmt.Printf("%10s%s%8s %5s %9s%s %s %8s%s %7s%s %8s%s %4s%s %5s%s\n", " ","ts", " ", "id"," ", "name"," ","len", " ","flag"," ", "notes", " ", "fv", " ", " ", "dv")
    var affectd int
    for rows.Next() {
        var ts string
        var name string
        var id int
        var len int8
        var flag bool
        var notes string
        var fv float32
        var dv float64

        err = rows.Scan(&ts, &id, &name, &len, &flag, &notes, &fv, &dv)
        checkErr(err, "select rows.Scan")

        fmt.Printf("%s\t", ts)
        fmt.Printf("%d\t",id)
        fmt.Printf("%10s\t",name)
        fmt.Printf("%d\t",len)
        fmt.Printf("%t\t",flag)
        fmt.Printf("%s\t",notes)
        fmt.Printf("%06.3f\t",fv)
        fmt.Printf("%09.6f\n",dv)

        affectd++
    }

    et := time.Now().Nanosecond()
    fmt.Printf("select data result:\n %d row(s) affectd (%6.6fs)\n\n", affectd, (float32(et-st))/1E9)
}

func checkErr(err error, prompt string) {
    if err != nil {
	fmt.Printf("%s\n", prompt)
        panic(err)
    }
}