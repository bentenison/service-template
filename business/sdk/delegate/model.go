package delegate

import (
	"context"
	"fmt"
)

type Func func(context.Context, Data) error

type Data struct {
	Domain    string
	Action    string
	RawParams []byte
}

func (d Data) String() string {
	return fmt.Sprintf(
		"Event{Domain:%#v, Action:%#v, RawParams:%#v}",
		d.Domain, d.Action, string(d.RawParams),
	)
}
