package loggers

import "log"

func Log(msg string){
log.Print(msg)
}

func Fatal(err error) {
	log.Fatalf("Some error occurred: %s", err)
}

func PanicIfNotNil(err error){
	if err != nil {
		log.Fatalf("Some error occurred: %s", err)
	}
}

