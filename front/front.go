package front

import (
	"fmt"
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/tube/one"
	"strings"
)

const BD = 4

func (rule *DataRule) PageHtml() (int, likdom.Domer) {
	html := rule.InitializePage(Version)
	if head, _ := html.GetDataTag("head"); head != nil {
		head.BuildItem("title").BuildString("АО РПТП Гранит")
		head.BuildString("<link rel='stylesheet' href='/lib/datatables.css'/>")
		head.BuildString("<script type='text/javascript' src='/lib/datatables.js'></script>")
		head.BuildString("<link rel='stylesheet' href='/lib/dropzone.css'/>")
		head.BuildString("<script type='text/javascript' src='/lib/dropzone.js'></script>")
		head.BuildString("<link rel='stylesheet' href='/lib/fotorama.css'/>")
		head.BuildString("<script type='text/javascript' src='/lib/fotorama.js'></script>")
		head.BuildString("<link rel='stylesheet' href='/js/grid.css?" + Version + "'/>")
		head.BuildString("<script type='text/javascript' src='/js/grid.js?" + Version + "'></script>")
		head.BuildString("<link rel='stylesheet' href='/js/styles.css?" + Version + "'/>")
		head.BuildString("<script type='text/javascript' src='/js/script.js?" + Version + "'></script>")
	}
	if body, _ := html.GetDataTag("body"); body != nil {
		if script := body.BuildItem("script"); script != nil {
			script.BuildString("jQuery(document).ready(function () { script_start(); });")
		}
		rule.BuildPage(body)
	}
	return 200, html
}

func (rule *DataRule) BuildPage(pater likdom.Domer) {
	path := rule.BuildPath()
	if path == "" || path == "/" {
		path = "/" + rule.GetCookie("canal")
		if path == "" || path == "/" {
			path = "/common"
		}
	}
	rule.PageSetPath(path)
	pater.AppendItem(rule.GenPage())
	pater.AppendItem(rule.BuildSidebar())
}

func (rule *DataRule) BuildFront() (int,lik.Seter) {
	rule.SeekPageSize()
	if rule.IsShift("marshal") {
		rule.doMarshal()
	} else if rule.IsShift("admin") {
		rule.ExecAdmin()
	} else if rule.IsShift("media") {
		rule.ExecMedia()
	} else if rule.IsShift("canal") {
		rule.ExecCanal()
	}
	return 200, rule.GetAllResponse()
}

func (rule *DataRule) doMarshal() {
	if !rule.Page.GetTrust() {
		rule.ItPage.ToPath = "/"
	}
	if rule.ItPage.ToPath != "" {
		rule.SetGoPart(rule.ItPage.ToPath)
		rule.ItPage.ToPath = ""
	} else if _,_,need := rule.ItPage.GetSizeFix(); need {
		rule.ItPage.NeedDraw = true
	} else if canal,ok := one.GetCanalName(rule.ItPage.Canal, rule.ItPage.Variant); ok && rule.ItPage.Generate != canal.Generate {
		rule.ItPage.NeedDraw = true
	}
	if rule.ItPage.NeedDraw {
		rule.ItPage.NeedDraw = false
		rule.PageRedraw()
	}
	if rule.ItPage.NeedImage {
		rule.ItPage.NeedImage = false
		rule.MediaImageShow()
	}
	if rule.ItPage.NeedUrl {
		rule.ItPage.NeedUrl = false
		rule.SetOnPart(rule.PageGetPath())
	}
}

func (rule *DataRule) PageRedraw() {
	rule.StorePage()
}

func (rule *DataRule) collectParms(prefix string) lik.Seter {
	parms := lik.BuildSet()
	if context := rule.GetAllContext(); context != nil {
		for _,set := range(context.Values()) {
			if strings.HasPrefix(set.Key, prefix) && set.Val != nil {
				str := set.Val.ToString()
				parms.SetItem(str, set.Key[len(prefix):])
			}
		}
	}
	return parms
}

func (rule *DataRule) StorePage() {
	rule.StoreItem(rule.GenPage())
}

func (rule *DataRule) GenPage() likdom.Domer {
	div := likdom.BuildItem("div","id=main", "class=fill")
	sx,sy := rule.ItPage.GetSize()
	if rule.ItPage.IsControl {
		div.AppendItem(rule.AdminGen(sx -8-BD, sy -8-BD))
	} else {
		div.AppendItem(rule.VisualGen(sx -8-BD, sy -8-BD))
	}
	return div
}

func (rule *DataRule) PageSetPath(path string) {
	parts := lik.PathToNames(path)
	canal := "common"
	variant := 0
	if len(parts) > 0 {
		canal = parts[0]
		parts = parts[1:]
		if len(parts) > 0 {
			variant = lik.StrToInt(parts[0])
			parts = parts[1:]
		}
	}
	rule.PageSetCanal(canal, variant)
}

func (rule *DataRule) PageSetCanal(canal string, variant int) {
	if canal != rule.ItPage.Canal {
		rule.SetCookie(canal,"canal")
		rule.ItPage.Canal = canal
		rule.ItPage.NeedUrl = true
	}
	if variant != rule.ItPage.Variant {
		rule.ItPage.Variant = variant
	}
}

func (rule *DataRule) PageGetPath() string {
	path := "/"
	if rule.ItPage.Canal != "" {
		path += rule.ItPage.Canal
	} else {
		path += "common"
	}
	if rule.ItPage.Variant > 0 {
		path += fmt.Sprintf("/%d", rule.ItPage.Variant)
	}
	return path
}

func (rule *DataRule) BuildSidebar() likdom.Domer {
	div := likdom.BuildItem("div", "class=sidebar")
	ul := div.BuildItem("ul","id=topon")
	ul.BuildItem("li","class=topon").
		BuildItem("a","title=Настройка", "href=#", "onclick=tube_command('/admin/control')")
	return div
}

