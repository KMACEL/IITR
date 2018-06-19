package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/KMACEL/IITR/errc"
	"github.com/KMACEL/IITR/logc"
)

/*
 ██████╗ ██████╗ ███╗   ██╗███╗   ██╗███████╗ ██████╗████████╗
██╔════╝██╔═══██╗████╗  ██║████╗  ██║██╔════╝██╔════╝╚══██╔══╝
██║     ██║   ██║██╔██╗ ██║██╔██╗ ██║█████╗  ██║        ██║
██║     ██║   ██║██║╚██╗██║██║╚██╗██║██╔══╝  ██║        ██║
╚██████╗╚██████╔╝██║ ╚████║██║ ╚████║███████╗╚██████╗   ██║
 ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═══╝╚══════╝ ╚═════╝   ╚═╝
*/

// Connect is In order to use the Connect Rest APIs, it performs login util to "Tenant".
// It takes two parameters;
//    1: userName:
//          UserName, IoT - You are logged in to the Ignite platform.
//    2: password:
//          Password, IoT - prompts you for the password information that you have registered
//          for the Ignite platform.
// This process performs the Post method. Information from the user is sent to the method post section.
// The header section is important here. This is given as a constant.
// This function returns "true" if the Connect operation has been performed. Sends "false" if the message
// is empty and "false" if it gets an error.
func Connect(userName string, password string) bool {
	setQueryAddress := loginLink()
	setBody := connectBodyLink(userName, password)

	setHeader := make(map[string]string)
	setHeader[contentType] = contentTypeApplicationKey
	setHeader[authorization] = authorizationKey
	visualFlag := Invisible

	query, errConnectPost := Query{}.PostQuery(setQueryAddress, setBody, setHeader, visualFlag)
	errc.ErrorCenter("Connect", errConnectPost)

	if query != nil {
		if string(query) != ResponseNotFound {
			json.Unmarshal(query, &getLogin)
			go tokenControl()
			log.Println("IoT - Ignite Connection : OK...")
			logc.ConnectionPrint("Connection OK - User Name : %s \nGet DATA : %s", userName, getLogin)
			return true
		}
		log.Println("IoT - Ignite Connection : NO - Response Not Found -404-!")
		logc.ConnectionPrint("Connection Response Not Found 404 - User Name : %s \nGet DATA : %s", userName, getLogin)
		return false
	}
	log.Println("IoT - Ignite Connection : NO - Query is nil ! Check Username - Password !")
	logc.ConnectionPrint("Connection Query is nil - User Name : %s \nGet DATA : %s", userName, getLogin)
	return false
}

/*
 ██████╗ ███████╗████████╗     █████╗  ██████╗ ██████╗███████╗███████╗    ████████╗ ██████╗ ██╗  ██╗███████╗███╗   ██╗
██╔════╝ ██╔════╝╚══██╔══╝    ██╔══██╗██╔════╝██╔════╝██╔════╝██╔════╝    ╚══██╔══╝██╔═══██╗██║ ██╔╝██╔════╝████╗  ██║
██║  ███╗█████╗     ██║       ███████║██║     ██║     █████╗  ███████╗       ██║   ██║   ██║█████╔╝ █████╗  ██╔██╗ ██║
██║   ██║██╔══╝     ██║       ██╔══██║██║     ██║     ██╔══╝  ╚════██║       ██║   ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╗██║
╚██████╔╝███████╗   ██║       ██║  ██║╚██████╗╚██████╗███████╗███████║       ██║   ╚██████╔╝██║  ██╗███████╗██║ ╚████║
 ╚═════╝ ╚══════╝   ╚═╝       ╚═╝  ╚═╝ ╚═════╝ ╚═════╝╚══════╝╚══════╝       ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝
*/

// GetAccessToken is; After the Connect is done, it sends back the Acces Token value.
// All subsequent query util will use this value. Anyone who wishes to know about
//Connect can take this value with the help of this function.
func GetAccessToken() string {
	return getLogin.AccessToken
}

