package db

func Init() {
	createUser := `create table if not exists Users(
					id int(11) NOT NULL,   
					name varchar(50) NOT NULL,       
					surname varchar(50)  NOT NULL
					)
					`
	DB.Query(createUser)
}

