package tooltip

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
)

type Tooltip struct {
	self             *js.Object
	isMouseInTooltip bool
	left             int
	top              int
}

func (tt *Tooltip) onMouseEnter() {
	tt.isMouseInTooltip = true
}

func (tt *Tooltip) onMouseLeave() {
	tt.isMouseInTooltip = false
	tt.hide()
}

func (tt *Tooltip) registerMouseEnterLeaveHandler() {
	tt.self.Set("onmouseenter", func() {
		tt.onMouseEnter()
	})
	tt.self.Set("onmouseleave", func() {
		tt.onMouseLeave()
	})
}

func (tt *Tooltip) hide() {
	if !tt.isMouseInTooltip {
		tt.self.Get("style").Set("left", "-9999px")
	}
}

func (tt *Tooltip) removeAllChildren() {
	for tt.self.Call("hasChildNodes").Bool() {
		tt.self.Call("removeChild", tt.self.Get("lastChild"))
	}
}

func (tt *Tooltip) setInnerHTML(html string) {
	tt.removeAllChildren()
	tt.self.Set("innerHTML", html)
}

func (tt *Tooltip) setPosition(left, top int) {
	tt.left = left
	tt.top = top
}

func (tt *Tooltip) show() {
	// move tooltip to the right
	// (don't cross the right side of browser inner window)
	offsetWidth := tt.self.Get("style").Get("offsetWidth").Int()
	right := tt.left + offsetWidth
	if right > viewWidthInt() {
		tt.left = right - viewWidthInt()
	}

	tt.self.Get("style").Set("left", strconv.Itoa(tt.left)+"px")
	tt.self.Get("style").Set("top", strconv.Itoa(tt.top)+"px")
}

func (tt *Tooltip) appendToBodyElement() {
	// insert tooltip at the end of body element
	js.Global.Get("document").Call("getElementsByTagName", "body").Call("item", 0).Call("appendChild", tt.self)
}

func (tt *Tooltip) createTooltipInstance() {
	tt.self = js.Global.Get("document").Call("createElement", "div")
	// set css class as tooltip
	tt.self.Get("classList").Call("add", "tooltip")
	// set max-width
	tt.self.Get("style").Set("max-width", viewWidth()+"px")
}

func (tt *Tooltip) appendCSSToHeadElement() {
	css := `.tooltip {
		position: absolute;
		left: -9999px;
		background-color: #CCFFFF;
		border-radius: 10px;
		font-family: Tahoma, Arial, serif;
		word-wrap: break-word;
	}`
	s := js.Global.Get("document").Call("createElement", "style")
	s.Set("innerHTML", css)
	// insert style of tooltip at the end of head element
	js.Global.Get("document").Call("getElementsByTagName", "head").Call("item", 0).Call("appendChild", s)
}

func NewTooltip() *Tooltip {
	tt := &Tooltip{
		isMouseInTooltip: false,
	}
	tt.appendCSSToHeadElement()
	tt.createTooltipInstance()
	tt.appendToBodyElement()
	tt.registerMouseEnterLeaveHandler()

	return tt
}

func viewWidth() string {
	return js.Global.Get("innerWidth").String()
}

func viewWidthInt() int {
	return js.Global.Get("innerWidth").Int()
}
