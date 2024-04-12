package main

type config struct {
	nextLocURL *string
	prevLocURL *string
}

func main() {
	cfg := config{}
	StartRepl(&cfg)
}
