package main

import (
	"fmt"
	"wahlb/internal"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/charmbracelet/glamour"
)

func main() {
	b := internal.GetTestFile("su.html")
	conveter := md.NewConverter("", true, nil)
	mdown, err := conveter.ConvertString(string(b))
	if err != nil {
		fmt.Println("Error:", err)
	}

	out, err := glamour.Render(mdown, "dark")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(out)
}
