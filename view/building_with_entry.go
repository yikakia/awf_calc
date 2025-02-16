package view

import (
	"strconv"

	"github.com/yikakia/awf_calc/internal/tools/gresult"
	"github.com/yikakia/awf_calc/mywidget"
	"github.com/yikakia/awf_calc/viewmodel"
)

type buildingWithEntry struct {
	*viewmodel.Building
	*mywidget.NumericalEntry
}

func newBuildingWithEntry(b *viewmodel.Building) *buildingWithEntry {
	entry := mywidget.NewNumericalEntry()

	entry.OnChanged = func(s string) {
		cnts := gresult.Of(strconv.Atoi(s)).ValueOrZero()
		b.Data.Count.Set(cnts)
	}

	bd := &buildingWithEntry{
		Building:       b,
		NumericalEntry: entry,
	}
	return bd
}
