package rest

import (
	"net/url"
)

/*
note : use :
		u, _ := url.ParseRequestURI(api)
		u.Path = l
		urlStr := u.String() // 'https://api.com/user/'
		https://stackoverflow.com/questions/19253469/make-a-url-encoded-post-request-using-http-newrequest

*/

/*
██╗     ██╗███╗   ██╗██╗  ██╗███████╗
██║     ██║████╗  ██║██║ ██╔╝██╔════╝
██║     ██║██╔██╗ ██║█████╔╝ ███████╗
██║     ██║██║╚██╗██║██╔═██╗ ╚════██║
███████╗██║██║ ╚████║██║  ██╗███████║
╚══════╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
*/
//This page is the part that shows the links that the queries will use.
//It is designed in such a way that the administration is easy.

// Const is
const (
	APIScheme = "https"
	APIHost   = "api.ardich.com"
	APIPath   = "api/v3/"
)

const (
	oauth = "login/oauth"
)

const (
	contentType   = "Content-Type"
	authorization = "Authorization"

	authorizationKey          = "Basic ZnJvbnRlbmQ6" // data := []byte("frontend:") 	str := base64.StdEncoding.EncodeToString(data)
	contentTypeApplicationKey = "application/x-www-form-urlencoded"

	grantType    = "grant_type"
	refleshToken = "refresh_token"
	headerBearer = "Bearer "
)

var (
	scheme string
	host   string
	path   string
)

func loginLink() string {
	u := GetAPITemplate()
	u.Path = u.Path + oauth
	return u.String()
}

func connectBodyLink(userName string, password string) string {
	data := url.Values{}
	data.Add(grantType, "password")
	data.Add("username", userName)
	data.Add("password", password)
	return data.Encode()
}

func refleshTokenBodyLink() string {
	data := url.Values{}
	data.Add(grantType, refleshToken)
	data.Add(refleshToken, GetRefreshToken())
	return data.Encode()
}

// GetAPITemplate is
func GetAPITemplate() url.URL {
	if len(scheme) == 0{
		scheme=APIScheme
	}

	if len(host) == 0{
		host=APIHost
	}

	if len(path) == 0{
		path=APIPath
	}

	return url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path}
}

// SetScheme is
func SetScheme(schm string)  {
	scheme= schm
}

// SetHost is
func SetHost(hst string)  {
	host= hst
}

// SetPath is
func SetPath(pth string)  {
	path= pth
}

