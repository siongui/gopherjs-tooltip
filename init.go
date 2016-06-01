package tooltip

import (
	"github.com/gopherjs/gopherjs/js"
	"time"
)

// when user's mouse hovers over words,
// delay a period of time before look up.
var DELAY_INTERVAL = 1000 * time.Millisecond
var isMouseInWord = false
var tooltipPtr *Tooltip

var onWordMouseOver = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	isMouseInWord = true
	this.Get("style").Set("color", "red")

	time.AfterFunc(DELAY_INTERVAL, func() {
		if this.Get("style").Get("color").String() == "red" {
			// mouse is still on word
			tooltipPtr.setInnerHTML(this.Get("dataset").Get("tooltipContent").String())
			tooltipPtr.setPosition(
				this.Call("getBoundingClientRect").Get("left").Int(),
				this.Call("getBoundingClientRect").Get("top").Int()+this.Get("offsetHeight").Int())
			tooltipPtr.show()
		}
	})
	return nil
})

var onWordMouseOut = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	isMouseInWord = false
	this.Get("style").Set("color", "")

	time.AfterFunc(DELAY_INTERVAL, func() {
		if !isMouseInWord {
			tooltipPtr.hide()
		}
	})
	return nil
})

func AddTooltipToElement(elm *js.Object, tooltipContent string) {
	elm.Get("dataset").Set("tooltipContent", tooltipContent)
	elm.Set("onmouseover", onWordMouseOver)
	elm.Set("onmouseout", onWordMouseOut)
}

func init() {
	tooltipPtr = NewTooltip()
}
