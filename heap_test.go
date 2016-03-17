package dpll

import (
	"reflect"
	"testing"
)

func TestActivityQueue(t *testing.T) {
	// initialize variables and activity values
	const (
		v1 = Var(1)
		v2 = Var(2)
		v3 = Var(3)
		v4 = Var(4)
		v5 = Var(5)
		v6 = Var(6)
	)
	activity := []float64{
		v1: 10,
		v2: 9,
		v3: 8,
		v4: 7,
		v5: 6,
		v6: 5,
	}

	// sanity check
	queue := newActivityQueue(&activity)
	queue.Push(v1)
	queue.Push(v2)
	vs := dequeueAll(queue)
	vsexpect := []Var{v1, v2}
	if !reflect.DeepEqual(vs, vsexpect) {
		t.Errorf("queue order: %v (!= %v)", vs, vsexpect)
	}

	// push everything and decrease some values, one at a time
	queue.Push(v1)
	queue.Push(v2)
	queue.Push(v3)
	queue.Push(v4)
	queue.Push(v5)
	queue.Push(v6)
	activity[v2] = 1e-10
	queue.Decrease(v2)
	activity[v1] = 7.7
	queue.Decrease(v1)
	activity[v5] = 1
	queue.Decrease(v5)
	vs = dequeueAll(queue)
	vsexpect = []Var{v3, v1, v4, v6, v5, v2}
	if !reflect.DeepEqual(vs, vsexpect) {
		t.Errorf("queue order: %v (!= %v)", vs, vsexpect)
	}

	// rebuilding from scratch with the same activity as before has the same
	// result.
	queue.Rebuild([]Var{v1, v2, v3, v4, v5, v6})
	vs = dequeueAll(queue)
	if !reflect.DeepEqual(vs, vsexpect) {
		t.Errorf("queue order: %v (!= %v)", vs, vsexpect)
	}
}

func dequeueAll(q *activityQueue) []Var {
	var vs []Var
	for {
		x, ok := q.RemoveMax()
		if !ok {
			return vs
		}
		vs = append(vs, x)
	}
}
