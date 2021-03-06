package attendeefind

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/alfredyang1986/ddsaas/bmmodel/attendee"
	"io"
	"net/http"
)

type BMAttendeeFindMulti struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *BMAttendeeFindMulti) Exec() error {
	var tmp attendee.BmAttendees
	err := tmp.FindMulti(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *BMAttendeeFindMulti) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	b.BrickInstance().Req = &req
	return nil
}

func (b *BMAttendeeFindMulti) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *BMAttendeeFindMulti) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *BMAttendeeFindMulti) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(attendee.BmAttendees)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *BMAttendeeFindMulti) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		results := b.BrickInstance().Pr.(attendee.BmAttendees)
		jsonapi.ToJsonAPI(results.Attendees, w)
	}
}

