package download

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/writefile"
)

/*
██████╗  ██████╗ ██╗    ██╗███╗   ██╗██╗      ██████╗  █████╗ ██████╗
██╔══██╗██╔═══██╗██║    ██║████╗  ██║██║     ██╔═══██╗██╔══██╗██╔══██╗
██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║██║     ██║   ██║███████║██║  ██║
██║  ██║██║   ██║██║███╗██║██║╚██╗██║██║     ██║   ██║██╔══██║██║  ██║
██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║███████╗╚██████╔╝██║  ██║██████╔╝
╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═════╝
*/

// Download package does  download the computer by naming the logs received by IoT-Ignite.
type Download struct {
}

/*
██████╗  ██████╗ ██╗    ██╗███╗   ██╗██╗      ██████╗  █████╗ ██████╗         ██╗      ██████╗  ██████╗
██╔══██╗██╔═══██╗██║    ██║████╗  ██║██║     ██╔═══██╗██╔══██╗██╔══██╗        ██║     ██╔═══██╗██╔════╝
██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║██║     ██║   ██║███████║██║  ██║        ██║     ██║   ██║██║  ███╗
██║  ██║██║   ██║██║███╗██║██║╚██╗██║██║     ██║   ██║██╔══██║██║  ██║        ██║     ██║   ██║██║   ██║
██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║███████╗╚██████╔╝██║  ██║██████╔╝        ███████╗╚██████╔╝╚██████╔╝
╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═════╝         ╚══════╝ ╚═════╝  ╚═════╝
*/

/*
	download.Download{}.DownloadLog(device.Device{}.DeviceID2Code("{YOUR_DEVICE_ID}"))
*/

//DownloadLog is function used to download all logs on the given device code.
func (d Download) DownloadLog(setDeviceCode string) {
	var logListJSON device.LogListJSON

	query := device.Device{}.GetLogList(setDeviceCode, rest.NOMarshal, rest.Invisible)
	if query != nil {
		err := json.Unmarshal(query, &logListJSON)
		errc.ErrorCenter("Download-Log-Json", err)

		for _, logs := range logListJSON {
			go d.readFile(logs.URL, logs.Token, logs.DeviceID+"_"+logs.Name)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// readFile is connected to the given log position and reads the data
func (d Download) readFile(url string, token string, fileName string) {
	req, err := http.NewRequest("GET", url, nil)
	errc.ErrorCenter("Download-Log-Json", err)

	req.Header.Set("X-Auth-Token", token)

	resp, errDO := http.DefaultClient.Do(req)
	errc.ErrorCenter("Download-Log-readFile", errDO)

	defer resp.Body.Close()

	respBody, errcRead := ioutil.ReadAll(resp.Body)
	errc.ErrorCenter("Download-Log-ReadAll", errcRead)

	d.saveFile(respBody, fileName)
}

// saveFile saves the read data to the file in the specified format.
func (d Download) saveFile(data []byte, fileName string) {
	writefile.CreateFile(fileName)
	var saveFile *os.File
	saveFile = writefile.OpenFile(saveFile, fileName)
	writefile.WriteByte(saveFile, data)
}
