package models

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo/bson"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	dao = Application{}
)

func TestCreateApplication(t *testing.T) {

	Convey("When user create an new Application", t, func() {

		Convey("the data could be saved to the DB", func() {
			newApp := Application{
				ID:          bson.NewObjectId(),
				Name:        "test App",
				Type:        "mini App",
				Description: "a mini app for testing",
			}

			insertError := dao.InsertApplication(newApp)
			So(insertError, ShouldBeNil)
		})
	})

	Convey("When user want to search the Application", t, func() {

		Convey("user can find all Application", func() {
			apps, error := dao.FindAllApplication()
			So(error, ShouldBeNil)
			So(apps, ShouldNotBeEmpty)
		})

		Convey("user can find an Application by ID", func() {
			newApp := Application{
				ID:          bson.NewObjectId(),
				Name:        "test App for find one",
				Type:        "mini App",
				Description: "test App for find one",
			}

			error := dao.InsertApplication(newApp)
			So(error, ShouldBeNil)

			app, error := dao.FindAppByID(fmt.Sprintf(`%x`, string(newApp.ID)))
			So(error, ShouldBeNil)
			So(app, ShouldNotBeNil)
		})
	})
}
