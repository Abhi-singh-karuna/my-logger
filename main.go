package mylogger

import "log"

func Loginfo(a string) {
	log.Printf("INFO - %v", a)
}

func LogWarning(a string) {
	log.Printf("WARN - %v", a)
}

func LogError(a string) {
	log.Printf("ERROR - %v", a)
}
