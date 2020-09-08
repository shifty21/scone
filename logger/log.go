package logger

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	//Trace type
	Trace *log.Logger
	//Info type
	Info *log.Logger
	//Warning type
	Warning *log.Logger
	//Error type
	Error *log.Logger
)

//InitLogger initializes all the handlers
func InitLogger() {
	traceHandle := ioutil.Discard
	infoHandle := os.Stdout
	warningHandle := os.Stdout
	errorHandle := os.Stderr

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
