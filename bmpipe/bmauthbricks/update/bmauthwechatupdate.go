package authupdate

import (
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	//"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type BMAuthWechatUpdateBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAuthWechatUpdateBrick) Exec() error {
	tmp := auth.BmPhone{}
	tmp.UpdateBMObject(*b.bk.Req)
	b.bk.Pr = tmp
	return nil
}

func (b *BMAuthWechatUpdateBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Req = &req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BMAuthWechatUpdateBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAuthWechatUpdateBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAuthWechatUpdateBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.BmPhone)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAuthWechatUpdateBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BmPhone = b.BrickInstance().Pr.(auth.BmPhone)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
