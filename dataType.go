package golib

import (
	"os"

	"github.com/joho/godotenv"
)

type DataType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | byte | string
}

func GetVersion() string {
	godotenv.Load("env")
	return os.Getenv("VERSION")
}
