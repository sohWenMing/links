package parsing

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var datatypes map[int]string = map[int]string{
	0: "ErrorNode",
	1: "TextNode",
	2: "DocumentNode",
	3: "ElementNode",
	4: "CommentNode",
	5: "DoctypeNode",
	6: "RawNode",
}

type Link struct {
	Href, Text string
}

func (l *Link) String() string {
	return fmt.Sprint("\n" +
		"Href: " + l.Href + "\n" +
		"Text: " + l.Text)
}

func GetLinks(input string) (links []Link, err error) {
	node, err := ParseHtmlToDoc(input)
	if err != nil {
		return []Link{}, err
	}
	return Visit(node), nil
}

func ParseHtmlToDoc(input string) (node *html.Node, err error) {
	node, err = html.Parse(strings.NewReader(input))
	return node, err
}

func Visit(entryNode *html.Node) (links []Link) {

	returnedLinks := []Link{}
	for currentNode := entryNode.FirstChild; currentNode != nil; currentNode = currentNode.NextSibling {
		if stripString(currentNode.Data) == "a" {
			for _, attr := range currentNode.Attr {
				if attr.Key == "href" {
					newLink := Link{stripString(attr.Val), stripString(currentNode.FirstChild.Data)}
					returnedLinks = append(returnedLinks, newLink)
				}
			}

		}
		returnedFromNode := Visit(currentNode)
		for _, node := range returnedFromNode {
			returnedLinks = append(returnedLinks, node)
		}
	}
	return returnedLinks
}

func CheckArgsAndGetHTML(args []string) (htmlString string, err error) {
	err = checkArgsLength(args)
	if err != nil {
		return "", err
	}
	htmlString, err = getHTML(args[1])
	return htmlString, err

}
func checkArgsLength(args []string) error {
	if len(args) != 2 {
		return errors.New("number of arguments cannot be more than 2")
	}
	return nil
}

func getHTML(filename string) (htmlString string, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func stripString(input string) string {
	return strings.TrimSpace(input)
}
