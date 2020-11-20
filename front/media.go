package front

import (
	"fmt"
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"io/ioutil"
	"sort"
	"strings"
)

var dirMain = "./media"

func (rule *DataRule) MediaShow(sx int, sy int) likdom.Domer {
	rule.ItPage.AMX = sx
	rule.ItPage.AMY = sy
	code := likdom.BuildTableClassId("fill", "win_media")
	row := code.BuildTr()
	dx := sx /2 - BD
	row.BuildTd(MakeSizes(dx, sy)...).AppendItem(rule.mediaLeft(dx, sy))
	row.BuildTd(MakeSizes(dx, sy)...).AppendItem(rule.mediaRight(dx, sy))
	return code
}

func (rule *DataRule) MediaExecute(cmd string) {
}

func (rule *DataRule) mediaLeft(sx int, sy int) likdom.Domer {
	return MakeWindow("win_medias", sx, sy, "Файлы", rule.mediaControl(sx, sy - 32))
}

func (rule *DataRule) mediaRight(sx int, sy int) likdom.Domer {
	rule.ItPage.APX = sx
	rule.ItPage.APY = sy
	div := likdom.BuildDivClassId("", "win_preview")
	div.AppendItem(rule.VisualSource(sx, sy, dirMain + rule.ItPage.FilePath))
	return div
}

func (rule *DataRule) mediaControl(sx int, sy int) likdom.Domer {
	lev := len(lik.PathToNames(rule.ItPage.DirPath))
	tbl := likdom.BuildTable("width=100%")
	tbl.BuildTrTd().AppendItem(rule.mediaCommands(sx, 20))
	tbl.BuildTrTd("hr")
	tbl.BuildTrTd().AppendItem(rule.mediaPath(sx, 20 + lev*20))
	tbl.BuildTrTd("hr")
	tbl.BuildTrTd().AppendItem(rule.mediaFiles(sx, sy - 50 - lev*20))
	return tbl
}

func (rule *DataRule) mediaCommands(sx int, sy int) likdom.Domer {
	return likdom.BuildString("Команды")
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
		link := LinkTextProc("cmd cmdd", name, fmt.Sprintf("tube_path('%s')", path))
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
						link := LinkTextProc("cmd cmdd", "[" + name + "]", fmt.Sprintf("tube_direct('%s')", full))
						tbl.BuildTrTd("colspan=2").AppendItem(link)
					} else if fase == 1 && !file.IsDir() {
						row := tbl.BuildTr()
						row.BuildTd("width=24px").BuildString("&nbsp;")
						link := LinkTextProc("cmd cmdf", name, fmt.Sprintf("tube_file('%s')", full))
						row.BuildTd(fmt.Sprintf("width=%dpx", sx-24)).AppendItem(link)
					}
				}
			}
		}
	}
	return container
}

func (rule *DataRule) MediaDoPath(path string) {
	rule.ItPage.DirPath = path
	rule.StoreItem(rule.MediaShow(rule.ItPage.AMX, rule.ItPage.AMY))
}

func (rule *DataRule) MediaDoDirect(path string) {
	rule.ItPage.DirPath = path
	rule.StoreItem(rule.MediaShow(rule.ItPage.AMX, rule.ItPage.AMY))
}

func (rule *DataRule) MediaDoFile(path string) {
	rule.ItPage.FilePath = path
	rule.StoreItem(rule.mediaRight(rule.ItPage.APX, rule.ItPage.APY))
}
