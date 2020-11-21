package front

import (
	"fmt"
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"io/ioutil"
	"sort"
	"strings"
)

var dirMain = "media"

func (rule *DataRule) MediaImageShow() {
	if !rule.ItPage.IsLoad {
		rule.MediaReImage()
	} else {
		rule.mediaReCommand()
	}
}

func (rule *DataRule) MediaReShow() {
	rule.StoreItem(rule.MediaShow(rule.ItPage.MediaSize.Sx, rule.ItPage.MediaSize.Sy))
}

func (rule *DataRule) MediaShow(sx int, sy int) likdom.Domer {
	rule.ItPage.MediaSize = Size{sx, sy}
	code := likdom.BuildTableClassId("fill", "win_media")
	dx := sx /2 - BD
	dy := sy - 32
	code.BuildTrTd("colspan=2").AppendItem(rule.mediaCommands())
	row := code.BuildTr()
	if !rule.ItPage.IsLoad {
		row.BuildTd(MakeSizes(dx, dy)...).AppendItem(rule.mediaBrowser(dx, dy))
		row.BuildTd(MakeSizes(dx, dy)...).AppendItem(rule.mediaImage(dx, dy))
	} else {
		td := row.BuildTd("colspan=2")
		td.SetAttr(MakeSizes(sx, dy)...)
		td.AppendItem(rule.mediaLoad(sx, dy))
	}
	return code
}

func (rule *DataRule) mediaReCommand() {
	rule.StoreItem(rule.mediaCommands())
}

func (rule *DataRule) mediaCommands() likdom.Domer {
	tbl := likdom.BuildTableClassId("mediacmd", "mediacmd")
	row := tbl.BuildTr()
	if !rule.ItPage.IsLoad {
		row.BuildTd().AppendItem(LinkTextProc("cmd", "Загрузить", "media_load()"))
		row.BuildTd().AppendItem(LinkTextProc("cmd", "Удалить", "media_delete()"))
	} else if len(rule.ItPage.Upload) > 0 {
		row.BuildTd().AppendItem(LinkTextProc("cmd", "Запомнить", "media_store()"))
		row.BuildTd().AppendItem(LinkTextProc("cmd", "Отменить", "media_cancel()"))
	} else {
		row.BuildTd().AppendItem(LinkTextProc("cmd", "Отменить", "media_cancel()"))
	}
	row.BuildTdClass("fill")
	return tbl
}

func (rule *DataRule) MediaReBrowser() {
	rule.StoreItem(rule.mediaBrowser(rule.ItPage.BrowserSize.Sx, rule.ItPage.BrowserSize.Sy))
}

func (rule *DataRule) mediaBrowser(sx int, sy int) likdom.Domer {
	rule.ItPage.BrowserSize = Size{sx, sy}
	tbl := likdom.BuildTable("width=100%", "id=win_browser")
	lev := len(lik.PathToNames(rule.ItPage.DirPath))
	tbl.BuildTrTd().AppendItem(rule.mediaPath(sx, lev*20))
	tbl.BuildTrTd().BuildString("<hr>")
	if !rule.ItPage.IsLoad {
		tbl.BuildTrTd().AppendItem(rule.mediaFiles(sx, sy-20-lev*20))
	} else {
	}
	return tbl
}

func (rule *DataRule) MediaReImage() {
	rule.StoreItem(rule.mediaImage(rule.ItPage.ImageSize.Sx, rule.ItPage.ImageSize.Sy))
}

func (rule *DataRule) mediaImage(sx int, sy int) likdom.Domer {
	rule.ItPage.ImageSize = Size{sx, sy}
	div := likdom.BuildDivClassId("win_visual", "win_preview")
	div.AppendItem(rule.VisualSource(sx - 8, sy - 8, dirMain + rule.ItPage.FilePath))
	return div
}

func (rule *DataRule) mediaLoad(sx int, sy int) likdom.Domer {
	tbl := likdom.BuildTableClass("")
	td := tbl.BuildTrTd()
	url := rule.BuildUrl("/front/media/upload?_mf=1")
	td.BuildItem("form", "class=dropzone", "id=mediaDropzone", "action", url)
	script := "var options = { addRemoveLinks: true };\n"
	script += "var myDropzone = new Dropzone(\"#mediaDropzone\", options);\n"
	td.BuildItem("script").BuildString("jQuery(function(){ " + script + " });")
	return tbl
}

