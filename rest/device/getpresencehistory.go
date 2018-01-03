package device

import (
	"encoding/json"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

//PresenceHistory is
func (d Device) PresenceHistory(deviceID string, setVisualFlag bool, setUnMarshal bool) []byte {
	var rq rest.Query

	queryLink := presenceHistoryLink(deviceID)
	visualFlag := setVisualFlag
	query, errDownloadAppList := rq.GetQuery(queryLink, visualFlag)
	errc.ErrorCenter(deviceID, errDownloadAppList)

	if query != nil {
		if setUnMarshal {
			json.Unmarshal(query, &presenceHistroyJSONVariable)
		}
		return query
	}
	return nil
}
