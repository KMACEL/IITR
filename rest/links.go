package rest

import "net/url"

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

const (
	login       = "https://api.ardich.com/api/v3/login/oauth"
	device      = "https://api.ardich.com:443/api/v3/device/"
	locationMap = "device-location-map"
	downloaded  = "/apps?type=downloaded"
	modePolicy  = "/current-and-active-profile"
)

const (
	contentType   = "Content-Type"
	authorization = "Authorization"

	authorizationKey          = "Basic ZnJvbnRlbmQ6"
	contentTypeApplicationKey = "application/x-www-form-urlencoded"

	grantType        = "grant_type="
	passwordUsername = "password&&username="
	passwordEntry    = "&&password="
	refleshToken     = "refresh_token&refresh_token="
	headerBearer     = "Bearer "
)

func loginLink() string {
	return login
}

func connectBodyLink(userName string, password string) string {
	return grantType + passwordUsername + userName + passwordEntry + url.QueryEscape(password)
}

func refleshTokenBodyLink() string {
	return grantType + refleshToken + GetRefreshToken()
}
