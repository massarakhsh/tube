package front

import (
	"fmt"
	"github.com/massarakhsh/lik/likdom"
)

func MakeSizes(sx int, sy int) []string {
	var attrs []string
	if sx > 0 {
		attrs = append(attrs, fmt.Sprintf("width=%dpx", sx))
	}
	if sy > 0 {
		attrs = append(attrs, fmt.Sprintf("height=%dpx", sy))
	}
	return attrs
}

func MakeWindow(name string, sx int, sy int, title string, data likdom.Domer) likdom.Domer {
	code := likdom.BuildTable()
	cls := "fill"
	if name != "" {
		cls = name
		code.SetAttr("id", name)
		code.SetAttr("class", name)
	}
	if sx > 0 {
		code.SetAttr("width", fmt.Sprintf("%dpx", sx))
	} else if sx < 0 {
		code.SetAttr("width", fmt.Sprintf("%d%%", -sx))
	}
	if sy > 0 {
		code.SetAttr("height", fmt.Sprintf("%dpx", sy))
	} else if sy < 0 {
		code.SetAttr("height", fmt.Sprintf("%d%%", -sy))
	}
	if title != "" {
		code.BuildTrTdClass("boxtitle", MakeSizes(sx, 24)...).BuildString(title)
		if sy >= 24 {
			sy -= 24
		}
	}
	if data != nil {
		code.BuildTrTdClass(cls, MakeSizes(sx, sy)...).AppendItem(data)
	}
	return code
}

func LinkTextCmd(cls string, text string, cmd string) likdom.Domer {
	return LinkTextProc(cls, text,"front_get('"+cmd+"')")
}

func LinkTextProc(cls string, text string, proc string) likdom.Domer {
	div := likdom.BuildDivClassId(cls, "", "onclick", proc)
	div.BuildString(text)
	return div
}

func LinkPicCmd(cls string, pic string, cmd string, title string) likdom.Domer {
	return LinkPicProc(cls, pic,"front_get('"+cmd+"')", title)
}

func LinkPicProc(cls string, pic string, proc string, title string) likdom.Domer {
	img := likdom.BuildUnpairItem("img")
	if cls != "" { img.SetAttr("class", cls) }
	if pic != ""{ img.SetAttr("src", pic) }
	if proc != "" { img.SetAttr("onclick", proc) }
	if title != "" { img.SetAttr("title", title) }
	return img
}

func MakeNamelyCanal(name string, variant int) string {
	title := name
	if title == "" {
		title = "Без имени"
	}
	if variant > 0 {
		title += fmt.Sprintf(" (черновик-%d)", variant)
	}
	return title
}

