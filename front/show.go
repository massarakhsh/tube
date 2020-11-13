package front

import (
	"github.com/massarakhsh/lik/likdom"
	"fmt"
)

func MakeSizes(sx int, sy int) []string {
	return []string{ fmt.Sprintf("width=%dpx", sx), fmt.Sprintf("height=%dpx", sy) }
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

