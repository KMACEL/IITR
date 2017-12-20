package operations

import (
	"github.com/KMACEL/IITR/rest/device"
	"github.com/KMACEL/IITR/rest/download"
)

/*
██████╗  ██████╗ ██╗    ██╗███╗   ██╗██╗      ██████╗  █████╗ ██████╗         ██╗      ██████╗  ██████╗ ███████╗
██╔══██╗██╔═══██╗██║    ██║████╗  ██║██║     ██╔═══██╗██╔══██╗██╔══██╗        ██║     ██╔═══██╗██╔════╝ ██╔════╝
██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║██║     ██║   ██║███████║██║  ██║        ██║     ██║   ██║██║  ███╗███████╗
██║  ██║██║   ██║██║███╗██║██║╚██╗██║██║     ██║   ██║██╔══██║██║  ██║        ██║     ██║   ██║██║   ██║╚════██║
██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║███████╗╚██████╔╝██║  ██║██████╔╝        ███████╗╚██████╔╝╚██████╔╝███████║
╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═════╝         ╚══════╝ ╚═════╝  ╚═════╝ ╚══════╝
*/
// For use Example:
//     var downloadLogOperations util.DownloadLogs
//     downloadLogOperations.Start("867377020740787")

// DownloadLogs is
type DownloadLogs struct{}

// Start is
func (d DownloadLogs) Start(downloadLogDeviceID ...string) {
	var (
		downloadVar download.Download
		deviceVar   device.Device
	)

	for _, deviceID := range downloadLogDeviceID {
		go downloadVar.DownloadLog(deviceVar.DeviceID2Code(deviceID))
	}
}
