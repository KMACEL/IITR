package errc

import (
	"log"
	"os"

	"github.com/KMACEL/IITR/timop"
)

/*
███████╗██████╗ ██████╗  ██████╗ ██████╗      ██████╗███████╗███╗   ██╗████████╗███████╗██████╗
██╔════╝██╔══██╗██╔══██╗██╔═══██╗██╔══██╗    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝██╔════╝██╔══██╗
█████╗  ██████╔╝██████╔╝██║   ██║██████╔╝    ██║     █████╗  ██╔██╗ ██║   ██║   █████╗  ██████╔╝
██╔══╝  ██╔══██╗██╔══██╗██║   ██║██╔══██╗    ██║     ██╔══╝  ██║╚██╗██║   ██║   ██╔══╝  ██╔══██╗
███████╗██║  ██║██║  ██║╚██████╔╝██║  ██║    ╚██████╗███████╗██║ ╚████║   ██║   ███████╗██║  ██║
╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝     ╚═════╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝  ╚═╝
*/

var Debug = true

// ErrorCenter is
func ErrorCenter(title string, err error) {
	if Debug {
		if err != nil {
			log.Println("Error IITR - "+title, " : ", err.Error())
			errorFile("Error IITR - "+title, " : ", err.Error())
		}
	}
}

func errorFile(args ...interface{}) {
	if Debug {
		if _, err := os.Stat("./errc"); os.IsNotExist(err) {
			os.MkdirAll("./errc", os.ModePerm)
		}

		f, err := os.OpenFile("errc/errorLogFile_"+timop.GetTimeNamesFormatDaysTYPE2()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error IITR - Error File : Error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(args...)
	}
}
