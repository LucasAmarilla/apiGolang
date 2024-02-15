package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Erro carregando variaveis de ambientes")
		}
	}

}
