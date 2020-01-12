package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)

}

func main() {

	loggingSettings("testlog.txt")
	_, err := os.Open("gagagaa")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	/*
		log.Println("logging!")
		log.Printf("%T, %v", "test", "test")

		log.Fatalln("error!!")
	*/
	// Fatalln以降は実行されない
	fmt.Println("ok!")

}
