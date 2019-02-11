package drom

import (
	"encoding/json"

	"github.com/KMACEL/IITR/rest"
)

/*
 ██████╗ ██████╗ ███╗   ██╗███████╗██╗ ██████╗ ██╗   ██╗██████╗  █████╗ ████████╗██╗ ██████╗ ███╗   ██╗        ██╗     ██╗███████╗████████╗
██╔════╝██╔═══██╗████╗  ██║██╔════╝██║██╔════╝ ██║   ██║██╔══██╗██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║        ██║     ██║██╔════╝╚══██╔══╝
██║     ██║   ██║██╔██╗ ██║█████╗  ██║██║  ███╗██║   ██║██████╔╝███████║   ██║   ██║██║   ██║██╔██╗ ██║        ██║     ██║███████╗   ██║
██║     ██║   ██║██║╚██╗██║██╔══╝  ██║██║   ██║██║   ██║██╔══██╗██╔══██║   ██║   ██║██║   ██║██║╚██╗██║        ██║     ██║╚════██║   ██║
╚██████╗╚██████╔╝██║ ╚████║██║     ██║╚██████╔╝╚██████╔╝██║  ██║██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║        ███████╗██║███████║   ██║
 ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝     ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚══════╝╚═╝╚══════╝   ╚═╝
*/

// ConfigurationList is
func (d Drom) ConfigurationList(setUnMarshal bool, setVisualFlag bool) []byte {
	var q rest.Query

	queryLink := configurationListLink()

	query, errBuiltInAppList := q.GetQuery(queryLink, setVisualFlag)
	//errc.ErrorCenter(getBuiltInApplicationListErrorTag, errBuiltInAppList)
	_ = errBuiltInAppList
	if query != nil {

		return query
	}
	return nil
}

/*
 ██████╗ ███████╗████████╗        ██████╗ ██████╗  ██████╗ ███╗   ███╗         ██████╗ ██████╗ ███╗   ██╗███████╗██╗ ██████╗ ██╗   ██╗██████╗  █████╗ ████████╗██╗ ██████╗ ███╗   ██╗
██╔════╝ ██╔════╝╚══██╔══╝        ██╔══██╗██╔══██╗██╔═══██╗████╗ ████║        ██╔════╝██╔═══██╗████╗  ██║██╔════╝██║██╔════╝ ██║   ██║██╔══██╗██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║
██║  ███╗█████╗     ██║           ██║  ██║██████╔╝██║   ██║██╔████╔██║        ██║     ██║   ██║██╔██╗ ██║█████╗  ██║██║  ███╗██║   ██║██████╔╝███████║   ██║   ██║██║   ██║██╔██╗ ██║
██║   ██║██╔══╝     ██║           ██║  ██║██╔══██╗██║   ██║██║╚██╔╝██║        ██║     ██║   ██║██║╚██╗██║██╔══╝  ██║██║   ██║██║   ██║██╔══██╗██╔══██║   ██║   ██║██║   ██║██║╚██╗██║
╚██████╔╝███████╗   ██║           ██████╔╝██║  ██║╚██████╔╝██║ ╚═╝ ██║        ╚██████╗╚██████╔╝██║ ╚████║██║     ██║╚██████╔╝╚██████╔╝██║  ██║██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║
 ╚═════╝ ╚══════╝   ╚═╝           ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝         ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝     ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
*/

// GetDromConfiguration is
func (d Drom) GetDromConfiguration(configurationName string, setUnMarshal bool, setVisualFlag bool) string {
	var configurationListJSON ConfigurationListJSON
	query := d.ConfigurationList(false, false)
	if query != nil {
		json.Unmarshal(query, &configurationListJSON)
		for _, dromConfiguration := range configurationListJSON {
			if dromConfiguration.Name == configurationName {
				return dromConfiguration.ConfigurationID
			}
		}
		return "Not Found"
	}
	return "nil"
}
