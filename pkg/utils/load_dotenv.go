package utils

import "github.com/joho/godotenv"

func LoadDotenv(stage string) {
	if err := godotenv.Load(stage); err != nil {
		panic(err.Error())
	}
}
