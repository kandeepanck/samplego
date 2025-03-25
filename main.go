package main

import (
	"log"
	"os"
	"samplego/inegration/helpers"
	"time"
)

func main() {

	lFile, lErr := os.OpenFile("./log/log"+time.Now().Format("02012006150405 ")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if lErr != nil {
		log.Println("error on creating file :", lErr)
		panic(lErr)
	}
	defer lFile.Close()
	log.SetOutput(lFile)
	log.Println("hello word")

	lDebug := new(helpers.Debug)
	lDebug.Init()

	// log.Println("")
}
