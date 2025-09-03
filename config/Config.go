package config

import(
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

func GetEnv(key string, defaultValue string) string{
	value, exsist := os.LookupEnv(key)
	if !exsist{
		return defaultValue
	}
	return value
}