/*
 ██████╗ ███████╗████████╗    ██████╗ ███████╗███████╗██████╗ ███████╗███████╗██╗  ██╗    ████████╗ ██████╗ ██╗  ██╗███████╗███╗   ██╗
██╔════╝ ██╔════╝╚══██╔══╝    ██╔══██╗██╔════╝██╔════╝██╔══██╗██╔════╝██╔════╝██║  ██║    ╚══██╔══╝██╔═══██╗██║ ██╔╝██╔════╝████╗  ██║
██║  ███╗█████╗     ██║       ██████╔╝█████╗  █████╗  ██████╔╝█████╗  ███████╗███████║       ██║   ██║   ██║█████╔╝ █████╗  ██╔██╗ ██║
██║   ██║██╔══╝     ██║       ██╔══██╗██╔══╝  ██╔══╝  ██╔══██╗██╔══╝  ╚════██║██╔══██║       ██║   ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╗██║
╚██████╔╝███████╗   ██║       ██║  ██║███████╗██║     ██║  ██║███████╗███████║██║  ██║       ██║   ╚██████╔╝██║  ██╗███████╗██║ ╚████║
 ╚═════╝ ╚══════╝   ╚═╝       ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝       ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝
*/

// GetRefreshToken is ; There is a period of Acces Token. When this time expires, the automatically
// executed util are interrupted and the program fails. You need to do the Connect again.
// But the Connect operation also gives us the Reflesh Token value. This value is used to automatically
// submit a new Access Token when the time expires. This function is used to rotate the desired value.
func GetRefreshToken() string {
	return getLogin.RefreshToken
}

/*
 ██████╗ ███████╗████████╗    ███████╗██╗  ██╗██████╗ ██╗██████╗ ███████╗███████╗██╗███╗   ██╗
██╔════╝ ██╔════╝╚══██╔══╝    ██╔════╝╚██╗██╔╝██╔══██╗██║██╔══██╗██╔════╝██╔════╝██║████╗  ██║
██║  ███╗█████╗     ██║       █████╗   ╚███╔╝ ██████╔╝██║██████╔╝█████╗  ███████╗██║██╔██╗ ██║
██║   ██║██╔══╝     ██║       ██╔══╝   ██╔██╗ ██╔═══╝ ██║██╔══██╗██╔══╝  ╚════██║██║██║╚██╗██║
╚██████╔╝███████╗   ██║       ███████╗██╔╝ ██╗██║     ██║██║  ██║███████╗███████║██║██║ ╚████║
 ╚═════╝ ╚══════╝   ╚═╝       ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝╚═╝╚═╝  ╚═══╝
*/

// getExpiresIn is; After this function is Connect, the Access Token returns the usage time.
func getExpiresIn() int {
	return getLogin.ExpiresIn
}

/*
████████╗ ██████╗ ██╗  ██╗███████╗███╗   ██╗     ██████╗ ██████╗ ███╗   ██╗████████╗██████╗  ██████╗ ██╗
╚══██╔══╝██╔═══██╗██║ ██╔╝██╔════╝████╗  ██║    ██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██╔═══██╗██║
   ██║   ██║   ██║█████╔╝ █████╗  ██╔██╗ ██║    ██║     ██║   ██║██╔██╗ ██║   ██║   ██████╔╝██║   ██║██║
   ██║   ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╗██║    ██║     ██║   ██║██║╚██╗██║   ██║   ██╔══██╗██║   ██║██║
   ██║   ╚██████╔╝██║  ██╗███████╗██║ ╚████║    ╚██████╗╚██████╔╝██║ ╚████║   ██║   ██║  ██║╚██████╔╝███████╗
   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝     ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝
*/

// tokenControl is ;This function is great. This function is started at runtime as Connect and
// controls the Access Token time on the backplane. At the end of the time, using the Reflesh Token value,
// zeroes the Access Token value.
// This function performs these util when the remaining time is 750 seconds.
// The first stage starts with a countdown from the value of getExpiresIn.
// Then it performs Post Query for the reset operation when the specified condition is fulfilled.
// This function performs continuous control and renewal until the program is closed
func tokenControl() {
retry:
	for i := getExpiresIn(); i > 750; i-- {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Refreshing Token")
	fmt.Println("exp : ", getExpiresIn(), " reflesh : ", GetRefreshToken())
	setQueryAddress := loginLink()
	setBody := refleshTokenBodyLink()

	setHeader := make(map[string]string)
	setHeader[contentType] = contentTypeApplicationKey
	setHeader[authorization] = authorizationKey
	visualFlag := Visible

	query, errRefreshToken := Query{}.PostQuery(setQueryAddress, setBody, setHeader, visualFlag)
	errc.ErrorCenter("Refresh Token", errRefreshToken)
	json.Unmarshal(query, &getLogin)
	goto retry
}
