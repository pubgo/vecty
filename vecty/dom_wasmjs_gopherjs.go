// +build js,wasm

package vecty

// Node returns the underlying JavaScript Element or TextNode.
//
// It panics if it is called before the DOM node has been attached, i.e. before
// the associated component's Mounter interface would be invoked.
func (h *HTML) Node() jsObject {
	if h.node == nil {
		panic("vecty: cannot call (*HTML).Node() before DOM node creation / component mount")
	}
	return h.node
}

// RenderIntoNode renders the given component into the existing HTML element by
// replacing it.
//
// If the Component's Render method does not return an element of the same type,
// an error of type ElementMismatchError is returned.
func RenderIntoNode(node jsObject, c Component) error {
	return renderIntoNode("RenderIntoNode", node, c)
}

func toLower(s string) string {
	// We must call the prototype method here to workaround a limitation of
	// syscall/js in both Go and GopherJS where we cannot call the
	// `toLowerCase` string method. See https://github.com/golang/go/issues/35917
	return global.Get("String").Get("prototype").Get("toLowerCase").Call("call", valueOf(s)).String()
}
