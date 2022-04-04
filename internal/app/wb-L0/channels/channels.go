package channels

var (
	ConnectedToNats = make(chan bool)
	ConnectedToDb   = make(chan bool)
)
