package link_manager_events

import (
	om "github.com/the-gigi/delinkcious/pkg/object_model"
)

type Event struct {
	EventType om.EventTypeEnum
	Username  string
	Link      *om.Link
}
