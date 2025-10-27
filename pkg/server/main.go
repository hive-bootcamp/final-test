package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/sqwdvrt/final-test/pkg/db" // путь к твоему пакету db
)

func Run() error {
	// Инициализация базы данных
	dbFile := "scheduler.db"
	if err := db.Init(dbFile); err != nil {
		return fmt.Errorf("не удалось инициализировать базу: %w", err)
	}

	// Определение порта
	port := 7540
	strPort := os.Getenv("TODO_PORT")
	if strPort != "" {
		if p, err := strconv.Atoi(strPort); err == nil {
			port = p
		}
	}

	http.Handle("/", http.FileServer(http.Dir("web")))
	fmt.Printf("Server started on port %d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {
	if err := Run(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

	log.Println("База данных успешно инициализирована")
}
