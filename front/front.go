package front

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
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
	if rule.IsShift("canal") {
		rule.doCanal()
	} else if rule.IsShift("admin") {
		rule.doAdmin()
	} else if rule.IsShift("tune") {
		rule.doTune()
	} else if rule.IsShift("media") {
		rule.doMedia()
	} else if rule.IsShift("marshal") {
		rule.doMarshal()
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
		rule.PageRedraw()
	}
	if rule.ItPage.NeedUrl {
		rule.ItPage.NeedUrl = false
		rule.SetClientPath()
	}
}

func (rule *DataRule) doCanal() {
	if canal := rule.Shift(); canal != "" {
		rule.PageSetPath("/" + canal)
		rule.PageRedraw()
	}
}

func (rule *DataRule) doAdmin() {
	rule.ItPage.IsAdmin = !rule.ItPage.IsAdmin
	rule.PageRedraw()
}

func (rule *DataRule) doTune() {
	key := rule.Shift()
	value := lik.StringFromXS(rule.Shift())
	rule.AdminTuneCanal(key, value)
	rule.PageRedraw()
}

func (rule *DataRule) doMedia() {
	key := rule.Shift()
	rule.MediaExecute(key)
}

func (rule *DataRule) PageRedraw() {
	rule.StorePage()
}

func (rule *DataRule) SetClientPath() {
	rule.SetOnPart("/" + rule.ItPage.Canal)
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
	var code likdom.Domer
	if rule.ItPage.IsAdmin {
		code = rule.AdminGen(sx -BD, sy -BD)
	} else {
		code = rule.CanalGen(sx -BD, sy -BD)
	}
	if code != nil {
		div.AppendItem(code)
	}
	return div
}

func (rule *DataRule) PageSetPath(path string) {
	parts := lik.PathToNames(path)
	canal := "common"
	if len(parts) > 0 {
		canal = parts[0]
		parts = parts[1:]
	}
	if canal != rule.ItPage.Canal || true {
		rule.SetCookie(canal,"canal")
		rule.ItPage.Canal = canal
		rule.ItPage.NeedUrl = true
	}
}

func (rule *DataRule) PageGetPath() string {
	path := "/"
	if rule.ItPage.Canal != "" {
		path += rule.ItPage.Canal
	} else {
		path += "common"
	}
	return path
}

func (rule *DataRule) BuildSidebar() likdom.Domer {
	div := likdom.BuildItem("div", "class=sidebar")
	ul := div.BuildItem("ul","id=topon")
	ul.BuildItem("li","class=topon").
		BuildItem("a","title=Настройка", "href=#", "onclick=topon_click()")
	return div
}

