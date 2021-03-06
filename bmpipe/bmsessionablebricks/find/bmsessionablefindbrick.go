package sessionablefind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/ddsaas/bmmodel/courseunit"
	"github.com/alfredyang1986/ddsaas/bmmodel/sessionable"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"io"
)

type BmSessionableFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmSessionableFindBrick) Exec() error {
	var tmp sessionable.BmSessionable
	err := tmp.FindOne(*b.bk.Req)
	tmp.ReSetProp()
	ReSetClassDate(&tmp)
	b.bk.Pr = tmp
	return err
}

func (b *BmSessionableFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Pr = req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BmSessionableFindBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmSessionableFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmSessionableFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(sessionable.BmSessionable)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmSessionableFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval sessionable.BmSessionable = b.BrickInstance().Pr.(sessionable.BmSessionable)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func ReSetClassDate(bd *sessionable.BmSessionable) error {

	if bd.Status != 2 {
		return nil
	}

	req := request.Request{}
	req.Res = "BmCourseUnit"
	var eqCondi []interface{}
	eq := request.Eqcond{}
	eq.Ky = "sessionableId"
	eq.Vy = bd.Id
	eqCondi = append(eqCondi, eq)
	c := req.SetConnect("conditions", eqCondi)
	req1 := c.(request.Request)
	var tmp courseunit.BmCourseUnits
	err := tmp.FindMulti(req1)
	if tmp.CourseUnits == nil {
		return nil
	}
	tmp.SortByStartDate(true)
	bd.StartDate = tmp.CourseUnits[0].StartDate
	tmp.SortByEndDate(false)
	bd.EndDate = tmp.CourseUnits[0].EndDate

	return err
}

