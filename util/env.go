package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/mattn/go-isatty"
)

// Gets the path to the .env file
// When running locally, the file is in the same folder
// When running compiled, the file is in the exe folder
func getEnvPath() string {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		return "./.env"
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dir, ".env")
}

// Loads the .env file
func LoadEnv() {
	err := godotenv.Load(getEnvPath())
	if err != nil {
		log.Fatal(err)
	}
}
