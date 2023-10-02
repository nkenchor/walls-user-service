package event

import (
	"walls-user-service/internal/core/helper/event-helper/eto"
)

type UserUpdatedEvent struct {
	eto.Event
}
