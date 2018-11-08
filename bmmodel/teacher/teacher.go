package teacher

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmTeacher struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Intro string `json:"intro" bson:"intro"`

	BrandId string `json:"brandId" bson:"brandId"`

	Name     string `json:"name" bson:"name"`
	Nickname string `json:"nickname" bson:"nickname"`
	Icon     string `json:"icon" bson:"icon"`
	Dob      int64 `json:"dob" bson:"dob"`
	Gender   int64 `json:"gender" bson:"gender"`
	RegDate  int64 `json:"reg_date" bson:"reg_date"`
	Contact  string `json:"contact" bson:"contact"`
	WeChat  string `json:"wechat" bson:"wechat"`

	//Person person.BmPerson `json:"person" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmTeacher) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmTeacher) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmTeacher) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmTeacher) QueryId() string {
	return bd.Id
}

func (bd *BmTeacher) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmTeacher) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmTeacher) SetConnect(tag string, v interface{}) interface{} {
	//switch tag {
	//case "person":
	//	bd.Person = v.(person.BmPerson)
	//}
	return bd
}

func (bd BmTeacher) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmTeacher) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmTeacher) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmTeacher) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
