package config

import (
	"os"
	"strconv"
)

var (
	DatabaseURL       string
	MaxDBConn         int
	MigrationFilePath string

	Port string

	TokenLen    int
	TokenExpiry int

	BasicAuthUserName string
	BasicAuthPassword string
)

func init() {
	DatabaseURL = os.Getenv("DatabaseURL")
	MaxDBConn, _ = strconv.Atoi(os.Getenv("MaxDBConn"))

	Port = os.Getenv("Port")

	TokenLen, _ = strconv.Atoi(os.Getenv("TokenLen"))
	if TokenLen == 0 {
		TokenLen = 12
	}

	TokenExpiry, _ = strconv.Atoi(os.Getenv("TokenExpiry"))
	if TokenExpiry == 0 {
		TokenExpiry = 7
	}

	BasicAuthUserName = os.Getenv("BasicAuthUserName")
	BasicAuthPassword = os.Getenv("BasicAuthPassword")
}
