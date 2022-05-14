package helpers

import "log"

func CheckAndThrowException(err error) {
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
}
