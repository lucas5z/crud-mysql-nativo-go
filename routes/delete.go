package routes

import (
	"database/sql"
	"fmt"
)

func Delete_user() {
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

	user := "delete from conductores where id=?"
	delete, err := db.Exec(user, 3)
	if err != nil {
		panic(err)
	}
	num, err := delete.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("eliminado - filas afectadas", num)

}
