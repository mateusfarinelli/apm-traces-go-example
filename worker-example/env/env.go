package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const DB_KIND = "sqlserver"

var (
	Environment      string
	Version          string
	HttpPort         int
	HttpReadTimeout  int
	HttpWriteTimeout int
)

func LoadEnvs() {
	godotenv.Load()

	Environment = os.Getenv("ENV")
	Version = os.Getenv("VERSION")
	HttpPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))
	HttpReadTimeout, _ = strconv.Atoi(os.Getenv("HTTP_READ_TIMEOUT"))
	HttpWriteTimeout, _ = strconv.Atoi(os.Getenv("HTTP_WRITE_TIMEOUT"))
}
