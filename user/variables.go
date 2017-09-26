package user

type devicesReportJSON struct {
	Type string `json:"type"`
	Case struct {
		Name     string   `json:"name"`
		Devices  []string `json:"devices"`
		Packages []string `json:"packages"`
	} `json:"case"`
}

type newTestCaseJSON struct {
	Type    string   `json:"type"`
	Code    string   `json:"code"`
	Name    string   `json:"name"`
	Devices []string `json:"devices"`
	Case    []struct {
		Loop  int      `json:"loop"`
		Steps []string `json:"steps"`
	} `json:"case"`
}

type savedTestCasesJSON struct {
	Name    string   `json:"name"`
	Devices []string `json:"devices"`
	Case    []struct {
		Loop  int      `json:"loop"`
		Steps []string `json:"steps"`
	} `json:"case"`
}
