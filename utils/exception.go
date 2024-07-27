package utils

import "log"

func PanicLogging(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
