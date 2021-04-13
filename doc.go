/*
Package daz is a library for building composable HTML components in Go. It is a functional alternative to using templates, and allows for nested components/lists.

Daz is a "functional" alternative to using templates, and allows for nested components/lists
Also enables template-free server-side rendered components with support for nested lists. It is inspired by https://github.com/hyperhype/hyperscript


Basic usage:


	element := H("div", Attr{"class": "bg-grey-50"})

	html := H("html", element)

	w.Write([]byte(html()))


Lists can be easily created without needing to embed a `range / end` in a template:


	items := []HTML{H("li", "item one"), H("li", "item two")}

	element := H("ul", Attr{"class": "bg-grey-50"})

	div := H("div", element)

*/
package daz
