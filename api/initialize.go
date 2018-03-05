package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Port   string
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	//a.Router.HandleFunc("/users", getUsers).Methods("GET")
	a.Router.HandleFunc("/user", a.CreateUser).Methods("POST")
}

func (a *App) Run(addr string) {
	fmt.Println("GoRestExample starting on :: ", a.GetPort())
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("GORestExample running on port %s :: ", a.Port)
	}
}

func (a *App) GetPort() string {
	a.Port = os.Getenv("PORT")
	if a.Port == "" {
		a.Port = "1340"
	}
	return a.Port
}
