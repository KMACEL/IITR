package user

import (
	"fmt"
	"os/exec"

	"github.com/KMACEL/IITR/timop"
)

func welcome() {

	formatText("Normal \\e[94mNormal \\e[49m")
	formatPicture(timop.Random(0, 3))

}

func formatText(commandString string) {
	out, err := exec.Command("echo", "-e", commandString).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}

func formatPicture(pictureNumber int) {

	if pictureNumber == 0 {
		fmt.Print(
			"██╗    ██╗███████╗██╗      ██████╗ ██████╗ ███╗   ███╗███████╗\n",
			"██║    ██║██╔════╝██║     ██╔════╝██╔═══██╗████╗ ████║██╔════╝\n",
			"██║ █╗ ██║█████╗  ██║     ██║     ██║   ██║██╔████╔██║█████╗  \n",
			"██║███╗██║██╔══╝  ██║     ██║     ██║   ██║██║╚██╔╝██║██╔══╝  \n",
			"╚███╔███╔╝███████╗███████╗╚██████╗╚██████╔╝██║ ╚═╝ ██║███████╗\n",
			" ╚══╝╚══╝ ╚══════╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝\n",
			"\n",
			"\t\t ████████╗ ██████╗ \n",
			"\t\t ╚══██╔══╝██╔═══██╗\n",
			"\t\t    ██║   ██║   ██║\n",
			"\t\t    ██║   ██║   ██║\n",
			"\t\t    ██║   ╚██████╔╝\n",
			"\t\t    ╚═╝    ╚═════╝ \n",
			"\n",
			"\t      ██╗██╗████████╗██████╗       \n",
			"\t      ██║██║╚══██╔══╝██╔══██╗      \n",
			"\t█████╗██║██║   ██║   ██████╔╝█████╗\n",
			"\t╚════╝██║██║   ██║   ██╔══██╗╚════╝\n",
			"\t      ██║██║   ██║   ██║  ██║      \n",
			"\t      ╚═╝╚═╝   ╚═╝   ╚═╝  ╚═╝\n\n\n",
		)

	} else if pictureNumber == 1 {
		fmt.Print(
			"____________________________________________________\n",
			"|.==================================================,|\n",
			"||  WELCOME TO THE IITR WELCOME TO THE IITR         ||\n",
			"||  WELCOME TO THE IITR WELCOME TO THE IITR         ||\n",
			"||  WELCOME TO THE IITR WELCOME TO THE IITR         ||\n",
			"||  W .----.TO T                                    ||\n",
			"||   / ><   \\  /                                    ||\n",
			"||  |        |/\\                                    ||\n",
			"||   \\______//\\/                                    ||\n",
			"||   _(____)/ /                                     ||\n",
			"||__/ ,_ _  _/______________________________________||\n",
			"'===\\___\\_) |========================================'\n",
			"    |______|\n",
			"    |  ||  |\n",
			"    |__||__|\n",
			"    (__)(__)\n\n\n")
	} else if pictureNumber == 2 {
		fmt.Print("^^          ..                                       ..\n",
			"            []                                       []\n",
			"          .:[]:_          ^^                       ,:[]:.\n",
			"        .: :[]: :-.=====WELCOME TO THE IITR=====,-: :[]: :.\n",
			"      .: : :[]: : :`._                       ,.': : :[]: : :.\n",
			"    .: : : :[]: : : : :-._               _,-: : : : :[]: : : :.\n",
			"_..: : : : :[]: : : : : : :-._________.-: : : : : : :[]: : : : :-._\n",
			"_:_:_:_:_:_:[]:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:[]:_:_:_:_:_:_\n",
			"!!!!!!!!!!!![]!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!![]!!!!!!!!!!!!!\n",
			"^^^^^^^^^^^^[]^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^[]^^^^^^^^^^^^^\n",
			"            []                                       []\n",
			"            []                                       []\n",
			"            []                                       []\n",
			" ~~^-~^_~^~/  \\~^-~^~_~^-~_^~-^~_^~~-^~_~^~-~_~-^~_^/  \\~^-~_~^-~~-\n",
			"~ _~~- ~^-^~-^~~- ^~_^-^~~_ -~^_ -~_-~~^- _~~_~-^_ ~^-^~~-_^-~ ~^\n",
			"~ ^- _~~_-  ~~ _ ~  ^~  - ~~^ _ -  ^~-  ~ _  ~~^  - ~_   - ~^_~\n",
			" ~-  ^_  ~^ -  ^~ _ - ~^~ _   _~^~-  _ ~~^ - _ ~ - _ ~~^ -\n",
			"    ~^ -_ ~^^ -_ ~ _ - _ ~^~-  _~ -_   ~- _ ~^ _ -  ~ ^-\n",
			"        ~^~ - _ ^ - ~~~ _ - _ ~-^ ~ __- ~_ - ~  ~^_-\n",
			"            ~ ~- ^~ -  ~^ -  ~ ^~ - ~~  ^~ - ~\n\n\n")
	}

}
