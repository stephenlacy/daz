### Carbon.now.sh

https://carbon.now.sh/68eRaNf1Tvi1t0q8I0Dq

```go

func Handler(w http.ResponseWriter, r *http.Request) {
	links := H("link", Attr{
		"href": "https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css",
		"rel":  "stylesheet",
	})
	meta := []HTML{
		H("meta", Attr{"charset": "utf-8"}),
		H("meta", Attr{
			"name":    "viewport",
			"content": "width=device-width, initial-scale=1.0",
		}),
	}
	head := H("head", H("title", "Example Server"), meta, links)
	style := Attr{"style": "background: #efefef;"}
	body := H("body", style)
	html := H("html", head, body)
	w.Write([]byte(html()))
}

```
