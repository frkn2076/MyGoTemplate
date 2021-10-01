package environments

import(
	"os"

	"github.com/joho/godotenv"
)

func init(){
	env := os.Getenv("ENV")
	if env == "UAT" {
		godotenv.Load("infra/environments/UAT.env")
	} else if env == "PROD" {
		godotenv.Load("infra/environments/PROD.env")
	} else {
		godotenv.Load("infra/environments/DEV.env")
	}
}