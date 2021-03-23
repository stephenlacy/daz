package daz

import (
	"fmt"
	"html"
	"strings"
)

var selfClosingTags = map[string]int{
	"area":  1,
	"br":    1,
	"hr":    1,
	"image": 1,
	"input": 1,
	"img":   1,
	"link":  1,
	"meta":  1,
}

// Attr is a HTML element attribute
// <a href="#"> => Attr{"href": "#"}
type Attr map[string]string

// H is the base HTML func
func H(el string, attrs ...interface{}) func() string {
	contents := []string{}
	attributes := ""
	for _, v := range attrs {
		switch v := v.(type) {
		case string:
			contents = append(contents, escape(v))
		case Attr:
			attributes = attributes + getAttributes(v)
		case []string:
			children := strings.Join(v, "")
			contents = append(contents, escape(children))
		case []func() string:
			children := subItems(v)
			contents = append(contents, children)
		case func() string:
			contents = append(contents, v())
		default:
			contents = append(contents, escape(fmt.Sprintf("%v", v)))
		}
	}
	return func() string {
		elc := escape(el)
		if _, ok := selfClosingTags[elc]; ok {
			return "<" + elc + attributes + " />"
		}
		return "<" + elc + attributes + ">" + strings.Join(contents, "") + "</" + elc + ">"
	}
}

func escape(str string) string {
	return html.EscapeString(str)
}

func subItems(attrs []func() string) string {
	res := []string{}
	for _, v := range attrs {
		res = append(res, v())
	}
	return strings.Join(res, "")
}

func getAttributes(attributes Attr) string {
	res := []string{}
	for k, v := range attributes {
		pair := fmt.Sprintf("%v='%v'", escape(k), escape(v))
		res = append(res, pair)
	}
	prefix := ""
	if len(res) > 0 {
		prefix = " "
	}
	return prefix + strings.Join(res, " ")
}
