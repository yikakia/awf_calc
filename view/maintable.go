package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/samber/lo"
	"github.com/yikakia/awf_calc/game/citizen"
	"github.com/yikakia/awf_calc/game/resource"
	"github.com/yikakia/awf_calc/i18n"
	"github.com/yikakia/awf_calc/internal/tools/gresult"
	viewmodel2 "github.com/yikakia/awf_calc/viewmodel"
)

type MainTable struct {
	buildings    []*buildingWithEntry
	detailTable  *widget.Table
	totalTable   *widget.Table
	onChanges    []func()
	peopleCnt    binding.Int
	allResources *viewmodel2.AllResources
	citizenCnt   map[citizen.Type]binding.Int
}

type MainTableDep struct {
	Buildings []*viewmodel2.Building
}

func (m *MainTable) GetDetailTableCanVas() fyne.CanvasObject {
	m.buildDetailTable()
	return m.detailTable
}

func NewMainTable(dep *MainTableDep, opts ...func(table *MainTable)) *MainTable {
	m := &MainTable{}
	m.allResources = viewmodel2.NewAllResources()
	m.citizenCnt = make(map[citizen.Type]binding.Int)
	for _, t := range citizen.AllTypes() {
		m.citizenCnt[t] = binding.NewInt()
	}

	m.buildings = lo.Map(dep.Buildings, func(item *viewmodel2.Building, _ int) *buildingWithEntry {
		return newBuildingWithEntry(item)
	})

	lo.ForEach(dep.Buildings, func(item *viewmodel2.Building, _ int) {
		item.Data.Count.AddListener(binding.NewDataListener(m.sumAllResources))
		item.Data.CitizenCount.AddListener(binding.NewDataListener(m.sumPeopleCnt))
	})

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (m *MainTable) buildDetailTable() {
	table := widget.NewTableWithHeaders(
		func() (rows int, cols int) {
			return len(m.buildings), len(detailHeaders)
		},
		func() fyne.CanvasObject {
			return container.NewStack()
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			c := object.(*fyne.Container)
			c.RemoveAll()
			c.Add(m.getDetailCol(id))
		},
	)

	table.CreateHeader = func() fyne.CanvasObject { return widget.NewLabel("") }
	table.UpdateHeader = buildUpdateHeaderFunc(detailHeaders)

	// 调整列宽
	for i, header := range detailHeaders {
		table.SetColumnWidth(i, max(float32(len(header.T())*8), 50))
		if i == 0 {
			table.SetColumnWidth(i, max(float32(len(header.T())*8), 90))

		}
	}
	m.detailTable = table
}

func (m *MainTable) getBuildings() []*buildingWithEntry {
	return m.buildings
}

func (m *MainTable) getDetailCol(id widget.TableCellID) fyne.CanvasObject {
	b := m.buildings[id.Row]
	switch detailHeaders[id.Col] {
	case i18n.TableDetailBuildingName:
		return widget.NewLabel(b.Data.Name)
	case i18n.TableDetailBuildingCountName:
		return b.NumericalEntry
	case i18n.TableDetailCitizenTypeName:
		return widget.NewLabel(b.Data.CitizenType.String())
	case i18n.TableDetailCitizenCountName:
		return widget.NewLabelWithData(binding.IntToString(b.Data.CitizenCount))
	case i18n.ResourceEnergyName:
		return widget.NewLabelWithData(b.Data.GetByAsString(resource.ResourceEnergy, "%.2f"))
	case i18n.ResourceWaterName:
		return widget.NewLabelWithData(b.Data.GetByAsString(resource.ResourceWater, "%.2f"))
	case i18n.ResourceWoodName:
		return widget.NewLabelWithData(b.Data.GetByAsString(resource.ResourceWood, "%.2f"))
	case i18n.ResourceJunkName:
		return widget.NewLabelWithData(b.Data.GetByAsString(resource.ResourceJunk, "%.2f"))
	case i18n.ResourceFoodName:
		return widget.NewLabelWithData(b.Data.GetByAsString(resource.ResourceFood, "%.2f"))
	case i18n.ResourceMetalName:
		return widget.NewLabelWithData(b.Data.GetByAsString(resource.ResourceMetal, "%.2f"))
	case i18n.ResourceFuelName:
		return widget.NewLabelWithData(b.Data.GetByAsString(resource.ResourceFuel, "%.2f"))
	default:
		return widget.NewLabel("unknown col")
	}

}

func (m *MainTable) updateDetailTableHeader(id widget.TableCellID, template fyne.CanvasObject) {
	if id.Row >= 0 {
		return
	}
	template.(*widget.Label).SetText(detailHeaders[id.Col].T())
}

func (m *MainTable) sumPeopleCnt() {
	grouped := lo.GroupBy(m.buildings, func(item *buildingWithEntry) citizen.Type {
		return item.Data.CitizenType
	})

	for cType, bs := range grouped {
		cCnt := lo.SumBy(bs, func(item *buildingWithEntry) int {
			return gresult.Of(item.Data.CitizenCount.Get()).MustGet()
		})
		m.citizenCnt[cType].Set(cCnt)
	}
}

func (m *MainTable) sumAllResources() {
	m.allResources.Clear()
	for _, t := range resource.AllTypes() {
		for _, building := range m.buildings {
			mustGet := gresult.Of(building.Building.Data.AllResources.GetBy(t).Get()).MustGet()
			m.allResources.Add(t, mustGet)
		}
	}
}

func (m *MainTable) GetTotalTable() fyne.CanvasObject {
	m.buildTotalTable()
	return m.totalTable
}

func (m *MainTable) buildTotalTable() {
	table := widget.NewTableWithHeaders(
		func() (rows int, cols int) {
			return 1, len(totalTableHeader)
		},
		func() fyne.CanvasObject {
			return container.NewStack()
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			c := object.(*fyne.Container)
			c.RemoveAll()
			c.Add(m.getTotalCol(id))
		},
	)

	table.CreateHeader = func() fyne.CanvasObject { return widget.NewLabel("") }
	table.UpdateHeader = buildUpdateHeaderFunc(totalTableHeader)

	// 调整列宽
	for i, header := range detailHeaders {
		table.SetColumnWidth(i, max(float32(len(header)*3), 50))
	}
	table.SetRowHeight(0, 10)
	m.totalTable = table
}

func (m *MainTable) updateTotalTableHeader(id widget.TableCellID, template fyne.CanvasObject) {
	if id.Row >= 0 {
		return
	}
	template.(*widget.Label).SetText(totalTableHeader[id.Col].T())
}

func (m *MainTable) getTotalCol(id widget.TableCellID) fyne.CanvasObject {
	switch totalTableHeader[id.Col] {
	case i18n.CitizenWorkerName:
		return widget.NewLabelWithData(m.getCitizenCntByTypeAsString(citizen.Worker))
	case i18n.CitizenSailorName:
		return widget.NewLabelWithData(m.getCitizenCntByTypeAsString(citizen.Sailor))
	case i18n.CitizenEngineerName:
		return widget.NewLabelWithData(m.getCitizenCntByTypeAsString(citizen.Sailor))
	case i18n.ResourceEnergyName:
		return widget.NewLabelWithData(m.allResources.GetByAsString(resource.ResourceEnergy, "%.2f"))
	case i18n.ResourceWaterName:
		return widget.NewLabelWithData(m.allResources.GetByAsString(resource.ResourceWater, "%.2f"))
	case i18n.ResourceWoodName:
		return widget.NewLabelWithData(m.allResources.GetByAsString(resource.ResourceWood, "%.2f"))
	case i18n.ResourceJunkName:
		return widget.NewLabelWithData(m.allResources.GetByAsString(resource.ResourceJunk, "%.2f"))
	case i18n.ResourceFoodName:
		return widget.NewLabelWithData(m.allResources.GetByAsString(resource.ResourceFood, "%.2f"))
	case i18n.ResourceMetalName:
		return widget.NewLabelWithData(m.allResources.GetByAsString(resource.ResourceMetal, "%.2f"))
	case i18n.ResourceFuelName:
		return widget.NewLabelWithData(m.allResources.GetByAsString(resource.ResourceFuel, "%.2f"))
	default:
		return widget.NewLabel("unknown col")
	}
}

func (m *MainTable) getCitizenCntByTypeAsString(t citizen.Type) binding.String {
	return binding.IntToString(m.citizenCnt[t])
}
