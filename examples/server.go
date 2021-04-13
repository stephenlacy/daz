package main

import (
	"fmt"
	"net/http"

	. "github.com/stevelacy/daz"
)

// User is an example prop for a component
type User struct {
	Name string
}

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Println("listening on :3000")
	http.ListenAndServe(":3000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	title := "Example Server"
	description := "Welcome to daz"

	user := User{Name: "Daz"}

	links := H("link", Attr{
		"href": "https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css",
		"rel":  "stylesheet",
	})

	meta := []HTML{
		H("meta", Attr{"charset": "UTF-8"}),
		H("meta", Attr{
			"name":    "viewport",
			"content": "width=device-width, initial-scale=1.0",
		}),
	}

	head := H("head", H("title", title), meta, links)
	style := Attr{"style": "background: #efefef"}

	body := H(
		"body",
		style,
		AppComponent(user, description),
	)
	html := H("html", head, body)
	w.Write([]byte(html()))
}

func navItems(user User) []HTML {
	// get itmes from somewhere such as a database
	items := []HTML{H("li", "item one"), H("li", "item two")}

	// example runtime modification
	lastElement := H("li", user.Name)
	items = append(items, lastElement)
	return items
}

// AppComponent is a daz component. It returns a daz.H func
func AppComponent(user User, description string) HTML {
	nav := H("nav", navItems(user))
	return H(
		"div", Attr{"class": "bg-grey-50"},
		H("div", Attr{"class": "max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:py-16 lg:px-8 lg:flex lg:items-center lg:justify-between"},
			H("h2", Attr{"class": "text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl"},
				H("span", Attr{"class": "block"}, description),
				H("span", Attr{"class": "block text-indigo-600"}, "This example uses Tailwind CSS"),
			),
			H("div", Attr{"class": "mt-8 lex lg:mt-0 lg:flex-shrink-0"},
				H("div", Attr{"class": "inline-flex rounded-md shadow"},
					H("a", Attr{"href": "#", "class": "inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700"}, "Get Started")),
			),
			H("div", Attr{"class": "ml-3 inline-flex rounded-md shadow"},
				H("a", Attr{"href": "#", "class": "inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-indigo-600 bg-white hover:bg-indigo-50"}, "Learn More")),
			H("div",
				Attr{"class": "mt-1 flex rounded-md shadow-sm"},
				H("span", Attr{"class": "inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm"}, "http://"),
				H("input", Attr{"type": "text", "name": "test", "class": "focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-none rounded-r-md sm:text-sm border-gray-300", "placeholder": "input's value"}),
			),
			nav,
		),
	)
}
