package utils

import (
	"log"
	"os"
)

type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

var CustomLogger Logger = Logger{
	log.New(os.Stdout, "INFO:  ", log.Ldate|log.Ltime),
	log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime),
}
