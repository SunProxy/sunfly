package entity

import "github.com/sunproxy/sunfly/dragonfly/event"

type Handler interface {
	HandleItemMerge(ctx *event.Context, oldItem, NewItem *Item)
}

type NopHandler struct{}

// Compile time check to make sure NopHandler implements Handler.
var _ Handler = (*NopHandler)(nil)

func (n NopHandler) HandleItemMerge(*event.Context, *Item, *Item) {}
