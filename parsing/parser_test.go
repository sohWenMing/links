package parsing

import (
	"slices"
	"testing"
)

func TestParsing(t *testing.T) {
	type test struct {
		filePath string
		expected []Link
	}
	tests := []test{
		{
			"../ex1.html",
			[]Link{
				{
					"/other-page",
					"A link to another page",
				},
			},
		},
		{
			"../ex2.html",
			[]Link{
				{
					"https://www.twitter.com/joncalhoun",
					"Check me out on twitter",
				},
				{
					"https://github.com/gophercises",
					"Gophercises is on",
				},
			},
		},
		{
			"../ex3.html",
			[]Link{
				{
					"#",
					"Login",
				},
				{
					"/lost",
					"Lost? Need help?",
				},
				{
					"https://twitter.com/marcusolsson",
					"@marcusolsson",
				},
			},
		},
		{
			"../ex4.html",
			[]Link{
				{
					"/dog-cat",
					"dog cat",
				},
			},
		},
	}
	for _, test := range tests {
		args := []string{"holder", test.filePath}
		html, err := CheckArgsAndGetHTML(args)
		if err != nil {
			t.Error(err)
		}
		doc, err := ParseHtmlToDoc(html)
		if err != nil {
			t.Error(err)
		}
		got := Visit(doc)
		if !slices.Equal(test.expected, got) {
			t.Errorf("got %v:\nwant%v\n", got, test.expected)
		}
	}
}
