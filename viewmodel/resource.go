package viewmodel

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/samber/lo"
	"github.com/yikakia/awf_calc/game/resource"
	"github.com/yikakia/awf_calc/internal/tools/gresult"
)

type AllResources struct {
	rs map[resource.Type]binding.Float
}

func NewAllResources() *AllResources {
	rs := make(map[resource.Type]binding.Float)
	for _, t := range resource.AllTypes() {
		rs[t] = binding.NewFloat()
	}
	return &AllResources{
		rs: rs,
	}
}

func (a *AllResources) Clear() {
	for r := range a.rs {
		a.Set(r, 0)
	}
}

func (a *AllResources) Set(p resource.Type, v float64) {
	a.rs[p].Set(v)
}

func (a *AllResources) Add(p resource.Type, v float64) {
	a.Set(p, gresult.Of(a.rs[p].Get()).MustGet()+v)
}

func (a *AllResources) GetBy(t resource.Type) binding.Float {
	return a.rs[t]
}

func (a *AllResources) GetByAsString(t resource.Type, format string) binding.String {
	return binding.FloatToStringWithFormat(a.rs[t], format)
}

func (a *AllResources) BuildLabels() []*widget.Label {
	return lo.Map(resource.AllTypes(), func(item resource.Type, index int) *widget.Label {
		countStr := binding.FloatToStringWithFormat(a.rs[item], "%.2f")
		return widget.NewLabelWithData(countStr)
	})
}
