package event

import "event/entities"

type EventRepository interface {
	CreateEvent(newAdd entities.Event) (entities.Event, error)
	GetAllEvent() ([]entities.Event, error)
	GetEventID(id uint) (entities.Event, error)
	UpdateEvent(id uint, UpdateEvent entities.Event, UserID uint) (entities.Event, error)
	DeleteEvent(id uint, UserID uint) error
}
