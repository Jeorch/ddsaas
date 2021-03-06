package room

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmRoom struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	YardId   string  `json:"yardId" bson:"yardId"`
	Title    string  `json:"title" bson:"title"`
	RoomType float64  `json:"roomType" bson:"roomType"`
	Capacity float64 `json:"capacity" bson:"capacity"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmRoom) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmRoom) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmRoom) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmRoom) QueryId() string {
	return bd.Id
}

func (bd *BmRoom) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmRoom) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmRoom) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd BmRoom) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmRoom) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmRoom) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmRoom) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *BmRoom) DeleteOne(req request.Request) error {
	return bmmodel.DeleteOne(req, bd)
}
