package main

import (
	"fmt"
	"net/http"

	. "github.com/stevelacy/daz"
)

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Println("listening on :3000")
	http.ListenAndServe(":3000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	links := H("link", Attr{"href": "https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css", "rel": "stylesheet"})

	meta := []func() string{
		H("meta", Attr{"charset": "UTF-8"}),
		H("meta", Attr{
			"name":    "viewport",
			"content": "width=device-width, initial-scale=1.0",
		}),
	}

	head := H("head", H("title", "Example Server"), meta, links)
	style := Attr{"style": "background: #efefef"}

	body := H("body", style, nested())
	html := H("html", head, body)
	w.Write([]byte(html()))
}

func navItems() []func() string {
	// get itmes from somewhere such as a database
	items := []func() string{H("li", "item one"), H("li", "item two")}

	// example runtime modification
	lastElement := H("li", "last item")
	items = append(items, lastElement)
	return items
}

func nested() func() string {
	nav := H("nav", navItems())
	return H(
		"div", Attr{"class": "bg-grey-50"},
		H("div", Attr{"class": "max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:py-16 lg:px-8 lg:flex lg:items-center lg:justify-between"},
			H("h2", Attr{"class": "text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl"},
				H("span", Attr{"class": "block"}, "Welcome to daz"),
				H("span", Attr{"class": "block text-indigo-600"}, "This example uses Tailwind CSS"),
			),
			H("div", Attr{"class": "mt-8 lex lg:mt-0 lg:flex-shrink-0"},
				H("div", Attr{"class": "inline-flex rounded-md shadow"},
					H("a", Attr{"href": "#", "class": "inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700"}, "Get Started")),
			),
			H("div", Attr{"class": "ml-3 inline-flex rounded-md shadow"},
				H("a", Attr{"href": "#", "class": "inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-indigo-600 bg-white hover:bg-indigo-50"}, "Learn More")),
			nav,
		),
	)
}
