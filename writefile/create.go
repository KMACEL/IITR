package writefile

import (
	"os"

	"github.com/KMACEL/IITR/errc"
)

var (
	//fi  *os.File
	fo *os.File
	//fc  *os.File
	err error
)

//CreateFile is
func CreateFile(fileName string) bool {
	//open output file
	fc, errCreate := os.Create(fileName)
	if errCreate != nil {
		errc.ErrorCenter("Create", errCreate)
		panic(errCreate)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if errCreate = fc.Close(); errCreate != nil {
			errc.ErrorCenter("Create", errCreate)
			panic(errCreate)
		}
	}()

	return false
}

/*// OpenFile is
func OpenFile(outFileName string) bool {
	if fo, err = os.OpenFile(outFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {
		panic(err)
	} else {
		return true
	}
}*/

// WriteFlag is
var WriteFlag = false

// WriteArray is
func WriteArray(writeText []string) {
	if WriteFlag == false {
		WriteFlag = true

		for _, text := range writeText {
			if _, err = fo.WriteString(text + ","); err != nil {
				panic(err)
			}
		}

		defer func() {
			if _, err = fo.WriteString("\n"); err != nil {
				panic(err)
			}
		}()
		WriteFlag = false
	}
}

// WriteByte is
func WriteByte(data []byte) {

	if _, err = fo.Write(data); err != nil {
		panic(err)
	}

	defer func() {
		if _, err = fo.WriteString("\n"); err != nil {
			panic(err)
		}

	}()
}

/*// WriteText is
func WriteText(writeText ...string) {

	for _, text := range writeText {
		if _, err = fo.WriteString(text + ","); err != nil {
			panic(err)
		}
	}

	defer func() {
		if _, err = fo.WriteString("\n"); err != nil {
			panic(err)
		}

	}()
}*/

// OpenFile is
func OpenFile(outFileName string, openFile *os.File) *os.File {
	if openFile, errOpenFile := os.OpenFile(outFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend); errOpenFile != nil {
		errc.ErrorCenter("Open", errOpenFile)
		panic(errOpenFile)
	} else {
		return openFile
	}
}

// WriteText is
func WriteText(openFile *os.File, writeText ...string) {
	for _, text := range writeText {
		if _, errWriteText := openFile.WriteString(text + ","); errWriteText != nil {
			errc.ErrorCenter("Open", errWriteText)
			panic(errWriteText)
		}
	}

	defer func() {
		if _, errWriteText := openFile.WriteString("\n"); errWriteText != nil {
			errc.ErrorCenter("Open", errWriteText)
			panic(errWriteText)
		}

	}()
}
