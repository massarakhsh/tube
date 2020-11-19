package front

import (
	"fmt"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/tube/one"
	"math/rand"
)

func (rule *DataRule) AdminGen(sx int, sy int) likdom.Domer {
	code := likdom.BuildTableClass("fill")
	row := code.BuildTr()
	dx := sx /2 - BD
	row.BuildTd(MakeSizes(dx, sy)...).AppendItem(rule.adminLeft(dx, sy))
	row.BuildTd(MakeSizes(dx, sy)...).AppendItem(rule.adminRight(dx, sy))
	return code
}

func (rule *DataRule) adminLeft(sx int, sy int) likdom.Domer {
	code := likdom.BuildTableClass("fill")
	dy := sy / 3 - BD
	code.BuildTrTdClass("control", MakeSizes(sx, dy)...).AppendItem(rule.CanalsList(sx, dy))
	code.BuildTrTdClass("control", MakeSizes(sx, dy)...).AppendItem(rule.AdminControl(sx, dy))
	code.BuildTrTdClass("control", MakeSizes(sx, dy)...).AppendItem(rule.adminNo(sx, dy))
	return code
}

func (rule *DataRule) adminRight(sx int, sy int) likdom.Domer {
	code := likdom.BuildTableClass("fill")
	dy := sy / 2 - BD
	code.BuildTrTdClass("control", MakeSizes(sx, dy)...).AppendItem(rule.AdminShow(sx, dy))
	code.BuildTrTdClass("control", MakeSizes(sx, dy)...).AppendItem(rule.adminNo(sx, dy))
	return code
}

func (rule *DataRule) adminNo(sx int, sy int) likdom.Domer {
	return rule.VisualMessage("NO")
}

func (rule *DataRule) AdminShow(sx int, sy int) likdom.Domer {
	var text string
	rule.ItPage.ASX = sx
	rule.ItPage.ASY = sy
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		text = canal.Name
		if text == "" { text = rule.ItPage.Canal }
	} else {
		text = "Канал не выбран"
	}
	return MakeWindow("win_admshow", sx, sy, text, rule.VisualGen(sx, sy - 24 - BD))
}

func (rule *DataRule) AdminControl(sx int, sy int) likdom.Domer {
	rule.ItPage.ACX = sx
	rule.ItPage.ACY = sy
	tbl := likdom.BuildTable()
	title := "Новый канал"
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		title = MakeNamelyCanal(canal.Name, canal.Variant)
		if canal.Variant > 0 {
			tbl.BuildTrTd().AppendItem(rule.adminControlEdit(&canal))
		} else {
			tbl.BuildTrTd().AppendItem(rule.adminControlCommand(&canal))
		}
	} else {
		tbl.AppendItem(rule.adminCmd("&nbsp;*&nbsp;", "Создать новый", "create"))
	}
	return MakeWindow("win_control", sx, sy, title, tbl)
}

func (rule *DataRule) adminControlEdit(canal *one.Canal) likdom.Domer {
	tbl := likdom.BuildTable()
	tbl.AppendItem(rule.adminCmd("&nbsp;*&nbsp;", "Записать изменения", "write"))
	tbl.AppendItem(rule.adminCmd("&nbsp;*&nbsp;", "Удалить черновик", "cancel"))
	tbl.BuildTrTd("colspan=2").BuildString("<hr>")
	if row := tbl.BuildTr(); row != nil {
		row.BuildTd("Формат")
		
	}
	return tbl
}

func (rule *DataRule) adminControlCommand(canal *one.Canal) likdom.Domer {
	tbl := likdom.BuildTable()
	tbl.AppendItem(rule.adminCmd("&nbsp;*&nbsp;", "Изменить", "edit"))
	tbl.AppendItem(rule.adminCmd("&nbsp;*&nbsp;", "Копировать", "copy"))
	tbl.AppendItem(rule.adminCmd("&nbsp;*&nbsp;", "Удалить", "delete"))
	tbl.AppendItem(rule.adminCmd("&nbsp;*&nbsp;", "Создать новый", "create"))
	return tbl
}

func (rule *DataRule) adminCmd(prompt string, text string, cmd string) likdom.Domer {
	row := likdom.BuildItem("tr")
	row.BuildTd().BuildString(prompt)
	row.BuildTd().AppendItem(LinkTextProc("cmd", text, fmt.Sprintf("tube_control('%s')", cmd)))
	return row
}

func (rule *DataRule) ExecAdmin() {
	if rule.IsShift("control") {
		rule.ItPage.IsControl = !rule.ItPage.IsControl
		rule.PageRedraw()
	} else if rule.IsShift("edit") {
		if rule.ItPage.Variant == 0 {
			rule.adminToEdit()
		}
	} else if rule.IsShift("create") {
		rule.adminCreate()
	} else if rule.IsShift("delete") {
		rule.adminDelete()
	} else if rule.IsShift("cancel") {
		if rule.ItPage.Variant > 0 {
			rule.adminCancelEdit()
		}
	} else if rule.IsShift("write") {
		if rule.ItPage.Variant > 0 {
			rule.adminWrite()
		}
	}
}

func (rule *DataRule) adminToEdit() {
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, 0); ok {
		for variant := 1; variant < 100; variant++ {
			if _,ok := one.GetCanalName(rule.ItPage.Canal, variant); !ok {
				rule.ItPage.Variant = variant
				break
			}
		}
		canal.Variant = rule.ItPage.Variant
		canal.Generate = rand.Intn(1000000000)
		canal.Create()
		rule.SetGoPart(rule.PageGetPath())
	}
}

func (rule *DataRule) adminCancelEdit() {
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		canal.Delete()
		rule.ItPage.Variant = 0
		rule.SetGoPart(rule.PageGetPath())
	}
}

func (rule *DataRule) adminWrite() {
	if draft,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		if canal,ok := one.GetCanalName(rule.ItPage.Canal, 0); ok {
			canal.Variant = 0
			canal.Generate = rand.Intn(1000000000)
			canal.Update()
		} else {
			canal = draft
			canal.Variant = 0
			canal.Generate = rand.Intn(1000000000)
			canal.Create()
		}
		draft.Delete()
		rule.ItPage.Variant = 0
		rule.SetGoPart(rule.PageGetPath())
	}
}

func (rule *DataRule) adminCreate() {
	canal := one.Canal{}
	rule.ItPage.Variant = 1
	canal.Code = "common"
	canal.Name = "Новый"
	canal.Variant = rule.ItPage.Variant
	canal.Generate = rand.Intn(1000000000)
	canal.FormatId = 2
	canal.Create()
	rule.SetGoPart(rule.PageGetPath())
}

func (rule *DataRule) adminDelete() {
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, 0); ok {
		for variant := 1; variant < 100; variant++ {
			if _,ok := one.GetCanalName(rule.ItPage.Canal, variant); !ok {
				rule.ItPage.Variant = variant
				break
			}
		}
		canal.Variant = rule.ItPage.Variant
		canal.Generate = rand.Intn(1000000000)
		canal.Create()
		rule.SetGoPart(rule.PageGetPath())
	}
}


func makeCanalData() {
}
