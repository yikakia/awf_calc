package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	i18n2 "github.com/yikakia/awf_calc/i18n"
)

func buildUpdateHeaderFunc(headers []i18n2.Key) func(id widget.TableCellID, template fyne.CanvasObject) {
	return func(id widget.TableCellID, template fyne.CanvasObject) {
		if id.Row >= 0 {
			return
		}
		template.(*widget.Label).SetText(headers[id.Col].T())
	}
}

var detailHeaders = []i18n2.Key{
	//"建筑",
	i18n2.TableDetailBuildingName,
	//"建筑个数",
	i18n2.TableDetailBuildingCountName,
	//"工人类别",
	i18n2.TableDetailCitizenTypeName,
	//"工人总数",
	i18n2.TableDetailCitizenCountName,
	//"电力",
	i18n2.ResourceEnergyName,
	//"水",
	i18n2.ResourceWaterName,
	//"木头",
	i18n2.ResourceWoodName,
	//"垃圾",
	i18n2.ResourceJunkName,
	//"食物",
	i18n2.ResourceFoodName,
	//"金属",
	i18n2.ResourceMetalName,
	//"燃料",
	i18n2.ResourceFuelName,
}

var totalTableHeader = []i18n2.Key{
	i18n2.CitizenWorkerName,
	i18n2.CitizenSailorName,
	i18n2.CitizenEngineerName,
	i18n2.ResourceEnergyName,
	//"水",
	i18n2.ResourceWaterName,
	//"木头",
	i18n2.ResourceWoodName,
	//"垃圾",
	i18n2.ResourceJunkName,
	//"食物",
	i18n2.ResourceFoodName,
	//"金属",
	i18n2.ResourceMetalName,
	//"燃料",
	i18n2.ResourceFuelName,
}
