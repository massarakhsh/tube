package front

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likbase"
	"github.com/massarakhsh/lik/likdom"
	"fmt"
)

func (rule *DataRule) AdminGen(sx int, sy int) likdom.Domer {
	code := likdom.BuildTable("class=fill")
	row := code.BuildTr()
	row.BuildTd("width=50%", "height=100%").AppendItem(rule.adminLeft(sx / 2 -BD, sy))
	row.BuildTd("width=50%", "height=100%").AppendItem(rule.adminRight(sx / 2 -BD, sy))
	return code
}

func (rule *DataRule) adminLeft(sx int, sy int) likdom.Domer {
	return rule.adminControl(sx, sy)
}

func (rule *DataRule) adminRight(sx int, sy int) likdom.Domer {
	code := likdom.BuildTable("class=fill")
	code.BuildTrTd("width=100%", "height=60%").AppendItem(rule.adminTop(sx, sy * 6 / 10 -BD))
	code.BuildTrTd("width=100%", "height=40%").AppendItem(rule.adminBottom(sx, sy * 4 / 10 -BD))
	return code
}

func (rule *DataRule) adminTop(sx int, sy int) likdom.Domer {
	return rule.adminCanal(sx, sy)
}

func (rule *DataRule) adminBottom(sx int, sy int) likdom.Domer {
	return rule.adminMedia(sx, sy)
}

func (rule *DataRule) adminControl(sx int, sy int) likdom.Domer {
	code := likdom.BuildTable("class=control", "id=control")
	code.BuildTrTdClass("bottom","height=60px").BuildString("&nbsp;")
	code.BuildTrTd().AppendItem(rule.adminListCanals(sx))
	if canal := FindCanal(rule.ItPage.Canal); canal != nil {
		code.BuildTrTd().BuildString("&nbsp;")
		code.BuildTrTd().AppendItem(rule.adminControlCanal(sx, canal))
	}
	code.BuildTrTdClass("fill")
	return code
}

func (rule *DataRule) adminCanal(sx int, sy int) likdom.Domer {
	code := likdom.BuildTable("class=fill", "id=control")
	var text string
	if canal := FindCanal(rule.ItPage.Canal); canal != nil {
		text = canal.GetString("title")
		if text == "" { text = rule.ItPage.Canal }
	} else {
		text = "Канал не выбран"
	}
	code.BuildTrTdClass("boxtitle").BuildString(text)
	code.BuildTrTdClass("boxcell").AppendItem(rule.CanalName(sx, sy - 20, rule.ItPage.Canal))
	return code
}

func (rule *DataRule) adminListCanals(sx int) likdom.Domer {
	tbl := likdom.BuildTable()
	tbl.BuildTrTdClass("boxtitle", "colspan=3").BuildString("Список каналов")
	for _, elm := range TableCanal.Elms {
		row := tbl.BuildTr()
		idc := elm.GetString("idc")
		name,isedit := rule.nameFromEdit(idc)
		if isedit {
			name += " (черновик)"
		}
		title := elm.GetString("title")
		cls := "boxcontrol"
		row.BuildTdClass(cls).AppendItem(LinkPicCmd("","/images/st.gif","/canal/"+idc,"Включить канал"))
		if name == rule.ItPage.Canal {
			cls += " boxsel"
		}
		row.BuildTdClass(cls).BuildString(name)
		row.BuildTdClass(cls).BuildString(title)
	}
	if row := tbl.BuildTr(); row != nil {
		row.BuildTdClass("boxcontrol").BuildString("&nbsp;")
		row.BuildTdClass("boxcontrol").BuildString("&nbsp;")
		row.BuildTdClass("boxcontrol").AppendItem(LinkTextProc("cmd", "Добавить канал", "tune_append()"))
	}
	return tbl
}

func (rule *DataRule) adminControlCanal(sx int, canal *likbase.ItElm) likdom.Domer {
	tbl := likdom.BuildTable("class=control")
	name,isedit := rule.nameFromEdit(rule.ItPage.Canal)
	if isedit {
		name += " (черновик)"
		rule.ItPage.IsEdit = true
	} else {
		rule.ItPage.IsEdit = false
	}
	tbl.BuildTrTdClass("boxtitle").BuildString("Канал '" + name + "'")
	if rule.ItPage.IsEdit {
		tbl.BuildTrTdClass("boxcontrol").AppendItem(LinkTextProc("cmd", "Записать изменения", "tune_edit_write()"))
		tbl.BuildTrTdClass("boxcontrol").AppendItem(LinkTextProc("cmd", "Отменить изменения", "tune_edit_cancel()"))
	} else {
		tbl.BuildTrTdClass("boxcontrol").AppendItem(LinkTextProc("cmd", "Внести изменения", "tune_edit_start()"))
		tbl.BuildTrTdClass("boxcontrol").AppendItem(LinkTextProc("cmd", "Удалить канал", "tune_canal_delete()"))
	}
	tbl.BuildTrTdClass("boxcontrol").BuildString("<hr>")
	tbl.BuildTrTd().AppendItem(rule.adminCanalInfo(sx, canal))
	return tbl
}

