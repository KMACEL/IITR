package rest

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
	ResponseNil      = "NIL Variable"
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
