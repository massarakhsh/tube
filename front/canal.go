package front

import (
	"fmt"
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/lik/liktable"
	"github.com/massarakhsh/tube/one"
)

func (rule *DataRule) CanalsList(sx int, sy int) likdom.Domer {
	rule.canalsPrepare()
	table := rule.ItPage.ListCanals.Initialize("/canal/list")
	return MakeWindow("win_canals", sx, sy, "Список каналов", table)
}

func (rule *DataRule) ExecCanal() {
	if rule.IsShift("listinit") {
		rule.canalsListInit()
	} else if rule.IsShift("listdata") {
		rule.canalsListData()
	} else if rule.IsShift("listselect") {
		rule.canalsListSelect()
	}
}

func (rule *DataRule) canalsListInit() {
	rule.canalsPrepare()
	grid := rule.ItPage.ListCanals.Show()
	grid.SetItem(rule.canalsBuildList(), "data")
	grid.SetItem(rule.ItPage.IdCanal, "likSelect")
	rule.SetResponse(grid, "grid")
}

func (rule *DataRule) canalsListData() {
	rule.canalsPrepare()
	rule.SetResponse(lik.BuildList(), "data")
	//rule.SetResponse(rule.canalsBuildList(), "data")
}

func (rule *DataRule) canalsListSelect() {
	rule.canalsPrepare()
	if id := lik.StrToIDB(rule.Shift()); id != 0 {
		rule.ItPage.IdCanal = id
		if canal,ok := one.GetCanal(id); ok {
			rule.ItPage.Canal = canal.Code
			rule.ItPage.Variant = canal.Variant
			rule.AdminReControl()
			rule.AdminReShow()
			rule.SetOnPart("/" + canal.Code)
		}

	}
}

func (rule *DataRule) canalsBuildList() lik.Lister {
	list := lik.BuildList()
	var canals []one.Canal
	one.DBCanal().Find(&canals)
	for _, canal := range canals {
		code := canal.Code
		if canal.Variant > 0 {
			code += fmt.Sprintf(" (%d)", canal.Variant)
		}
		list.AddItemSet("id", canal.Id, "code", code, "name", canal.Name)
		if canal.Code == rule.ItPage.Canal {
			rule.ItPage.IdCanal = lik.IDB(canal.Id)
		}
	}
	return list
}

func (rule *DataRule) canalsPrepare() {
	if rule.ItPage.ListCanals == nil {
		rule.ItPage.ListCanals = liktable.New("serverSide=false", "pageLength=5")
		rule.ItPage.ListCanals.AddColumn("data=code", "title=Код", "width=50")
		rule.ItPage.ListCanals.AddColumn("data=name", "title=Наименование", "width=150")
	}
}

