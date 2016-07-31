package myapi

const (
	GREET = "hello"
)

func Greeting() string {
	return GREET
}

func Greet(name string) string {
	if name != "" {
		return GREET + " " + name
	}
	return Greeting()
}
