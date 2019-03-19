package application

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

// Application is
type Application struct {
}

/*
     ██╗███████╗ ██████╗ ███╗   ██╗███████╗
     ██║██╔════╝██╔═══██╗████╗  ██║██╔════╝
     ██║███████╗██║   ██║██╔██╗ ██║███████╗
██   ██║╚════██║██║   ██║██║╚██╗██║╚════██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║███████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝
*/

// ResponseApplicationList is
type ResponseApplicationList struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Content []struct {
		PackageName   string      `json:"packageName"`
		Notify        bool        `json:"notify"`
		Name          string      `json:"name"`
		VersionCode   string      `json:"versionCode"`
		VersionName   string      `json:"versionName"`
		Type          interface{} `json:"type"`
		Size          int         `json:"size"`
		Code          string      `json:"code"`
		Desc          string      `json:"desc"`
		Version       string      `json:"version"`
		OsVersion     string      `json:"osVersion"`
		Location      string      `json:"location"`
		DeviceModel   string      `json:"deviceModel"`
		URL           string      `json:"url"`
		Token         string      `json:"token"`
		ExternalURL   string      `json:"externalUrl"`
		Hidden        bool        `json:"hidden"`
		DownloadCount int         `json:"downloadCount"`
		RatingCount   int         `json:"ratingCount"`
		CreateDate    int64       `json:"createDate"`
		Rating        int         `json:"rating"`
		AppType       string      `json:"appType"`
		Category      struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"category"`
		Icons       []interface{} `json:"icons"`
		Videos      []interface{} `json:"videos"`
		Screenshots []interface{} `json:"screenshots"`
		Comments    []interface{} `json:"comments"`
		Tags        []interface{} `json:"tags"`
		Permissions interface{}   `json:"permissions"`
		Links       []interface{} `json:"links"`
	} `json:"content"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}
