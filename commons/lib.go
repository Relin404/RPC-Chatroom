package commons

type Args struct {
	Message string
	Name    string
}

func GetServerAddress() string {
	return "0.0.0.0:7422"
}
