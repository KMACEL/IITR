package errc

import (
	"errors"
	"log"
)

/*
███████╗██████╗ ██████╗  ██████╗ ██████╗      ██████╗███████╗███╗   ██╗████████╗███████╗██████╗
██╔════╝██╔══██╗██╔══██╗██╔═══██╗██╔══██╗    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝██╔════╝██╔══██╗
█████╗  ██████╔╝██████╔╝██║   ██║██████╔╝    ██║     █████╗  ██╔██╗ ██║   ██║   █████╗  ██████╔╝
██╔══╝  ██╔══██╗██╔══██╗██║   ██║██╔══██╗    ██║     ██╔══╝  ██║╚██╗██║   ██║   ██╔══╝  ██╔══██╗
███████╗██║  ██║██║  ██║╚██████╔╝██║  ██║    ╚██████╗███████╗██║ ╚████║   ██║   ███████╗██║  ██║
╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝     ╚═════╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝  ╚═╝
*/

// Rest -> Query Constant
const (
	RequestGet = "Request Get : "
	DoGet      = "Do Get : "
	BodyGet    = "Body Get : "

	RequestPost = "Request Get :"
	DoPost      = "Do Get : "
	BodyPost    = "Body Get : "

	RequestPut = "Request Get : "
	DoPut      = "Do Get : "
	BodyPut    = "Body Get : "
)

// Rest -> Device Constant
const (
	Summary = "Summary : "
)

// writefile -> create Constant
const (
	CreateFile = "Create New File :"
	OpenFile   = "Open File :"

	WriteArray = "Write Array : "
	WriteByte  = "Write Byte : "
	WriteText  = "Write Text : "
)

// Errors Type
var (
	ErrorNotFound404 = errors.New("Request is 404 Not Found. Please check variables, queries, links and other parameters")
)

// ErrorCenter is
func ErrorCenter(title string, err error) {
	if err != nil {
		log.Println("Error IITR - "+title, err.Error())
	}
}
