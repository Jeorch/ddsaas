package authfind

import (
	"fmt"
	//"github.com/alfredyang1986/ddsaas/bmcommon/bmsingleton/bmconf"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/ddsaas/bmmodel/auth"
	//"github.com/alfredyang1986/ddsaas/bmpipe/bmauthbricks/push"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"reflect"
)

type BMAuthPhoneFindBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAuthPhoneFindBrick) Exec() error {
	var tmp auth.BmPhone
	err := tmp.FindOne(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *BMAuthPhoneFindBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	fmt.Println(req)
	b.BrickInstance().Req = &req
	//b.bk.Req = &req
	return nil
}

func (b *BMAuthPhoneFindBrick) Done(pkg string, idx int64, e error) error {
	if e != nil && e.Error() == "not found" {
		reval := auth.BmAuth{}
		reval.Phone = auth.BmPhone{}
		reval.Phone.PhoneNo = b.BrickInstance().Req.CondiQueryVal("phone_no", "BmPhone").(string)
		b.BrickInstance().Pr = reval

		bmrouter.NextBrickRemote("insertauth", 0, b)
	} else {
		bmrouter.NextBrickRemote("phone2auth", 0, b)
	}

	return nil
}

func (b *BMAuthPhoneFindBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAuthPhoneFindBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	if reflect.ValueOf(pr).Type().Name() == "BmPhone" {
		tmp := pr.(auth.BmPhone)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	} else {
		tmp := pr.(auth.BmAuth)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	}
}

func (b *BMAuthPhoneFindBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.BmAuth = b.BrickInstance().Pr.(auth.BmAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
