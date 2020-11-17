package front

import (
	"fmt"
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/tube/one"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

func (rule *DataRule) VisualGen(sx int, sy int) likdom.Domer {
	var code likdom.Domer
	if name := rule.ItPage.Canal; name == "" {
	} else if canal,ok := one.GetCanalName(name, rule.ItPage.Variant); ok {
		code = rule.VisualGenCanal(sx, sy, &canal)
	} else {
		code = rule.CanalNoCanal(name)
	}
	return MakeWindow("win_visual", sx, sy, "", code)
}

func (rule *DataRule) VisualGenCanal(sx int, sy int, canal *one.Canal) likdom.Domer {
	var code likdom.Domer
	if canal == nil {
	} else if format,ok := one.GetFormat(lik.IDB(canal.FormatId)); !ok {
		code = rule.VisualMessage(fmt.Sprintf("Неизвестный формат \"<b>ID=%d</b>\"", canal.FormatId))
	} else if format.Code == "1" {
		code = rule.VisualFormat1(sx, sy, canal)
	} else if format.Code == "4" {
		code = rule.VisualFormat4(sx, sy, canal)
	} else {
		code = rule.VisualMessage(fmt.Sprintf("Неизвестный формат \"<b>%s</b>\"", format.Code))
	}
	return code
}

func (rule *DataRule) VisualFormat1(sx int, sy int, canal *one.Canal) likdom.Domer {
	return rule.VisualWindow(sx, sy, lik.IDB(canal.Source0Id))
}

func (rule *DataRule) VisualFormat4(sx int, sy int, canal *one.Canal) likdom.Domer {
	dx := sx / 2 - BD
	dy := sy / 2 - BD
	code := likdom.BuildTableClass("fill")
	if row := code.BuildTr(); row != nil {
		row.BuildTd(MakeSizes(dx, dy)...).AppendItem(rule.VisualWindow(dx, dy, lik.IDB(canal.Source0Id)))
		row.BuildTd(MakeSizes(dx, dy)...).AppendItem(rule.VisualWindow(dx, dy, lik.IDB(canal.Source1Id)))
	}
	if row := code.BuildTr(); row != nil {
		row.BuildTd(MakeSizes(dx, dy)...).AppendItem(rule.VisualWindow(dx, dy, lik.IDB(canal.Source2Id)))
		row.BuildTd(MakeSizes(dx, dy)...).AppendItem(rule.VisualWindow(dx, dy, lik.IDB(canal.Source3Id)))
	}
	return code
}

func (rule *DataRule) VisualWindow(sx int, sy int, ids lik.IDB) likdom.Domer {
	var code likdom.Domer
	if ids == 0 {
	} else if source,ok := one.GetSource(ids); !ok {
		code = rule.VisualMessage(fmt.Sprintf("Неизвестный источник \"<b>ID=%d</b>\"", int(ids)))
	} else {
		code = rule.VisualSource(sx, sy, &source)
	}
	return code
}

func (rule *DataRule) VisualSource(sx int, sy int, source *one.Source) likdom.Domer {
	var code likdom.Domer
	if source != nil {
		if lik.RegExCompare(strings.ToLower(source.Path), "(jpg|png|gif)$") {
			code = rule.VisualImage(sx, sy, source.Path)
		} else if lik.RegExCompare(strings.ToLower(source.Path), "(avi|mpg|mts)$") {
			code = rule.VisualVideo(sx, sy, source.Path)
		}
	}
	return code
}

func (rule *DataRule) CanalNoCanal(name string) likdom.Domer {
	return rule.VisualMessage(fmt.Sprintf("Канал \"<b>%s</b>\" не найден", name))
}

func (rule *DataRule) CanalNoFile(file string) likdom.Domer {
	return rule.VisualMessage(fmt.Sprintf("Файл \"<b>%s</b>\" не найден", file))
}

func (rule *DataRule) VisualMessage(text string) likdom.Domer {
	code := likdom.BuildItemClass("table","fill")
	code.BuildTrTdClass("fill").BuildString(text)
	return code
}

func (rule *DataRule) VisualImage(sx int, sy int, path string) likdom.Domer {
	div := likdom.BuildDivClass("fill")
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

func (rule *DataRule) VisualVideo(sx int, sy int, path string) likdom.Domer {
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

