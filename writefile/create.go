package writefile

import (
	"os"

	"github.com/KMACEL/IITR/errc"
)

/*
██╗    ██╗██████╗ ██╗████████╗███████╗        ███████╗██╗██╗     ███████╗
██║    ██║██╔══██╗██║╚══██╔══╝██╔════╝        ██╔════╝██║██║     ██╔════╝
██║ █╗ ██║██████╔╝██║   ██║   █████╗          █████╗  ██║██║     █████╗
██║███╗██║██╔══██╗██║   ██║   ██╔══╝          ██╔══╝  ██║██║     ██╔══╝
╚███╔███╔╝██║  ██║██║   ██║   ███████╗        ██║     ██║███████╗███████╗
 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝   ╚═╝   ╚══════╝        ╚═╝     ╚═╝╚══════╝╚══════╝
*/
// WriteFile is a package created to print tests, reports, errors, and files.
// This package uses the 'os' library, especially the 'csv' format, for tabling data.
// This package can become direct communication from other packages.

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗        ███████╗██╗██╗     ███████╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝        ██╔════╝██║██║     ██╔════╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗          █████╗  ██║██║     █████╗
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝          ██╔══╝  ██║██║     ██╔══╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗        ██║     ██║███████╗███████╗
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝        ╚═╝     ╚═╝╚══════╝╚══════╝
*/

// CreateFile is used to create a new file. This function externally takes only
// a value of type "string" with the name "fileName". This value determines
// the name of the file to be created. This function performs util using the command "os.Create".
// The file is closed after it is created. This is the intent of the "defer" block.
func CreateFile(fileName string) {
	//open output file
	fc, errCreate := os.Create(fileName)
	if errCreate != nil {
		errc.ErrorCenter(createFileTag, errCreate)
		panic(errCreate)
	}

	defer func() {
		if errCreate = fc.Close(); errCreate != nil {
			errc.ErrorCenter(createFileTag, errCreate)
			panic(errCreate)
		}
	}()
}

/*
 ██████╗ ██████╗ ███████╗███╗   ██╗        ███████╗██╗██╗     ███████╗
██╔═══██╗██╔══██╗██╔════╝████╗  ██║        ██╔════╝██║██║     ██╔════╝
██║   ██║██████╔╝█████╗  ██╔██╗ ██║        █████╗  ██║██║     █████╗
██║   ██║██╔═══╝ ██╔══╝  ██║╚██╗██║        ██╔══╝  ██║██║     ██╔══╝
╚██████╔╝██║     ███████╗██║ ╚████║        ██║     ██║███████╗███████╗
 ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═══╝        ╚═╝     ╚═╝╚══════╝╚══════╝
*/

// OpenFile performs the opening process to write or read a file.
// This function takes two external values.
// openingFileName: This gets the name with the extension of the file to be opened.
//     If it is in the same folder as the program, only the name is enough.
// openedOSFile: This variable is important. Several files can be opened
//     at the same time and separate util can be performed on those files.
//     If you have a single "os.File" variable, you can not do the same operation.
//     For this reason, we expect the user to define a variable "var fileVarible * os.File"
//    instead of a constant variable and send this variable value to the "OpenFile" function.
//    This function returns the file that it opened in the form of "os.File" when opening.
//    Thus, when the file is opened, writing or reading util can be performed through the variable that returns easily.
//   Example Usage:
//      var openFile1 *os.File
//      writefile.CreateFile("testFile.csv")
//      openFile1 = writefile.OpenFile2("testFile.csv", openFile1)
func OpenFile(openedOSFile *os.File, openingFileName string) *os.File {
	if openedOSFile, errOpenFile := os.OpenFile(openingFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend); errOpenFile != nil {
		errc.ErrorCenter(openFileTag, errOpenFile)
		panic(errOpenFile)
	} else {
		return openedOSFile
	}
}

/*
██╗    ██╗██████╗ ██╗████████╗███████╗         █████╗ ██████╗ ██████╗  █████╗ ██╗   ██╗
██║    ██║██╔══██╗██║╚══██╔══╝██╔════╝        ██╔══██╗██╔══██╗██╔══██╗██╔══██╗╚██╗ ██╔╝
██║ █╗ ██║██████╔╝██║   ██║   █████╗          ███████║██████╔╝██████╔╝███████║ ╚████╔╝
██║███╗██║██╔══██╗██║   ██║   ██╔══╝          ██╔══██║██╔══██╗██╔══██╗██╔══██║  ╚██╔╝
╚███╔███╔╝██║  ██║██║   ██║   ███████╗        ██║  ██║██║  ██║██║  ██║██║  ██║   ██║
 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝   ╚═╝   ╚══════╝        ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝
*/

