package logc

import (
	"log"
	"os"

	"github.com/KMACEL/IITR/timop"
)

// QueryPrint is
func QueryPrint(args ...interface{}) {
	if _, err := os.Stat("./logc"); os.IsNotExist(err) {
		os.MkdirAll("./logc", os.ModePerm)
	}

	f, err := os.OpenFile("logc/queryLogFile_"+timop.GetTimeNamesFormatDays()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
}

// ConnectionPrint is
func ConnectionPrint(args ...interface{}) {
	if _, err := os.Stat("./logc"); os.IsNotExist(err) {
		os.MkdirAll("./logc", os.ModePerm)
	}

	f, err := os.OpenFile("logc/connectionLogFile_"+timop.GetTimeNamesFormatDays()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
}

// OperationPrint is
func OperationPrint(args ...interface{}) {
	if _, err := os.Stat("./logc"); os.IsNotExist(err) {
		os.MkdirAll("./logc", os.ModePerm)
	}

	f, err := os.OpenFile("logc/_"+timop.GetTimeNamesFormatDays()+"operationLogFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
}

// ReportPrint is
func ReportPrint(args ...interface{}) {
	if _, err := os.Stat("./logc"); os.IsNotExist(err) {
		os.MkdirAll("./logc", os.ModePerm)
	}

	f, err := os.OpenFile("logc/reportLogFile_"+timop.GetTimeNamesFormatDays()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
}

// TestPrint is
func TestPrint(args ...interface{}) {
	if _, err := os.Stat("./logc"); os.IsNotExist(err) {
		os.MkdirAll("./logc", os.ModePerm)
	}

	f, err := os.OpenFile("logc/testLogFile_"+timop.GetTimeNamesFormatDays()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
	//fmt.Fprint(xx, "ddf")
}

// GlobalPrint is
func GlobalPrint(args ...interface{}) {
	if _, err := os.Stat("./logc"); os.IsNotExist(err) {
		os.MkdirAll("./logc", os.ModePerm)
	}

	f, err := os.OpenFile("logc/globalLogFile_"+timop.GetTimeNamesFormatDays()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
}
