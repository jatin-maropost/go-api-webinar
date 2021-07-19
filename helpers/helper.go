package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func GetEnvVariable(key string) (string,bool) {
	return  os.LookupEnv(key)
  }

  // Load env variables defined in env file
  func LoadEnvFile(file string)  {
	// load .env file
	err := godotenv.Load(file)
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  }