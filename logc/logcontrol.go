package logc

import (
	"log"
	"os"
)

// OperationPrint is
func OperationPrint(args ...interface{}) {
	if _, err := os.Stat("./logc"); os.IsNotExist(err) {
		os.MkdirAll("./logc", os.ModePerm)
	}

	f, err := os.OpenFile("logc/operationLogFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	f, err := os.OpenFile("logc/reportLogFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	f, err := os.OpenFile("logc/testLogFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	f, err := os.OpenFile("logc/globalLogFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
}
