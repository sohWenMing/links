package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sohWenMing/links/parsing"
)

func main() {
	args := os.Args
	htmlString, err := parsing.CheckArgsAndGetHTML(args)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := parsing.ParseHtmlToDoc(htmlString)
	if err != nil {
		log.Fatal(err)
	}
	returnedLinks := parsing.Visit(doc)
	fmt.Println("##### PRINTING LINKS #####")
	for _, link := range returnedLinks {
		fmt.Println("link: ", link.String())
	}

}

/*
	If I want to recursively go through the tree
	the parent has a first child
	i want to go through everything in the first child
	then after that, i would want to go to the next child
*/
