package utils

import "github.com/joho/godotenv"

func LoadDotenv() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
}