func (rule *DataRule) adminCanalInfo(sx int, canal *likbase.ItElm) likdom.Domer {
	code := likdom.BuildTable("class=control")
	if row := code.BuildTr(); row != nil {
		row.BuildTdClass("boxcontrol").BuildString("Код канала")
		value,_ := rule.nameFromEdit(canal.GetString("idc"))
		row.BuildTdClass("boxcontrol info").BuildString(value)
		td := row.BuildTdClass("boxcontrol")
		if rule.ItPage.IsEdit {
			td.AppendItem(LinkTextProc("cmd", "Изменить", fmt.Sprintf("tune_canal('idc','%s')", value)))
		}
		row.BuildTdClass("fill")
	}
	if row := code.BuildTr(); row != nil {
		row.BuildTdClass("boxcontrol").BuildString("Наименование канала")
		value := canal.GetString("title")
		row.BuildTdClass("boxcontrol info").BuildString(value)
		td := row.BuildTdClass("boxcontrol")
		if rule.ItPage.IsEdit {
			td.AppendItem(LinkTextProc("cmd", "Изменить", fmt.Sprintf("tune_canal('title','%s')", value)))
		}
		row.BuildTdClass("fill")
	}
	if row := code.BuildTr(); row != nil {
		row.BuildTdClass("boxcontrol").BuildString("Код страницы")
		html := likdom.BuildItem("textarea", "id=canalcode")
		html.SetAttr("style", fmt.Sprintf("width:%dpx;height:%dpx;", sx - 250, (sx - 250) / 2))
		html.BuildString(canal.Info.Format(""))
		if !rule.ItPage.IsEdit {
			html.SetAttr("readonly", "")
		}
		row.BuildTdClass("boxcontrol info").AppendItem(html)
		td := row.BuildTdClass("boxcontrol")
		if rule.ItPage.IsEdit {
			td.AppendItem(LinkTextProc("cmd", "Изменить", "tune_canal_code()"))
		}
		row.BuildTdClass("fill")
	}
	return code
}

func (rule *DataRule) AdminTuneCanal(key string, value string) {
	if key == "append" {
		name,_ := rule.nameFromEdit(value)
		if canal := FindCanal(name); canal == nil {
			canal := TableCanal.CreateElm()
			canal.SetValue(name + "_", "idc")
		}
		rule.ItPage.Canal = name + "_"
	} else if canal := FindCanal(rule.ItPage.Canal); canal != nil {
		if key == "idc" {
			name,_ := rule.nameFromEdit(value)
			if FindCanal(name) == nil && FindCanal(name + "_") == nil {
				canal.SetValue(name + "_", "idc")
			}
		} else if key == "title" {
			canal.SetValue(value, "title")
		} else if key == "canalcode" {
			if info := lik.SetFromRequest(value); info != nil {
				canal.Info = info
				canal.OnModify()
			}
		} else if key == "editstart" {
			if name,isedit := rule.nameFromEdit(rule.ItPage.Canal); !isedit {
				if edit := FindCanal(name + "_"); edit == nil {
					edit := TableCanal.CreateElm()
					edit.Info = canal.Info.Clone().ToSet()
					edit.SetValue(name + "_", "idc")
				}
				rule.ItPage.Canal = name + "_"
			}
		} else if key == "editwrite" {
			if name,isedit := rule.nameFromEdit(rule.ItPage.Canal); isedit {
				hard := FindCanal(name)
				if hard == nil {
					hard = TableCanal.CreateElm()
				}
				hard.Info = canal.Info.Clone().ToSet()
				hard.SetValue(name, "idc")
				canal.Delete()
				rule.ItPage.Canal = name
			}
		} else if key == "editcancel" {
			if name,isedit := rule.nameFromEdit(rule.ItPage.Canal); isedit {
				canal.Delete()
				rule.ItPage.Canal = name
			}
		} else if key == "delete" {
			if _,isedit := rule.nameFromEdit(rule.ItPage.Canal); !isedit {
				canal.Delete()
				rule.ItPage.Canal = ""
			}
		}
	}
}

func (rule *DataRule) nameFromEdit(name string) (string,bool) {
	edit,ok := name,false
	if match := lik.RegExParse(name, "(.*)_$"); match != nil {
		edit = match[1]
		ok = true
	}
	return edit,ok
}

func (rule *DataRule) nameToEdit(name string) (string,bool) {
	edit,ok := name,false
	if match := lik.RegExParse(name, "(.*)_$"); match == nil {
		edit += "_"
		ok = true
	}
	return edit,ok
}

