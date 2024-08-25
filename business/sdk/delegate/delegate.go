package delegate

import (
	"context"

	"github.com/bentenison/microservice/foundation/logger"
)

type (
	domain string
	action string
)
type Delegate struct {
	log   *logger.CustomLogger
	funcs map[domain]map[action][]Func
}

func New(log *logger.CustomLogger) *Delegate {
	return &Delegate{
		log:   log,
		funcs: make(map[domain]map[action][]Func),
	}
}

// Register delegate function to use as a event
func (d *Delegate) Register(domainType string, actionType string, fn Func) {
	aMap, ok := d.funcs[domain(domainType)]
	if !ok {
		aMap := make(map[action][]Func)
		d.funcs[domain(domainType)] = aMap
	}
	funcs := aMap[action(actionType)]
	funcs = append(funcs, fn)
	aMap[action(actionType)] = funcs
}

// Call executes all functions registered for the specified domain and
// action. These functions are executed synchronously on the G making the call.
func (d *Delegate) Call(ctx context.Context, data Data) error {
	d.log.Info("delegate call", map[string]interface{}{
		"domain": data.Domain,
		"action": data.Action,
		"params": data.RawParams,
		"status": "started",
	})
	defer d.log.Info("delegate call", map[string]interface{}{
		"status": "completed",
	})

	if dMap, ok := d.funcs[domain(data.Domain)]; ok {
		if funcs, ok := dMap[action(data.Action)]; ok {
			for _, fn := range funcs {
				d.log.Info("delegate call", map[string]interface{}{
					"status": "sending",
				})

				if err := fn(ctx, data); err != nil {
					d.log.Error("delegate call", map[string]interface{}{
						"error": err,
					})
				}
			}
		}
	}

	return nil
}
