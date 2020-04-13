package main

func main() {
	cli()
	if err := run(); err != nil {
		panic(err)
	}
}
