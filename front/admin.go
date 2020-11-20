package front

import (
	"fmt"
	"github.com/massarakhsh/lik"
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
	tbl.AppendItem(rule.adminCmd("*", "Записать изменения", "write"))
	tbl.AppendItem(rule.adminCmd("*", "Удалить черновик", "cancel"))
	tbl.BuildTrTd("colspan=2").BuildString("<hr>")
	if row := tbl.BuildTr(); row != nil {
		row.BuildTdClass("title").BuildString("Код:")
		sel := likdom.BuildSpace()
		if canal.Code != "" {
			sel.BuildItemClass("B", "edit").BuildString(canal.Code)
		} else {
			sel.BuildString("(без кода)")
		}
		sel.BuildString("&nbsp;&nbsp;")
		sel.AppendItem(LinkTextProc("cmd", "Изменить", fmt.Sprintf("tube_code('%s')", canal.Code)))
		row.BuildTdClass("info").BuildItem("nobr").AppendItem(sel)
	}
	if row := tbl.BuildTr(); row != nil {
		row.BuildTdClass("title").BuildString("Наименование:")
		sel := likdom.BuildSpace()
		if canal.Name != "" {
			sel.BuildItemClass("B", "edit").BuildString(canal.Name)
		} else {
			sel.BuildString("(без имени)")
		}
		sel.BuildString("&nbsp;&nbsp;")
		sel.AppendItem(LinkTextProc("cmd", "Изменить", fmt.Sprintf("tube_name('%s')", canal.Name)))
		row.BuildTdClass("info").BuildItem("nobr").AppendItem(sel)
	}
	if row := tbl.BuildTr(); row != nil {
		row.BuildTdClass("title").BuildString("Формат:")
		sf := canal.Format
		if GetFormat(sf) == nil {
			sf = ""
		}
		sel := likdom.BuildItem("select", "onchange='tube_format(this)'")
		lsf := GetListFormats()
		for nf := 0; nf <= len(lsf); nf++ {
			opt := sel.BuildItem("option")
			if nf == 0 {
				opt.BuildString("Выберите формат")
			} else {
				opt.SetAttr("value", lsf[nf - 1].Code)
				opt.BuildString(lsf[nf - 1].Name)
			}
		}
		row.BuildTdClass("info").BuildItem("nobr").AppendItem(sel)
	}
	tbl.BuildTrTd("colspan=2").BuildString("<hr>")
	for ns := 0; ns < 4; ns++ {
		if row := tbl.BuildTr(); row != nil {
			row.BuildTdClass("title").BuildString(fmt.Sprintf("Окно %d", 1 + ns))
			sel := likdom.BuildSpace()
			path := ""
			if ns == 0 {
				path = canal.Source0
			} else if ns == 1 {
				path = canal.Source1
			} else if ns == 2 {
				path = canal.Source2
			} else if ns == 3 {
				path = canal.Source3
			}
			input := sel.BuildUnpairItem("input", "type=text", "id", fmt.Sprintf("src%d", ns))
			if path != "" {
				input.SetAttr("value", path)
			}
			sel.BuildString("&nbsp;&nbsp;")
			sel.AppendItem(LinkTextProc("cmd", "Изменить", fmt.Sprintf("tube_source(%d)", ns)))
			row.BuildTdClass("info").BuildItem("nobr").AppendItem(sel)
		}
	}
	return tbl
}

func (rule *DataRule) adminControlCommand(canal *one.Canal) likdom.Domer {
	tbl := likdom.BuildTable()
	tbl.AppendItem(rule.adminCmd("*", "Изменить", "edit"))
	tbl.AppendItem(rule.adminCmd("*", "Копировать", "copy"))
	tbl.AppendItem(rule.adminCmd("*", "Удалить", "delete"))
	tbl.AppendItem(rule.adminCmd("*", "Создать новый", "create"))
	return tbl
}

func (rule *DataRule) adminCmd(prompt string, text string, cmd string) likdom.Domer {
	row := likdom.BuildItem("tr")
	row.BuildTdClass("title").BuildString(prompt)
	row.BuildTdClass("info").AppendItem(LinkTextProc("cmd", text, fmt.Sprintf("tube_command('%s')", cmd)))
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
		if rule.ItPage.Variant == 0 {
			rule.adminDelete()
		}
	} else if rule.IsShift("cancel") {
		if rule.ItPage.Variant > 0 {
			rule.adminCancelEdit()
		}
	} else if rule.IsShift("write") {
		if rule.ItPage.Variant > 0 {
			rule.adminWrite()
		}
	} else if rule.IsShift("format") {
		if rule.ItPage.Variant > 0 {
			rule.adminFormat()
		}
	} else if rule.IsShift("code") {
		if rule.ItPage.Variant > 0 {
			rule.adminCode()
		}
	} else if rule.IsShift("name") {
		if rule.ItPage.Variant > 0 {
			rule.adminName()
		}
	} else if rule.IsShift("source") {
		if rule.ItPage.Variant > 0 {
			rule.adminSource()
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
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		if org,ok := one.GetCanalName(rule.ItPage.Canal, 0); ok {
			org.Delete()
		}
		canal.Variant = 0
		canal.Generate = rand.Intn(1000000000)
		canal.Update()
		rule.ItPage.Variant = 0
		rule.PageRedraw()
	}
}

func (rule *DataRule) adminCreate() {
	canal := one.Canal{}
	rule.ItPage.Variant = 1
	canal.Code = "common"
	canal.Name = "Новый"
	canal.Variant = rule.ItPage.Variant
	canal.Generate = rand.Intn(1000000000)
	canal.Format = "1"
	canal.Create()
	rule.PageRedraw()
}

func (rule *DataRule) adminDelete() {
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		canal.Delete()
		rule.PageRedraw()
	}
}

func (rule *DataRule) adminFormat() {
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		if val := lik.StringFromXS(rule.Shift()); val != canal.Format {
			canal.Format = val
			canal.Generate = rand.Intn(1000000000)
			canal.Update()
			rule.PageRedraw()
		}
	}
}

func (rule *DataRule) adminCode() {
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		if val := lik.StringFromXS(rule.Shift()); val != canal.Code {
			canal.Code = val
			canal.Generate = rand.Intn(1000000000)
			canal.Update()
			rule.ItPage.Canal = val
			rule.PageRedraw()
		}
	}
}

func (rule *DataRule) adminName() {
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		if val := lik.StringFromXS(rule.Shift()); val != canal.Name {
			canal.Name = val
			canal.Generate = rand.Intn(1000000000)
			canal.Update()
			rule.PageRedraw()
		}
	}
}

func (rule *DataRule) adminSource() {
	if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok {
		if ns := lik.StrToInt(rule.Shift()); ns >= 0 && ns < 4 {
			val := lik.StringFromXS(rule.Shift())
			if ns == 0 {
				canal.Source0 = val
			} else if ns == 1 {
				canal.Source1 = val
			} else if ns == 2 {
				canal.Source2 = val
			} else if ns == 3 {
				canal.Source3 = val
			}
			canal.Generate = rand.Intn(1000000000)
			canal.Update()
			rule.PageRedraw()
		}
	}
}

