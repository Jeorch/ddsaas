package student

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BMStudents struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Students []BMStudent `json:"students" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BMStudents) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BMStudents) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BMStudents) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BMStudents) QueryId() string {
	return bd.Id
}

func (bd *BMStudents) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BMStudents) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BMStudents) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BMStudents) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BMStudents) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BMStudents) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BMStudents) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BMStudents) FindMulti(req request.Request) error {
	err := bmmodel.FindMutil(req, &bd.Students)
	for i, r := range bd.Students {
		r.ResetIdWithId_()
		bd.Students[i] = r
	}
	return err
}