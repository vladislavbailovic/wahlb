package main

import (
	"bytes"
	"fmt"
	"wahlb/internal"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/glamour"
)

func main() {
	b := internal.GetTestFile("su.html")
	buffer := bytes.NewBuffer(b)

	doc, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		panic(err)
	}

	conveter := md.NewConverter("", true, nil)
	/*
		mdown, err := conveter.ConvertString(doc.Find("body"))
		if err != nil {
			panic(err)
		}
	*/
	mdown := conveter.Convert(doc.Find("[itemprop='mainEntity']"))

	out, err := glamour.Render(mdown, "dark")
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
