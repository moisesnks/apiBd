package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	host := "172.17.0.3"
	port := 5432
	user := "postgres"
	password := "123456"
	dbname := "bdtest"

	// Cadena de conexión de PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Abre una conexión a la base de datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Intenta realizar alguna operación con la base de datos para verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos.")

	// Configura las rutas con Gorilla Mux
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Conexión exitosa a la base de datos."))
	}).Methods("GET")

	serverPort := 9090 // Cambia el puerto según tus necesidades
	serverAddr := fmt.Sprintf(":%d", serverPort)
	log.Printf("Servidor escuchando en %s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
