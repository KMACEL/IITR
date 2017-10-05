package download

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/writefile"
)

// Download is
type Download struct {
	logListJSON device.LogListJSON
}

//DownloadLog is
func (d Download) DownloadLog(setDeviceCode string) {
	query := device.Device{}.GetLogList(setDeviceCode, rest.NOMarshal, rest.Invisible)
	json.Unmarshal(query, &d.logListJSON)
	for _, logs := range d.logListJSON {
		go d.readFile(logs.URL, logs.Token, logs.DeviceID+"_"+logs.Name)
	}
}

func (d Download) readFile(url string, token string, fileName string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Set("X-Auth-Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	d.saveFile(respBody, fileName)
}

func (d Download) saveFile(data []byte, fileName string) {
	writefile.CreateFile(fileName)
	writefile.OpenFile(fileName)
	writefile.WriteByte(data)
}
