package vi

import "testing"

type Parameters struct {
	Source   []int
	Target   int
	Expected []int
}

func TestGetRanges(t *testing.T) {
	src := []int{1, 4, 4, 10}
	target := 4
	want := []int{1, 2}
	paramList := []Parameters{}
	paramList = append(paramList, NewParameters(src, target))

	target = 10
	want = []int{3, 3}
	paramList = append(paramList, Parameters{src, target, want})

	for _, p := range paramList {
		UseParameters(t, p)
	}
}

func UseParameters(t *testing.T, parameters Parameters) {
	got := GetRanges(parameters.Source, parameters.Target)
	want := parameters.Expected
	in := parameters.Source
	if len(want) != len(got) {
		t.Errorf("Got %#v and not %#v for %#v", got, want, in)
	}
	for i, v := range got {
		if i >= len(want) || want[i] != v {
			t.Errorf("Got %#v and not %#v for %#v", got, want, in)
		}
	}
}
func NewParameters(src []int, target int) Parameters {
	p := Parameters{}
	p.Source = src
	p.Target = target
	p.Expected = NaiveGetRanges(src, target)
	return p
}