// WriteArray writes the data received in the Array type to the opened file and closes that file. It takes two values.
// writeTextArray: takes an array of type string. This array is written to the file.
// openedFile: The file to be opened.
func WriteArray(openedFile *os.File, writeTextArray []string) {
	for _, text := range writeTextArray {
		if _, errWriteArray := openedFile.WriteString(text + ","); errWriteArray != nil {
			errc.ErrorCenter(writeArrayTag, errWriteArray)
			panic(errWriteArray)
		}
	}
	/*defer func() {
		if _, errWriteArray := openedFile.WriteString("\n"); errWriteArray != nil {
			errc.ErrorCenter(writeArrayTag, errWriteArray)
			panic(errWriteArray)
		}
	}()*/

}

/*
██╗    ██╗██████╗ ██╗████████╗███████╗        ██████╗ ██╗   ██╗████████╗███████╗
██║    ██║██╔══██╗██║╚══██╔══╝██╔════╝        ██╔══██╗╚██╗ ██╔╝╚══██╔══╝██╔════╝
██║ █╗ ██║██████╔╝██║   ██║   █████╗          ██████╔╝ ╚████╔╝    ██║   █████╗
██║███╗██║██╔══██╗██║   ██║   ██╔══╝          ██╔══██╗  ╚██╔╝     ██║   ██╔══╝
╚███╔███╔╝██║  ██║██║   ██║   ███████╗        ██████╔╝   ██║      ██║   ███████╗
 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝   ╚═╝   ╚══════╝        ╚═════╝    ╚═╝      ╚═╝   ╚══════╝
*/

// WriteByte writes the data in the byte type received in the Array type to the opened file and closes that file. It takes two values.
// writeTextByte: takes an array of type byte. This array is written to the file.
// openedFile: The file to be opened.
func WriteByte(openedFile *os.File, writeTextByte []byte) {
	if _, errWriteByte := openedFile.Write(writeTextByte); errWriteByte != nil {
		errc.ErrorCenter(writeByteTag, errWriteByte)
		panic(errWriteByte)
	}
	defer func() {
		if _, errWriteByte := openedFile.WriteString("\n"); errWriteByte != nil {
			errc.ErrorCenter(writeByteTag, errWriteByte)
			panic(errWriteByte)
		}

	}()
}

/*
██╗    ██╗██████╗ ██╗████████╗███████╗        ████████╗███████╗██╗  ██╗████████╗
██║    ██║██╔══██╗██║╚══██╔══╝██╔════╝        ╚══██╔══╝██╔════╝╚██╗██╔╝╚══██╔══╝
██║ █╗ ██║██████╔╝██║   ██║   █████╗             ██║   █████╗   ╚███╔╝    ██║
██║███╗██║██╔══██╗██║   ██║   ██╔══╝             ██║   ██╔══╝   ██╔██╗    ██║
╚███╔███╔╝██║  ██║██║   ██║   ███████╗           ██║   ███████╗██╔╝ ██╗   ██║
 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝   ╚═╝   ╚══════╝           ╚═╝   ╚══════╝╚═╝  ╚═╝   ╚═╝
*/

// WriteText retrieves data in series. A serial data stream gets the first line of a csv formatted data.
// Thus, how many columns have been filled in so many columns, how many times this function is called,
// will create as many lines. Unlike the others, the variable "os.File" comes first because it needs to be a variable terminator.
// openedFile: Specifies the file to be written.
// writeText: a serialized variable of type string. This function can be sent as a string type by putting "," between them.
func WriteText(openedFile *os.File, writeText ...string) {
	for _, text := range writeText {
		if _, errWriteText := openedFile.WriteString(text + ","); errWriteText != nil {
			errc.ErrorCenter(writeTextTag, errWriteText)
			panic(errWriteText)
		}
	}
	defer func() {
		if _, errWriteText := openedFile.WriteString("\n"); errWriteText != nil {
			errc.ErrorCenter(writeTextTag, errWriteText)
			panic(errWriteText)
		}
	}()
}
