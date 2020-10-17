package db

import(
	"database/sql"
)

func InitScripts(db *sql.DB) {
	createUserTable := `create table if not exists Users(
					id int(11) NOT NULL,   
					name varchar(50) NOT NULL,       
					surname varchar(50)  NOT NULL
					)
					`
	db.Query(createUserTable)
}
