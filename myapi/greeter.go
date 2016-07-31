package myapi

const (
	GREET = "hello"
)

func Greeting() string {
	return GREET
}

func Greet(name string) string{
	return GREET + " " + name
}
