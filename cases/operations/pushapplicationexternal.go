package operations

import "github.com/KMACEL/IITR/rest/workingset"

type PushApplicationExternal struct{
	ApplicationCode string
	DeviceID []string
}

func(p PushApplicationExternal) Start (){
	var(
		//workingsets workingset.Workingset
	)

	//workingsets.PushApplicationsExternal(p.ApplicationCode,workingset.NotNotifyUser,p.DeviceID...)
}