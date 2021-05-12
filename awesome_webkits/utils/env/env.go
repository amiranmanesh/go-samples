package env

import (
	"fmt"
	"github.com/juju/errors"
	"os"

	"github.com/joho/godotenv"
)

/*
This is a func to get a .env data by its key
Its get a key as a string and return a string as the value
*/
func GetEnvItem(envKey string) string {
	if err := godotenv.Load(); err != nil {
		panic(errors.Trace(err))
	}
	value := ""
	value = os.Getenv(envKey)
	return value
}

/*
This is a func to get all .env keys and values as a map
*/
func GetEnvItems() map[string]string {
	if err := godotenv.Load(); err != nil {
		panic(errors.Trace(err))
	}
	myEnv, err := godotenv.Read()
	if err != nil {
		return map[string]string{}
	}
	return myEnv
}

/*
This is func to append a key value map into .env file
*/
func AppendInEnvFile(data map[string]string) bool {
	if err := godotenv.Load(); err != nil {
		panic(errors.Trace(err))
	}
	allData := GetEnvItems()
	for key, val := range data {
		allData[key] = val
	}
	var s string
	for key, val := range allData {
		s += fmt.Sprintf("%s=%s\n", key, val)
	}
	env, err := godotenv.Unmarshal(s)
	err = godotenv.Write(env, "./.env")
	return err == nil
}
