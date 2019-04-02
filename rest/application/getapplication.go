package application

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/KMACEL/IITR/rest"
)

// GetAllApplications is
func (a Application) GetAllApplications() []byte {

	setAddress := getAllApplicationsLink()
	query, _ := rest.Query{}.GetQuery(setAddress, false)
	fmt.Println(string(query))
	if query != nil {
		return query
	}
	return nil
}

// GetApplications is
func (a Application) GetApplications(packageName string, versionCode string) string {
	setAddress := getApplicationsLink(url.QueryEscape(packageName))
	query, _ := rest.Query{}.GetQuery(setAddress, false)

	var responseApplicationList ResponseApplicationList
	if query != nil {
		json.Unmarshal(query, &responseApplicationList)
		for _, control := range responseApplicationList.Content {
			if control.PackageName == packageName && control.VersionCode == versionCode {
				return control.Code
			}
		}
		return "Not Found"
	}
	return "Query is Null"
}
