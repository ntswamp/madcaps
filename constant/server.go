package constant

import "os"

var SERVER_TYPE string = os.Getenv("SERVER_TYPE")

func CheckServerEnv() string {
	return SERVER_TYPE
}

func IsProduction() bool {
	switch CheckServerEnv() {
	case "production":
		return true
	}
	return false
}
