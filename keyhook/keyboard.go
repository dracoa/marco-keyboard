package keyhook

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	hook "github.com/robotn/gohook"
)

var observable rxgo.Observable

func init() {
	ch := make(chan rxgo.Item)
	go func() {
		EvChan := hook.Start()
		defer hook.End()
		for ev := range EvChan {
			ch <- rxgo.Item{
				V: ev,
				E: nil,
			}
		}
	}()
	observable = rxgo.FromChannel(ch)
}

func KeyboardHook() rxgo.Observable {
	return observable.Filter(func(item interface{}) bool {
		e := item.(hook.Event)
		return e.Mask == 8 && e.Kind == 5
	}).Map(func(context context.Context, item interface{}) (interface{}, error) {
		e := item.(hook.Event)
		return e.Rawcode, nil
	})
}
