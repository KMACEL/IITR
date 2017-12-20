package operations

import "github.com/KMACEL/IITR/rest/workingset"

//PushApplicationExternal is
type PushApplicationExternal struct {
	ApplicationCode string
	DeviceID        []string
	URL             string
}

//Start is
func (p PushApplicationExternal) Start() {
	var (
		workingsets workingset.Workingset
	)

	workingsets.PushApplicationsExternal(p.ApplicationCode, p.URL, workingset.NotNotifyUser, p.DeviceID...)
}
