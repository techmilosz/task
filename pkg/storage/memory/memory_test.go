package memory

import (
	"testing"

	"repartners/pkg/tests"
)

func TestGetAll(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		vals map[int]struct{}
		want []int
	}

	cases := []testCase{
		{
			name: "get happy path",
			vals: map[int]struct{}{1: {}, 2: {}, 3: {}},
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m := New[int]()
			m.values = tc.vals
			res := m.GetAll()

			if !tests.CompareUnorderedSlice(tc.want, res) {
				t.Fatalf("expected %v, got %v", tc.want, res)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		vals []int
		want []int
	}

	cases := []testCase{
		{
			name: "add happy path - 0 values",
			vals: []int{},
			want: []int{},
		},
		{
			name: "add happy path - 3 values",
			vals: []int{3, 1, 2},
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m := New[int]()
			for _, val := range tc.vals {
				m.Add(val)
			}

			res := m.GetAll()
			if !tests.CompareUnorderedSlice(tc.want, res) {
				t.Fatalf("expected %v, got %v", tc.want, res)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		valsAdd    []int
		valsRemove []int
		want       []int
	}

	cases := []testCase{
		{
			name:       "happy path - add nothing",
			valsAdd:    []int{},
			valsRemove: []int{},
			want:       []int{},
		},
		{
			name:       "happy path - add nothing, remove non existing",
			valsAdd:    []int{},
			valsRemove: []int{1},
			want:       []int{},
		},
		{
			name:       "happy path - add something, remove non existing",
			valsAdd:    []int{1},
			valsRemove: []int{2},
			want:       []int{1},
		},
		{
			name:       "happy path - add and remove item",
			valsAdd:    []int{4},
			valsRemove: []int{4},
			want:       []int{},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m := New[int]()
			for _, val := range tc.valsAdd {
				m.Add(val)
			}

			for _, val := range tc.valsRemove {
				m.Remove(val)
			}

			res := m.GetAll()
			if !tests.CompareUnorderedSlice(tc.want, res) {
				t.Fatalf("expected %v, got %v", tc.want, res)
			}
		})
	}
}
