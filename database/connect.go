package database

import (
	"CRM-test/structures"
	"CRM-test/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Link *sql.DB

func Connection(config *structures.Setting) {
	var e error
	Link, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName))
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	e = Link.Ping()
	if e != nil {
		utils.Logger.Println(e)
		return
	}
}
