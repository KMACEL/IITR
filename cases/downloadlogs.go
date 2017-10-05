package cases

import "github.com/KMACEL/IITR/rest/download"

// DownloadLogs is
type DownloadLogs struct{}

//Delete Packet

// Start is
func (d DownloadLogs) Start() {
	getLogList := []string{"c89ce6a816704b06aa9b8c93d4085edc", "a2f26ff97910475bb24b5e8bb9fad7fb", "9f737cd7bd7e47e5a4dcf98f415d293b", "7b1cee83329d453dbdad40e9953a86d0", "814df77a28ce4b99952412653303a3d6", "2731663e9afe499f8fa681298a7d07ac", "6745c042e79d4dad8fcb126dcef359de", "4e341a6b45454ef98430fb510d655ff4", "1537d4894f424318b173c3d39bdf06f1", "401ec27380c74c6d82b575ea4fd1c1aa", "ead6d1bf8d4e4ddcac11068e505623a7", "518261278eaa41578e17a329e0b70e7f", "2e265ee5436c47ad9c139c0e9f48a3c6", "e4d993d6abd4479c92e89c5fec94d4a4", "0e2cf4d59c144856ad3113a66b9345a3", "2e79e3a45cc94a88bc668f6f37ad8267", "49d249b936f54ec4ac59aa55a5238794", "f9ec6efd27b646648d12707dc7fd0a67"}
	for _, deviceCode := range getLogList {
		go download.Download{}.DownloadLog(deviceCode)
	}
}
