package db

import(
	"database/sql"
)

//Additional scripts like trigger, SP etc. 
func InitScripts(db *sql.DB) {
	createUserTable := `create table if not exists User(
					id int(11) NOT NULL,   
					name varchar(50) NOT NULL,       
					surname varchar(50)  NOT NULL
					)
					`
	db.Query(createUserTable)
}
