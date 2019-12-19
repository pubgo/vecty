package components

import (
	"github.com/pubgo/vecty/example/todomvc/actions"
	"github.com/pubgo/vecty/example/todomvc/dispatcher"
	"github.com/pubgo/vecty/example/todomvc/store"
	"github.com/pubgo/vecty/example/todomvc/store/model"
	"github.com/pubgo/vecty/vecty"
	"github.com/pubgo/vecty/vecty/elem"
	"github.com/pubgo/vecty/vecty/event"
)

// FilterButton is a vecty.Component which allows the user to select a filter
// state.
type FilterButton struct {
	vecty.Core

	Label  string            `vecty:"prop"`
	Filter model.FilterState `vecty:"prop"`
	IsShow bool
}

func (b *FilterButton) Mount() {
	vecty.Rerender(b)
}

func (b *FilterButton) SkipRender(prev vecty.Component) bool {
	return b.IsShow
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

// Render implements the vecty.Component interface.
func (b *FilterButton) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.Markup(
				vecty.MarkupIf(store.Filter == b.Filter, vecty.Class("selected")),
				prop.Href("#"),
				event.Click(b.onClick).PreventDefault(),
			),

			vecty.Text(b.Label),
		),
	)
}
