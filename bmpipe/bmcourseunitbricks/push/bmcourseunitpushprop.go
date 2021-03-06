package courseunitpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/courseunit"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BmCourseUnitPushProp struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmCourseUnitPushProp) Exec() error {
	tmp := b.bk.Pr.(courseunit.BmCourseUnit)
	bindSessionable(tmp)
	bindTeacher(tmp)
	bindRoom(tmp)
	b.bk.Pr = tmp
	return nil
}

func (b *BmCourseUnitPushProp) Prepare(pr interface{}) error {
	req := pr.(courseunit.BmCourseUnit)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmCourseUnitPushProp) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmCourseUnitPushProp) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmCourseUnitPushProp) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(courseunit.BmCourseUnit)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmCourseUnitPushProp) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval courseunit.BmCourseUnit = b.BrickInstance().Pr.(courseunit.BmCourseUnit)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func bindSessionable(courseUnit courseunit.BmCourseUnit) error {
	sessionabel := courseUnit.Sessionable
	one2one := courseunit.BmCourseUnitBindSessionable{}
	one2one.CourseUnitId = courseUnit.Id
	one2one.SessionableId = sessionabel.Id
	one2one.Id_ = bson.NewObjectId()
	one2one.ResetIdWithId_()
	one2one.CheckExist()
	err := one2one.InsertBMObject()
	if err != nil {
		panic("courseUnit bind error")
	}
	return err
}

func bindTeacher(courseUnit courseunit.BmCourseUnit) error {
	t := courseUnit.Teacher
	one2one := courseunit.BmCourseUnitBindTeacher{}
	one2one.CourseUnitId = courseUnit.Id
	one2one.TeacherId = t.Id
	one2one.Id_ = bson.NewObjectId()
	one2one.ResetIdWithId_()
	one2one.CheckExist()
	err := one2one.InsertBMObject()
	if err != nil {
		panic("courseUnit bind error")
	}
	return err
}

func bindRoom(courseUnit courseunit.BmCourseUnit) error {
	r := courseUnit.Room
	one2one := courseunit.BmCourseUnitBindRoom{}
	one2one.CourseUnitId = courseUnit.Id
	one2one.RoomId = r.Id
	one2one.Id_ = bson.NewObjectId()
	one2one.ResetIdWithId_()
	one2one.CheckExist()
	err := one2one.InsertBMObject()
	if err != nil {
		panic("courseUnit bind error")
	}
	return err
}

