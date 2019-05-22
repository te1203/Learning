package models

import "github.com/globalsign/mgo/bson"

// Application Model
type Application struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Type        string        `bson:"app_type" json:"app_type"`
	Description string        `bson:"description" json:"description"`
}

const (
	db         = "Application"
	collection = "WeChat_Application"
)

// InsertApplication to the DB
func (m *Application) InsertApplication(app Application) error {
	return Insert(db, collection, app)
}

// FindAllApplication from the DB
func (m *Application) FindAllApplication() ([]Application, error) {
	var result []Application
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

// FindAppByID from the DB
func (m *Application) FindAppByID(id string) (Application, error) {
	var result Application
	err := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	return result, err
}

// UpdateApp is used for update
func (m *Application) UpdateApp(app Application) error {
	return Update(db, collection, bson.M{"_id": app.ID}, app)
}

// RemoveApp is used for remove
func (m *Application) RemoveApp(id string) error {
	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
