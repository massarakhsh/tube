package front

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
)

func (rule *DataRule) CanalGen(sx int, sy int) likdom.Domer {
	return rule.CanalName(sx, sy, rule.ItPage.Canal)
}

func (rule *DataRule) CanalName(sx int, sy int, name string) likdom.Domer {
	var code likdom.Domer
	if name == "" {
	} else if canal := FindCanal(name); canal != nil {
		code = rule.CanalInfo(sx, sy, canal.Info)
	} else {
		code = rule.CanalNone(name)
	}
	return code
}

func (rule *DataRule) CanalInfo(sx int, sy int, info lik.Seter) likdom.Domer {
	var code likdom.Domer
	if what := info.GetString("what"); what == "" {
	} else if match := lik.RegExParse(what, "^canal_(.+)"); match != nil {
		name := match[1]
		code = rule.CanalName(sx, sy, name)
	} else if match := lik.RegExParse(what, "^row_(\\d+)"); match != nil {
		nx := lik.StrToInt(match[1])
		code = rule.CanalInfoTable(sx, sy, info, nx, 1)
	} else if match := lik.RegExParse(what, "^column_(\\d+)"); match != nil {
		ny := lik.StrToInt(match[1])
		code = rule.CanalInfoTable(sx, sy, info, 1, ny)
	} else if match := lik.RegExParse(what, "^table_(\\d+)_(\\d+)"); match != nil {
		ny := lik.StrToInt(match[1])
		nx := lik.StrToInt(match[2])
		code = rule.CanalInfoTable(sx, sy, info, nx, ny)
	} else if what == "html" {
		code = rule.CanalInfoHtml(sx, sy, info)
	} else if what == "image" {
		code = rule.CanalInfoImage(sx, sy, info.GetString("path"))
	} else if what == "video" {
		code = rule.CanalInfoVideo(sx, sy, info.GetString("path"))
	} else if what == "album" {
		code = rule.CanalInfoAlbum(sx, sy, info.GetString("path"))
	} else if lik.RegExCompare(what,"/.*\\.(jpg|jpeg|png|tif|tiff|gif)$") {
		code = rule.CanalInfoImage(sx, sy, what)
	} else if lik.RegExCompare(what,"/.*\\.(mpg|mpeg|mp4|avi)$") {
		code = rule.CanalInfoVideo(sx, sy, what)
	} else {
		code = rule.CanalBadFormat(what)
	}
	return code
}

func (rule *DataRule) CanalNone(name string) likdom.Domer {
	return rule.CanalMessage(fmt.Sprintf("Канал \"<b>%s</b>\" не найден", name))
}

func (rule *DataRule) CanalBadFormat(what string) likdom.Domer {
	return rule.CanalMessage(fmt.Sprintf("Неизвестный формат \"<b>%s</b>\"", what))
}

func (rule *DataRule) CanalNoFile(file string) likdom.Domer {
	return rule.CanalMessage(fmt.Sprintf("Файл \"<b>%s</b>\" не найден", file))
}

func (rule *DataRule) CanalMessage(text string) likdom.Domer {
	code := likdom.BuildItemClass("table","fill")
	code.BuildTrTdClass("fill").BuildString(text)
	return code
}

func (rule *DataRule) CanalInfoTable(sx int, sy int, info lik.Seter, nx int, ny int) likdom.Domer {
	if nx <= 0 || ny <= 0 {
		return nil
	}
	code := likdom.BuildTable(MakeSizes(sx, sy)...)
	dx := sx / nx
	dy := sy / ny
	for y := 0; y < ny; y++ {
		tr := code.BuildTr()
		for x := 0; x < nx; x++ {
			td := tr.BuildTdClass("boxcell", MakeSizes(dx, dy)...)
			cell := info.GetItem(fmt.Sprintf("cell_%d_%d", y, x))
			if cell == nil && nx == 1 {
				cell = info.GetItem(fmt.Sprintf("cell_%d", y))
			}
			if cell == nil && ny == 1 {
				cell = info.GetItem(fmt.Sprintf("cell_%d", x))
			}
			if cell != nil {
				if cell.IsString() || cell.IsInt() {
					if cd := rule.CanalName(dx-BD, dy-BD, cell.ToString()); cd != nil {
						td.AppendItem(cd)
					}
				} else if set := cell.ToSet(); set != nil {
					set := cell.ToSet()
					if align := set.GetString("cls"); align != "" {
						td.SetAttr("class", td.GetAttr("class") + " " + align)
					}
					if cd := rule.CanalInfo(dx-BD, dy-BD, set); cd != nil {
						td.AppendItem(cd)
					}
				}
			}
		}
	}
	return code
}

func (rule *DataRule) CanalInfoHtml(sx int, sy int, info lik.Seter) likdom.Domer {
	code := info.GetString("html")
	return rule.CanalMessage(code)
}

func (rule *DataRule) CanalInfoImage(sx int, sy int, path string) likdom.Domer {
	div := likdom.BuildItem("div")
	id := fmt.Sprintf("id%d", rand.Int31n(1000000))
	code := div.BuildDiv("id", id, "style", fmt.Sprintf("width:%dpx;height:%dpx;", sx, sy))
	code.SetAttr(fmt.Sprintf("data-width=%d", sx), fmt.Sprintf("data-height=%d", sy))
	code.SetAttr(fmt.Sprintf("data-maxwidth=%d", sx), fmt.Sprintf("data-maxheight=%d", sy))
	code.BuildUnpairItem("img", "src", path)
	init := fmt.Sprintf("jQuery('#%s').fotorama()", id)
	script := fmt.Sprintf("jQuery(document).ready(function() { %s });", init)
	div.BuildItem("script").BuildString(script)
	return div
}

func (rule *DataRule) CanalInfoVideo(sx int, sy int, path string) likdom.Domer {
	if _, err := os.Stat("."+ path); err != nil {
		return rule.CanalNoFile(path)
	}
	code := likdom.BuildItem("video","width=100%", "height=100%", "controls=yes", "autoplay=yes", "loop=yes", "muted=yes")
	if match := regexp.MustCompile("\\.(.+)$").FindStringSubmatch(path); match != nil {
		format := match[1]
		code.BuildUnpairItem("source", "src", path, "type", "video/"+format)
	}
	return code
}

func (rule *DataRule) CanalInfoAlbum(sx int, sy int, path string) likdom.Domer {
	files, err := ioutil.ReadDir("."+path)
	if err != nil {
		return rule.CanalNoFile(path)
	}
	div := likdom.BuildItem("div")
	id := fmt.Sprintf("id%d", rand.Int31n(1000000))
	code := div.BuildDiv("id", id, "style", fmt.Sprintf("width:%dpx;height:%dpx;", sx, sy))
	code.SetAttr(fmt.Sprintf("data-width=%d", sx), fmt.Sprintf("data-height=%d", sy-96))
	code.SetAttr(fmt.Sprintf("data-maxwidth=%d", sx), fmt.Sprintf("data-maxheight=%d", sy-96))
	code.SetAttr("data-hash=true", "data-nav=thumbs", "data-loop=true", "data-autoplay=5000")
	for _, file := range files {
		if match := regexp.MustCompile("([^/]+)$").FindStringSubmatch(file.Name()); match != nil {
			code.BuildUnpairItem("img", "src", path + "/" + match[1])
		}
	}
	init := fmt.Sprintf("jQuery('#%s').fotorama()", id)
	script := fmt.Sprintf("jQuery(document).ready(function() { %s });", init)
	div.BuildItem("script").BuildString(script)
	return div
}

