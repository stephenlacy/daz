package daz

import (
	"testing"
)

var fixture1 = "<div class='app view'><header>test 1<nav>Welcome</nav></header><div>&lt;escaped&gt;</div></div>"
var fixture2 = "<div>onetwothree</div>"
var fixture3 = "<div><div>one</div>one<>text</></div>"
var fixture4 = "<div class='bg-grey-50' data-id='div-1'>content</div>"
var fixture5 = "<div>O&#39;Brian<input type='text' value='input value&#39;s' /></div>"
var fixture6 = "<div><img src='https://example.com/image.png' /><br /></div>"
var fixture7 = "<div>&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;</div>"
var fixture8 = "<div><script>alert('xss')</script></div>"

func TestBasicRender(t *testing.T) {
	attrs := Attr{"class": "app view"}
	nav := H("nav", "Welcome")
	header := H("header", "test 1", nav)
	escaped := H("div", "<escaped>")
	root := H("div", attrs, header, escaped)
	res := root()
	if res != fixture1 {
		t.Errorf("got: %v wanted: %v", res, fixture1)
	}
}

func TestStringItems(t *testing.T) {
	items := []string{"one", "two", "three"}
	root := H("div", items)
	res := root()
	if res != fixture2 {
		t.Errorf("got: %v wanted: %v", res, fixture2)
	}
}

func TestItems1(t *testing.T) {
	one := H("div", "one")
	two := func() string { return "one" }
	three := H("", "text")
	items := []HTML{one, two, three}

	root := H("div", items)
	res := root()
	if res != fixture3 {
		t.Errorf("got: %v wanted: %v", res, fixture3)
	}
}
func TestItems2(t *testing.T) {
	one := H("div", "one")
	two := func() string { return "one" }
	three := H("", "text")
	items := []HTML{one, two, three}

	root := H("div", items)
	res := root()
	if res != fixture3 {
		t.Errorf("got: %v wanted: %v", res, fixture3)
	}
}

func TestAttrs(t *testing.T) {
	attr1 := Attr{"class": "bg-grey-50"}
	attr2 := Attr{"data-id": "div-1"}
	root := H("div", attr1, "content", attr2)
	res := root()
	if res != fixture4 {
		t.Errorf("got: %v wanted: %v", res, fixture4)
	}
}
func TestQuoted(t *testing.T) {
	value := "input value's"
	input := H("input", Attr{"type": "text", "value": value})
	root := H("div", "O'Brian", input)
	res := root()
	if res != fixture5 {
		t.Errorf("got: %v wanted: %v", res, fixture5)
	}
}

func TestSelfClosing(t *testing.T) {
	root := H("div", H("img", Attr{"src": "https://example.com/image.png"}), H("br"))
	res := root()
	if res != fixture6 {
		t.Errorf("got: %v wanted: %v", res, fixture6)
	}
}

func TestXSS1(t *testing.T) {
	root := H("div", "<script>alert('xss')</script>")
	res := root()
	if res != fixture7 {
		t.Errorf("got: %v wanted: %v", res, fixture7)
	}
}

func TestUnsafeContent(t *testing.T) {
	injection := "<script>alert('xss')</script>"
	root := H("div", UnsafeContent(injection))
	res := root()
	if res != fixture8 {
		t.Errorf("got: %v wanted: %v", res, fixture8)
	}
}

func BenchmarkBasicRender(b *testing.B) {
	attrs := Attr{"class": "app view"}
	nav := H("nav", "Welcome")
	header := H("header", "test 1", nav)
	root := H("div", attrs, header)
	root()
}
