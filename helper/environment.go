package helper

import(
	"github.com/joho/godotenv"
)

func init(){
	//Undo for local environment, other enviroments will be managed by docker-compose parameters
	godotenv.Load("environments/LOCAL.env")
}