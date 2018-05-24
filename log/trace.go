package log

import (
	"fmt"
	"sync"
	"time"
)

var traceTrees map[string]*traceTree
var traceMutex sync.RWMutex

func init() {
	traceTrees = make(map[string]*traceTree)
}

func traceBegin(contextID string, message string, keysAndValues ...interface{}) {
	traceMutex.RLock()
	t, ok := traceTrees[contextID]
	traceMutex.RUnlock()
	if !ok {
		traceMutex.Lock()
		traceTrees[contextID] = &traceTree{
			make([]*trace, 0),
			nil,
		}
		traceMutex.Unlock()

		traceMutex.RLock()
		t = traceTrees[contextID]
		traceMutex.RUnlock()
	}
	t.begin(message, keysAndValues...)
}

func traceEnd(contextID string, state interface{}) {
	traceMutex.RLock()
	t, ok := traceTrees[contextID]
	traceMutex.RUnlock()
	if !ok {
		return
	}
	t.end(state)
}

func traceEndRoot(contextID string, state interface{}) {
	traceMutex.RLock()
	t, ok := traceTrees[contextID]
	traceMutex.RUnlock()
	if !ok {
		return
	}
	t.endRoot(state)
}

func traceSum(contextID string) (string, []interface{}) {
	defer func() {
		traceMutex.Lock()
		delete(traceTrees, contextID)
		traceMutex.Unlock()
	}()
	traceMutex.RLock()
	t, ok := traceTrees[contextID]
	traceMutex.RUnlock()
	if !ok {
		return "", nil
	}
	return t.sum()
}

type traceTree struct {
	stack []*trace
	root  *trace
}

func (t *traceTree) begin(message string, keysAndValues ...interface{}) *trace {
	stk := t.stack
	trc := &trace{}
	trc.begin(message, keysAndValues...)
	l := len(stk)
	if l > 0 {
		parent := stk[l-1]
		parent.addChild(trc)
	} else {
		t.root = trc
	}
	t.stack = append(stk, trc)
	return trc
}

func (t *traceTree) end(state interface{}) {
	stk := t.stack
	current := len(stk) - 1
	if current == -1 {
		return
	}
	trc := stk[current]
	trc.end(state)
	t.stack = stk[:current]
}

func (t *traceTree) endRoot(state interface{}) {
	t.root.end(state)
	t.stack = t.stack[:0]
}

func (t *traceTree) sum() (string, []interface{}) {
	return t.root.sum()
}

type trace struct {
	message   string
	alias     string
	state     interface{}
	startTime time.Time
	endTime   time.Time
	children  []*trace
}

func (t *trace) begin(message string, keysAndValues ...interface{}) {
	t.message = message
	if len(keysAndValues) <= 2 {
		t.alias = fmt.Sprintf("%v-%v", keysAndValues...)
	} else {
		t.alias = fmt.Sprintf("%v-%v", keysAndValues[0], keysAndValues[1])
	}
	t.startTime = time.Now()
}

func (t *trace) end(state interface{}) {
	if state == nil {
		t.state = "done"
	} else {
		t.state = state
	}
	t.endTime = time.Now()
}

func (t *trace) addChild(child *trace) {
	if t.children == nil {
		t.children = make([]*trace, 0)
	}
	t.children = append(t.children, child)
}

func (t *trace) sum() (string, []interface{}) {
	var message string
	var namesAndDurations = make([]interface{}, 0)
	if t.state == nil {
		t.end("not a ended trace")
		message = t.message + "(not a ended trace)"
		namesAndDurations = []interface{}{
			t.alias,
			"not a ended trace",
		}
	} else {
		message = t.message + fmt.Sprintf("(%v)", t.state)
		namesAndDurations = []interface{}{
			t.alias,
			t.endTime.Sub(t.startTime),
		}
	}
	for i, child := range t.children {
		m, nad := child.sum()
		if i == 0 {
			message = message + ">>" + m
		} else {
			message = message + ">" + m
		}
		namesAndDurations = append(namesAndDurations, nad...)
	}
	return message, namesAndDurations
}
