package event

import (
	"event/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type EventDB struct {
	Db *gorm.DB
}

// Get Access DB
func NewDB(db *gorm.DB) *EventDB {
	return &EventDB{
		Db: db,
	}
}

func (e *EventDB) CreateEvent(newAdd entities.Event) (entities.Event, error) {
	if err := e.Db.Create(&newAdd).Error; err != nil {
		log.Warn(err)
		return newAdd, err
	}
	return newAdd, nil
}
func (e *EventDB) GetAllEvent() ([]entities.Event, error) {
	var event []entities.Event
	if err := e.Db.Find(&event).Error; err != nil {
		log.Warn("Error Get Data", err)
		return event, err
	}
	return event, nil
}

func (e *EventDB) GetEventID(id uint) (entities.Event, error) {
	var event entities.Event
	if err := e.Db.Where("id= ?", id).Find(&event).Error; err != nil {
		log.Warn("Error Get By ID", err)
		return event, err
	}
	return event, nil
}

func (e *EventDB) UpdateEvent(id uint, UpdateEvent entities.Event, UserID uint) (entities.Event, error) {
	var event entities.Event

	if err := e.Db.Where("id =? AND user_id =?", id, UserID).First(&event).Updates(&UpdateEvent).Find(&event).Error; err != nil {
		log.Warn("Update Error", err)
		return event, err
	}

	return event, nil
}

func (e *EventDB) DeleteEvent(id uint, UserID uint) error {

	var delete entities.Event
	if err := e.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}
