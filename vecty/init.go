// +build js,wasm

package vecty

import (
	"fmt"
	"github.com/pubgo/godom"
	"github.com/pubgo/godom/version"
)

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

func init() {
	fmt.Printf("vecty Version %s, BuildV %s, CommitV %s\n", version.Version, version.BuildV, version.CommitV)
}
