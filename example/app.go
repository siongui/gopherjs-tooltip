package main

import (
	"github.com/gopherjs/gopherjs/js"
	tt "github.com/siongui/gopherjs-tooltip"
	"strings"
)

func main() {
	spans := js.Global.Get("document").Call("getElementById", "container").Call("querySelectorAll", "span")
	// access individual span
	length := spans.Get("length").Int()
	for i := 0; i < length; i++ {
		span := spans.Call("item", i)
		word := strings.ToLower(span.Get("innerHTML").String())
		tooltipContent := word + " " + word + "<br>" + "<span>" + word + "</span>" + " " + word
		tt.AddTooltipToElement(span, tooltipContent)
	}
}
