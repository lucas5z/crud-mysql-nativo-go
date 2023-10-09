package routes

import (
	"database/sql"
	"fmt"
)

func Read_user() {
	dbUser := "root"
	dbPass := ""
	dbName := "unajmabus"
	dbHost := "localhost" // Cambia esto si tu base de datos está en un servidor remoto
	dbPort := "3306"      // Cambia esto si tu base de datos utiliza un puerto diferente

	// Crea la cadena de conexión
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	user := "select id,nombre,edad from conductores"

	read, err := db.Query(user)
	if err != nil {
		panic(err)
	}
	for read.Next() {
		var id int
		var nombre string
		var edad int
		err = read.Scan(&id, &nombre, &edad)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID:%d, Nombre:%s, edad:%d\n", id, nombre, edad)
	}

}
