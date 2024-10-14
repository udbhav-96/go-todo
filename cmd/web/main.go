package main

import(
	// "fmt"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/udbhav-96/go-todo/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct{
		errorLog 	*log.Logger
		infoLog 	*log.Logger
		tasks		*models.TaskModel
	}

func main(){

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/gotodo?parseTime=true", "/gotodo")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	// database connection -------------------------------------------------
	db, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }

    defer db.Close()
    // ---------------------------------------------------------------------

	app := &application{
		errorLog: 	errorLog,
		infoLog:	infoLog,
		tasks: &models.TaskModel{DB: db},
	}

    srv := &http.Server{
    	Addr: *addr,
    	ErrorLog: errorLog,
    	Handler: app.routes(),
    }

	infoLog.Printf("Starting Server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)	

}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}