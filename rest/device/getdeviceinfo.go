package device

import (
	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/rest"
)

/*
██████╗ ██████╗ ███████╗███████╗███████╗███╗   ██╗ ██████╗███████╗        ██╗███╗   ██╗███████╗ ██████╗
██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝████╗  ██║██╔════╝██╔════╝        ██║████╗  ██║██╔════╝██╔═══██╗
██████╔╝██████╔╝█████╗  ███████╗█████╗  ██╔██╗ ██║██║     █████╗          ██║██╔██╗ ██║█████╗  ██║   ██║
██╔═══╝ ██╔══██╗██╔══╝  ╚════██║██╔══╝  ██║╚██╗██║██║     ██╔══╝          ██║██║╚██╗██║██╔══╝  ██║   ██║
██║     ██║  ██║███████╗███████║███████╗██║ ╚████║╚██████╗███████╗        ██║██║ ╚████║██║     ╚██████╔╝
╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═══╝ ╚═════╝╚══════╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

//PresenceInfo is
func (d Device) PresenceInfo(setDeviceID string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := presenceInfoLink(setDeviceID)

	query, errorPresenceInfo := q.GetQuery(setQueryAddress, visualFlag)
	errc.ErrorCenter(presenceInfoErrorTag, errorPresenceInfo)

	if query != nil {
		if setUnMarshal {
			//json.Unmarshal(query, &presenceInfoJSONVariable)
		}
		return query
	}
	return nil
}

/*
 █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗        ██╗███╗   ██╗███████╗ ██████╗
██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║        ██║████╗  ██║██╔════╝██╔═══██╗
███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║        ██║██╔██╗ ██║█████╗  ██║   ██║
██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║        ██║██║╚██╗██║██╔══╝  ██║   ██║
██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║        ██║██║ ╚████║██║     ╚██████╔╝
╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

//ApplicationInfo is
func (d Device) ApplicationInfo(setDeviceID string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := applicationInfoLink(setDeviceID)

	query, _ := q.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				//json.Unmarshal(query, &applicationInfoJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}

/*
 ██████╗ ███████╗        ██████╗ ██████╗  ██████╗ ███████╗██╗██╗     ███████╗        ██╗███╗   ██╗███████╗ ██████╗
██╔═══██╗██╔════╝        ██╔══██╗██╔══██╗██╔═══██╗██╔════╝██║██║     ██╔════╝        ██║████╗  ██║██╔════╝██╔═══██╗
██║   ██║███████╗        ██████╔╝██████╔╝██║   ██║█████╗  ██║██║     █████╗          ██║██╔██╗ ██║█████╗  ██║   ██║
██║   ██║╚════██║        ██╔═══╝ ██╔══██╗██║   ██║██╔══╝  ██║██║     ██╔══╝          ██║██║╚██╗██║██╔══╝  ██║   ██║
╚██████╔╝███████║        ██║     ██║  ██║╚██████╔╝██║     ██║███████╗███████╗        ██║██║ ╚████║██║     ╚██████╔╝
 ╚═════╝ ╚══════╝        ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚══════╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

// OSProfileInfo is
func (d Device) OSProfileInfo(setDeviceID string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := osProfileInfoLink(setDeviceID)

	query, _ := q.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				//json.Unmarshal(query, &osProfileInfoJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}

/*
██╗███╗   ██╗███████╗████████╗ █████╗ ███╗   ██╗████████╗         █████╗ ██████╗ ██████╗ ██╗     ██╗ ██████╗ █████╗ ████████╗██╗ ██████╗ ███╗   ██╗        ██╗███╗   ██╗███████╗ ██████╗
██║████╗  ██║██╔════╝╚══██╔══╝██╔══██╗████╗  ██║╚══██╔══╝        ██╔══██╗██╔══██╗██╔══██╗██║     ██║██╔════╝██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║        ██║████╗  ██║██╔════╝██╔═══██╗
██║██╔██╗ ██║███████╗   ██║   ███████║██╔██╗ ██║   ██║           ███████║██████╔╝██████╔╝██║     ██║██║     ███████║   ██║   ██║██║   ██║██╔██╗ ██║        ██║██╔██╗ ██║█████╗  ██║   ██║
██║██║╚██╗██║╚════██║   ██║   ██╔══██║██║╚██╗██║   ██║           ██╔══██║██╔═══╝ ██╔═══╝ ██║     ██║██║     ██╔══██║   ██║   ██║██║   ██║██║╚██╗██║        ██║██║╚██╗██║██╔══╝  ██║   ██║
██║██║ ╚████║███████║   ██║   ██║  ██║██║ ╚████║   ██║           ██║  ██║██║     ██║     ███████╗██║╚██████╗██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║        ██║██║ ╚████║██║     ╚██████╔╝
╚═╝╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝           ╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

// InstantApplicationInfo is
func (d Device) InstantApplicationInfo(setDeviceID string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := instantApplicationInfoLink(setDeviceID)

	query, _ := q.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				//json.Unmarshal(query, &instantApplicationInfoJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}

/*
███╗   ███╗ ██████╗ ██████╗ ██╗██╗   ██╗███████╗██████╗ ███████╗███████╗        ██╗███╗   ██╗███████╗ ██████╗
████╗ ████║██╔═══██╗██╔══██╗██║██║   ██║██╔════╝██╔══██╗██╔════╝██╔════╝        ██║████╗  ██║██╔════╝██╔═══██╗
██╔████╔██║██║   ██║██║  ██║██║██║   ██║█████╗  ██████╔╝███████╗█████╗          ██║██╔██╗ ██║█████╗  ██║   ██║
██║╚██╔╝██║██║   ██║██║  ██║██║╚██╗ ██╔╝██╔══╝  ██╔══██╗╚════██║██╔══╝          ██║██║╚██╗██║██╔══╝  ██║   ██║
██║ ╚═╝ ██║╚██████╔╝██████╔╝██║ ╚████╔╝ ███████╗██║  ██║███████║███████╗        ██║██║ ╚████║██║     ╚██████╔╝
╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝        ╚═╝╚═╝  ╚═══╝╚═╝      ╚═════╝
*/

// ModiverseInfo is
func (d Device) ModiverseInfo(setDeviceID string, setUnMarshal bool, visualFlag bool) []byte {
	setQueryAddress := modiverseInfoLink(setDeviceID)
	query, _ := q.GetQuery(setQueryAddress, visualFlag)

	if query != nil {
		if string(query) != rest.ResponseNotFound {
			if setUnMarshal {
				//json.Unmarshal(query, &instantApplicationInfoJSONVariable)
			}
			return query
		}
		return []byte(rest.ResponseNotFound)
	}
	return nil
}
