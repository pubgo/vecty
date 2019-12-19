package vecty

import "github.com/pubgo/godom"

var (
	global    = godom.Global()
	undefined = godom.Undefined()
	window    = godom.GetWindow()
	document  = window.Document()
)

type jsObject = godom.Value
type jsFunc = godom.Func

var valueOf = godom.ValueOf
var funcOf = godom.FuncOf
var wrapEvent = godom.WrapEvent

type Event = godom.Event
type BasicEvent = godom.BasicEvent
