package main

import (
	"fmt"
	"log"
	"net/http"
	"testProj/src/common"
	"testProj/src/controllers"
	database "testProj/src/db"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	if initServer() != nil {
		return
	}
	defer database.Db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/api/event", controllers.HandleEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8101", router))
	logrus.Info("старт сервера")
}

func initServer() error {
	var err error

	err = common.LoadConfig()
	if err != nil {
		logrus.Error(err)
		return err
	}

	db := &database.ClickHouseDB{}
	err = db.Init()

	if err != nil {
		logrus.Error("ошибка инициализации соединения с бд", err)
		return fmt.Errorf("ошибка инициализации соединения с бд")
	}

	return nil
}
