package define

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

var MysqlDSN = os.Getenv("MysqlDSN")

type M map[string]interface{}

type UserClaim struct {
	Id       uint   `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var (
	JwtKey = "iot-platform"
)
