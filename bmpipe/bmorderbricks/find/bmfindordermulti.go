package orderfind

import (
	"github.com/alfredyang1986/blackmirror-modules/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmmodel/order"
	"github.com/alfredyang1986/blackmirror-modules/bmmodel/request"
	"github.com/alfredyang1986/blackmirror-modules/bmerror"
	"github.com/alfredyang1986/blackmirror-modules/bmpipe"
	"github.com/alfredyang1986/blackmirror-modules/bmrouter"
	"github.com/alfredyang1986/blackmirror-modules/jsonapi"
	"net/http"
	"io"
)

type BMOrderFindMultiBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMOrderFindMultiBrick) Exec() error {
	var tmp order.Orders
	err := tmp.FindMulti(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *BMOrderFindMultiBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Pr = req
	b.BrickInstance().Req = &req
	return nil
}

func (b *BMOrderFindMultiBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMOrderFindMultiBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMOrderFindMultiBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(order.Orders)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMOrderFindMultiBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval order.Orders = b.BrickInstance().Pr.(order.Orders)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
