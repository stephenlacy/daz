# daz
> Composable HTML components in Golang

<p align="center">
	<img src="https://github.com/stevelacy/daz/raw/master/daz.go.png" width="300">
</p>

[![GoDoc](https://godoc.org/github.com/stevelacy/daz?status.svg)](https://godoc.org/github.com/stevelacy/daz)![Go](https://github.com/stevelacy/daz/workflows/Go/badge.svg)


![daz carbon example](./carbon.png)

Daz is a "functional" alternative to using templates, and allows for nested components/lists
Also enables template-free server-side rendered components with support for nested lists. It is inspired by [HyperScript](https://github.com/hyperhype/hyperscript).


A component can be created and used with simple functions:
```golang
// Example prop for a component
type User struct {
	Name string
	// ...
}

func MyComponenet(user User) HTML {
	return H(
		"div",
		Attr{"class": "bg-grey-50"},
		user.Name,
	)
}

func Root() HTML {
	user := User{Name: "Daz"}
	return H("html", MyComponenet(user))
}

// And used in a handler:

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(Root()()))
}
```

Lists can be easily created without needing to embed a `range / end` in a template:
```golang
items := []HTML{
	H("li", "item one"),
	H("li", "item two"),
}

element := H("ul", Attr{"class": "bg-grey-50"}, items)

div := H("div", element)
```


### Install

```
import (
	"github.com/stevelacy/daz"
)

```

### Usage

#### func `H`

Create a HTML element:
```golang
H("div", ...attrs)

```

#### struct `Attr`

HTML attributes:
```golang
Attr{
	"class": "app",
	"onClick": "javascriptFunc()",
}
```

#### func `UnsafeContent`

This will bypass HTML sanitization and allow for direct injecting
```golang

injection := "<script>alert('xss')</script>"
root := H("div", UnsafeContent(injection))
// <script>alert('xss')</script>
```
