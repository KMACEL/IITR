package main

import (
	"os"
	"log"
)

func InıtLog(){

}

func Println(args ...interface{}) {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
}


func main(){
	Println("fdfdf","fefdf")
	Println("qqqq","rrrr","ggggg")
}