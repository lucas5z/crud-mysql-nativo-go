package routes

import (
	"database/sql"
	"fmt"
)

func Crea_user() {
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

	user := "INSERT INTO conductores(nombre,edad) values(?,?)"

	llenado, err := db.Exec(user, "juan", 42)
	if err != nil {
		panic(err)
	}
	num, err := llenado.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("nuevo canduntor creado ", num)

}

func Crea_table() {
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

	table := `
	CREATE TABLE IF NOT EXISTS garajes(
		id INT AUTO_INCREMENT PRIMARY KEY,
		id_buces INT NOT NULL,
		nombre VARCHAR(30) NOT NULL,
		direccion VARCHAR(40) NOT NULL,
		FOREIGN KEY (id_buces) REFERENCES buces(id)
			ON UPDATE CASCADE
			ON DELETE RESTRICT

	);
	`

	creado, err := db.Exec(table)
	if err != nil {
		fmt.Println("pipi")
	}
	num, err := creado.RowsAffected()
	fmt.Println("creado", num)

}
