package rest

import "errors"

/*
██╗   ██╗ █████╗ ██████╗ ██╗ █████╗ ██████╗ ██╗     ███████╗███████╗
██║   ██║██╔══██╗██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝
██║   ██║███████║██████╔╝██║███████║██████╔╝██║     █████╗  ███████╗
╚██╗ ██╔╝██╔══██║██╔══██╗██║██╔══██║██╔══██╗██║     ██╔══╝  ╚════██║
 ╚████╔╝ ██║  ██║██║  ██║██║██║  ██║██████╔╝███████╗███████╗███████║
  ╚═══╝  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝╚══════╝
*/

// After login is done, a json is sent back. It contains variables handling the json parameter
var (
	getLogin loginJSON
)

/*
 ██████╗ ██████╗ ███╗   ██╗███████╗████████╗
██╔════╝██╔═══██╗████╗  ██║██╔════╝╚══██╔══╝
██║     ██║   ██║██╔██╗ ██║███████╗   ██║
██║     ██║   ██║██║╚██╗██║╚════██║   ██║
╚██████╗╚██████╔╝██║ ╚████║███████║   ██║
 ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝   ╚═╝
*/

// These constants are used throughout the program. It has been used to make variable types more meaningful.
// ResponseOK: This constant is used to check if we have a question when we are back positive, ie "200 OK".
// ResponseNotFound: This constant is used to check if we have a negative response, ie "404 Not Found" when we make a query.
// ResponseCreated: This constant is used to check if "201 Created" is coming back when we made a query.
// ResponseNil: This constant is used to return a blank when we have a question.
const (
	ResponseOK       = "200 OK"
	ResponseNotFound = "404 Not Found"
	ResponseCreated  = "201 Created"
	ResponseNil      = "Null Value"
)

// Response Code
const (
	ResponseOKCode            = 200
	ResponseCreatedCode       = 201
	ResponseBadRequestCode    = 400
	ResponseUnauthorizedCode  = 401
	ResponseForbiddenCode     = 403
	ResponseNotFoundCode      = 404
	ResponseServerProblemCode = 500
)

// Errors
var (
	ErrorResponseBadRequestCode400   = errors.New("Request is 40 Bad Request. Please check your request")
	ErrorResponseUnauthorizedCode401 = errors.New("Request is 401 Unauthorized. Please check the login information")
	ErrorResponseForbiddenCode403    = errors.New("Request is 403 Forbidden. You are not authorized for this query")
	ErrorNotFound404                 = errors.New("Request is 404 Not Found. Please check variables, queries, links and other parameters")
	ErrorServerProblemCode500        = errors.New("Request is 500 Server Problem. There is a server problem. Please try later")
	ErrorElseProblem                 = errors.New("No problem could be detected. Please check the information. If the problem is not resolved, consult the program owner")
	ErrorResponseNil                 = errors.New("Null Response Message")
	ErrorResponseNilRequest          = errors.New("Null Response Request Message")
)

// Query Type
const (
	POST = "POST"
	GET  = "GET"
	PUT  = "PUT"
)

// These constants are used throughout the program. They have functions such as displaying yawning values.
// Visible: Show the user without processing the incoming json value.
// Invisible: The incoming json value does not represent the user.
// OKMarshal: Processes automatically when data arrives. However, since this process is address-based, there may be data overlap. This is why it is necessary to pay attention.
// NOMarshal: If this option is selected, the data will be sent directly without processing. The incoming value is processed in the returned field.
const (
	Visible   = true
	Invisible = false
	OKMarshal = true
	NOMarshal = false
)

//This section is the online - offline status check of the Presence value taken from the device.
const (
	Online  = "ONLINE"
	Offline = "OFFLINE"
)

//
const (
	NoApplication = "NOAPPLICATION"
	Running       = "RUNNING"
	NotRunning    = "NOTRUNNING"
	Blocked       = "BLOCKED"
	NotBlocked    = "NOTBLOCKED"
	UnKnow        = "UNKNOW"
)

// Rest -> Query Constant
const (
	requestGetTag = "Request Get : "
	doGetTag      = "Do Get : "
	bodyGetTag    = "Body Get : "

	requestPostTag = "Request Get :"
	doPostTag      = "Do Get : "
	bodyPostTag    = "Body Get : "

	requestPutTag = "Request Get : "
	doPutTag      = "Do Get : "
	bodyPutTag    = "Body Get : "
)

/*
     ██╗███████╗ ██████╗ ███╗   ██╗
     ██║██╔════╝██╔═══██╗████╗  ██║
     ██║███████╗██║   ██║██╔██╗ ██║
██   ██║╚════██║██║   ██║██║╚██╗██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║
╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
*/

// When it is a login operation, the data comes from a JSON type. It's a big thing for us to use
// and interpret this data. This is done with the UnMarshal method and sends to this part.
// "Access Token", "Reflesh Token" and "Expires In" are the information we use very much.
type loginJSON struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}
