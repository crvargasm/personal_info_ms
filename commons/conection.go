package commons

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() (db *sql.DB) {
	driver := "mysql"
	user := "root"
	contrasenia := "poiuasdoiweq123we1232130asd"
	nombreDB := "personalInfoDb"
	direccionIP := "172.17.0.2"
	port := "3306"

	db, err := sql.Open(driver, user+":"+contrasenia+"@tcp("+direccionIP+":"+port+")/"+nombreDB)
	if err != nil {
		panic(err.Error())
	}

	return db
}
