package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/samber/lo"
	"github.com/yikakia/awf_calc/i18n"
	"github.com/yikakia/awf_calc/internal/cfgs"
	"github.com/yikakia/awf_calc/view"
	"github.com/yikakia/awf_calc/viewmodel"
)

func main() {

	a := app.New()
	w := a.NewWindow(i18n.MainTitle.T())
	w.SetMaster()

	bConfs := cfgs.GetBuildingConfigs()
	bs := lo.Map(bConfs, func(item *viewmodel.BuildingConfig, index int) *viewmodel.Building {
		return viewmodel.NewBuilding(item)
	})

	table := view.NewMainTable(&view.MainTableDep{
		Buildings: bs,
	})

	total := table.GetTotalTable()
	total.Resize(fyne.NewSize(800, 50))
	detail := table.GetDetailTableCanVas()
	c := container.NewVSplit(total, detail)
	c.SetOffset(0.1)
	w.SetContent(c)

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()
}
