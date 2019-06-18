package logc

import (
	"log"
	"os"

	"github.com/KMACEL/IITR/timop"
)

var Debug = true
var ConnectionPrintDebug = true
var QueryPrintDebug = true

// QueryPrint is
func QueryPrint(args ...interface{}) {
	if ConnectionPrintDebug {
		if _, err := os.Stat("./logc"); os.IsNotExist(err) {
			os.MkdirAll("./logc", os.ModePerm)
		}

		f, err := os.OpenFile("logc/queryLogFile_"+timop.GetTimeNamesFormatDaysTYPE2()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
	}
}

// ConnectionPrint is
func ConnectionPrint(args ...interface{}) {
	if ConnectionPrintDebug {
		if _, err := os.Stat("./logc"); os.IsNotExist(err) {
			os.MkdirAll("./logc", os.ModePerm)
		}

		f, err := os.OpenFile("logc/connectionLogFile_"+timop.GetTimeNamesFormatDaysTYPE2()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
	}
}

// OperationPrint is
func OperationPrint(args ...interface{}) {
	if Debug {
		if _, err := os.Stat("./logc"); os.IsNotExist(err) {
			os.MkdirAll("./logc", os.ModePerm)
		}

		f, err := os.OpenFile("logc/_"+timop.GetTimeNamesFormatDaysTYPE2()+"operationLogFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
	}
}

// ReportPrint is
func ReportPrint(args ...interface{}) {
	if Debug {
		if _, err := os.Stat("./logc"); os.IsNotExist(err) {
			os.MkdirAll("./logc", os.ModePerm)
		}

		f, err := os.OpenFile("logc/reportLogFile_"+timop.GetTimeNamesFormatDaysTYPE2()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
	}
}

// TestPrint is
func TestPrint(args ...interface{}) {
	if Debug {
		if _, err := os.Stat("./logc"); os.IsNotExist(err) {
			os.MkdirAll("./logc", os.ModePerm)
		}

		f, err := os.OpenFile("logc/testLogFile_"+timop.GetTimeNamesFormatDaysTYPE2()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
		//fmt.Fprint(xx, "ddf")
	}
}

// GlobalPrint is
func GlobalPrint(args ...interface{}) {
	if Debug {
		if _, err := os.Stat("./logc"); os.IsNotExist(err) {
			os.MkdirAll("./logc", os.ModePerm)
		}

		f, err := os.OpenFile("logc/globalLogFile_"+timop.GetTimeNamesFormatDaysTYPE2()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
	}
}
