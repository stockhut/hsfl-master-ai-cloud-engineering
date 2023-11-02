package main

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/termcolor"
)

func main() {

	fgMagenta := termcolor.Fg(termcolor.FgMagenta)
	bgMagenta := termcolor.Bg(termcolor.BgMagenta)

	fmt.Printf("%s %s\n", fgMagenta("hello"), bgMagenta("world"))

	fmt.Println(fgMagenta(bgMagenta("hello world")))

	c := termcolor.
		NewBuilder().
		Text(termcolor.FgWhite).
		Background(termcolor.BgRed).
		Build()

	fmt.Println(c("hello world"))

	c = termcolor.
		NewBuilder().
		Background(termcolor.BgRed).
		Build()

	fmt.Println(c("hello world"))
}
