package main

import (
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucas5z/crud-nativo-practica/routes"
)

func main() {
	//routes.Open_mysql()
	routes.Crea_user()
	routes.Read_user()
	routes.Update_user()
	routes.Delete_user()
	routes.Crea_table()
}
