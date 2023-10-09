package routes

import (
	"database/sql"
	"fmt"
)

func Update_user() {
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
	user := "UPDATE conductores set nombre=?, edad=? where id=?"
	actualiar, err := db.Exec(user, "lucas", 56, 2)
	if err != nil {
		panic(err)
	}
	num, err := actualiar.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("actualizado  - filas fectada ", num)

}
