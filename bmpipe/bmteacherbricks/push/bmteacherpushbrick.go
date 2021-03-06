package teacherpush

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/teacher"
	"io"
	"net/http"
	"time"
)

type BmTeacherPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BmTeacherPushBrick) Exec() error {
	var tmp teacher.BmTeacher = b.bk.Pr.(teacher.BmTeacher)
	tmp.CreateTime = float64(time.Now().UnixNano() / 1e6)
	tmp.InsertBMObject()
	b.bk.Pr = tmp
	return nil
}

func (b *BmTeacherPushBrick) Prepare(pr interface{}) error {
	req := pr.(teacher.BmTeacher)
	b.BrickInstance().Pr = req
	return nil
}

func (b *BmTeacherPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BmTeacherPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BmTeacherPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(teacher.BmTeacher)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BmTeacherPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval teacher.BmTeacher = b.BrickInstance().Pr.(teacher.BmTeacher)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
