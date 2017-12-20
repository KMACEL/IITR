package errc

import (
	"log"
	"os"
)

/*
███████╗██████╗ ██████╗  ██████╗ ██████╗      ██████╗███████╗███╗   ██╗████████╗███████╗██████╗
██╔════╝██╔══██╗██╔══██╗██╔═══██╗██╔══██╗    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝██╔════╝██╔══██╗
█████╗  ██████╔╝██████╔╝██║   ██║██████╔╝    ██║     █████╗  ██╔██╗ ██║   ██║   █████╗  ██████╔╝
██╔══╝  ██╔══██╗██╔══██╗██║   ██║██╔══██╗    ██║     ██╔══╝  ██║╚██╗██║   ██║   ██╔══╝  ██╔══██╗
███████╗██║  ██║██║  ██║╚██████╔╝██║  ██║    ╚██████╗███████╗██║ ╚████║   ██║   ███████╗██║  ██║
╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝     ╚═════╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// ErrorCenter is
func ErrorCenter(title string, err error) {
	if err != nil {
		log.Println("Error IITR - "+title, " : ", err.Error())
		errorFile("Error IITR - "+title, " : ", err.Error())
	}
}

func errorFile(args ...interface{}) {
	if _, err := os.Stat("./errc"); os.IsNotExist(err) {
		os.MkdirAll("./errc", os.ModePerm)
	}

	f, err := os.OpenFile("errc/errorLogFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error IITR - Error File : Error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(args...)
}
