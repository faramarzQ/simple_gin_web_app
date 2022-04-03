package config

import h "gin/helpers"

var db = map[string]string{
	"DRIVER_NAME": h.ENV("DRIVER_NAME"),
	"NAME":        h.ENV("DB_DATABASE"),
	"PASSWORD":    h.ENV("DB_PASSWORD", ""),
	"USERNAME":    h.ENV("DB_USERNAME"),
	"HOST":        h.ENV("DB_HOST"),
	"PORT":        h.ENV("DB_PORT"),
}

func DB(key string) string {
	return db[key]
}
