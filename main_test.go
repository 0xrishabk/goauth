package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnvironmen(t *testing.T) {
	godotenv.Load(".env")
	fmt.Println(os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}
