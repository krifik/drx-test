package exception

import "log"

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func LogIfNeeded(err interface{}) {
	if err != nil {
		log.Println("This error can be ignored. Error: ", err)
	}
}
