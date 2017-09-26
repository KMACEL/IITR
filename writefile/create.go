package writefile

import (
	"os"
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
	fc, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err = fc.Close(); err != nil {
			panic(err)
		}
	}()

	return false
}

func OpenFile(outFileName string) bool {
	if fo, err = os.OpenFile(outFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {
		panic(err)
	} else {
		return true
	}
}

var WriteFlag = false

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
}

func WriteFileByte(data []byte) {

	if _, err = fo.Write(data); err != nil {
		panic(err)
	}

	defer func() {
		if _, err = fo.WriteString("\n"); err != nil {
			panic(err)
		}

	}()
}