func (rule *DataRule) mediaPath(sx int, sy int) likdom.Domer {
	tbl := likdom.BuildTable()
	names := lik.PathToNames(rule.ItPage.DirPath)
	ml := len(names)
	for nl := 0; nl <= ml; nl++ {
		row := tbl.BuildTr()
		for ns := 0; ns < nl; ns++ {
			row.BuildTd("width=24px").BuildString("&nbsp;")
		}
		name := "/"
		path := ""
		if nl > 0 {
			name = names[nl-1]
			path = "/" + strings.Join(names[:nl], "/")
		}
		link := LinkTextProc("cmd cmdd", name, fmt.Sprintf("media_path('%s')", path))
		row.BuildTd(fmt.Sprintf("width=%dpx", sx - nl * 24), fmt.Sprintf("colspan=%d", ml+1-nl)).AppendItem(link)
	}
	return tbl
}

func (rule *DataRule) mediaFiles(sx int, sy int) likdom.Domer {
	container := likdom.BuildDiv("style", fmt.Sprintf("height:%dpx; overflow-x:scroll; overflow-y:scroll;", sy))
	tbl := container.BuildTable("width=100%")
	if files, err := ioutil.ReadDir(dirMain + rule.ItPage.DirPath); err == nil {
		sort.Slice(files, func(i,j int) bool {
			return files[i].Name() >= files[j].Name()
		})
		for fase := 0; fase < 2; fase++ {
			for _, file := range files {
				if name := file.Name(); name != "" {
					full := rule.ItPage.DirPath + "/" + name
					if fase == 0 && file.IsDir() {
						link := LinkTextProc("cmd cmdd", "[" + name + "]", fmt.Sprintf("media_path('%s')", full))
						tbl.BuildTrTd("colspan=2").AppendItem(link)
					} else if fase == 1 && !file.IsDir() {
						row := tbl.BuildTr()
						row.BuildTd("width=24px").BuildString("&nbsp;")
						link := LinkTextProc("cmd cmdf", name, fmt.Sprintf("media_file('%s')", full))
						row.BuildTd(fmt.Sprintf("width=%dpx", sx-24)).AppendItem(link)
					}
				}
			}
		}
	}
	return container
}

func (rule *DataRule) ExecMedia() {
	if rule.IsShift("path") {
		rule.mediaDoPath()
	} else if rule.IsShift("file") {
		rule.mediaDoFile()
	} else if rule.IsShift("delete") {
		rule.mediaDoDelete()
	} else if rule.IsShift("load") {
		rule.mediaDoLoad()
	} else if rule.IsShift("upload") {
		rule.mediaDoUpload()
	} else if rule.IsShift("store") {
		rule.mediaDoStore()
	} else if rule.IsShift("cancel") {
		rule.mediaDoCancel()
	}
}

func (rule *DataRule) mediaDoPath() {
	rule.ItPage.DirPath = lik.StringFromXS(rule.Shift())
	rule.MediaReBrowser()
}

func (rule *DataRule) mediaDoFile() {
	rule.ItPage.FilePath = lik.StringFromXS(rule.Shift())
	rule.MediaReImage()
}

func (rule *DataRule) mediaDoDelete() {
}

func (rule *DataRule) mediaDoLoad() {
	rule.ItPage.IsLoad = true
	rule.ItPage.Upload = []Pot{}
	rule.MediaReShow()
}

func (rule *DataRule) mediaDoUpload() {
	if buffers := rule.GetBuffers(); buffers != nil {
		rule.ItPage.LoadSync.Lock()
		for name, val := range buffers {
			pot := Pot{ Name: name, Data: val }
			rule.ItPage.Upload = append(rule.ItPage.Upload, pot)
		}
		rule.ItPage.LoadSync.Unlock()
		rule.ItPage.NeedImage = true
	}
}

func (rule *DataRule) mediaDoStore() {
	for _,pot := range rule.ItPage.Upload {
		file := lik.Transliterate(pot.Name)
		path := dirMain + "/" + file
		_ = ioutil.WriteFile(path, pot.Data, 0666)
	}
	rule.ItPage.IsLoad = false
	rule.ItPage.Upload = []Pot{}
	rule.MediaReShow()
}

func (rule *DataRule) mediaDoCancel() {
	rule.ItPage.IsLoad = false
	rule.ItPage.Upload = []Pot{}
	rule.MediaReShow()
}

