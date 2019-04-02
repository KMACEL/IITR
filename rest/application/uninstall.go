package application

import (
	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
██╗   ██╗███╗   ██╗██╗███╗   ██╗███████╗████████╗ █████╗ ██╗     ██╗
██║   ██║████╗  ██║██║████╗  ██║██╔════╝╚══██╔══╝██╔══██╗██║     ██║
██║   ██║██╔██╗ ██║██║██╔██╗ ██║███████╗   ██║   ███████║██║     ██║
██║   ██║██║╚██╗██║██║██║╚██╗██║╚════██║   ██║   ██╔══██║██║     ██║
╚██████╔╝██║ ╚████║██║██║ ╚████║███████║   ██║   ██║  ██║███████╗███████╗
 ╚═════╝ ╚═╝  ╚═══╝╚═╝╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚══════╝
*/

//Uninstall is
func (a Application) Uninstall(notifyUser bool, packageName []string, deviceID ...string) bool {
	setQueryAddress := uninstallLink()
	body := uninstallBody(notifyUser, packageName, deviceID)

	query, err := rest.Query{}.PutQuery(setQueryAddress, body, contentTypeJSON(), true)
	errc.ErrorCenter("Uninstall Application :", err)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			//json.Unmarshal(query, &responsePushApplicationJSONVariable)
			//todo : succes bilgisini kontrol et
			return true
		}
		return false
	}
	return false
}
