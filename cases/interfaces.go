package cases

//Starter is
type Starter interface {
	Start()
}

//Test is
func Test(s Starter) {
	s.Start()
}

//Report is
func Report(s Starter) {
	s.Start()
}
