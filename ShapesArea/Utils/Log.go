package Utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Log(shape string, area float64) {
	fileName := getCurrentDate() + ".txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println(shape+" area calculated.", area)
}

func getCurrentDate() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%v-%v-%v", year, int(month), day)
}

func UserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}
