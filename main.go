// main.go

package main

func main() {
	a := App{}
	a.Initialize(
		"go-app",
		"UMErR4NLBjRJqYEUaUs0qw7Ap",
		"postgres")

	a.Run(":8010")
}
