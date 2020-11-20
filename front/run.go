package front

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likapi"
	"github.com/massarakhsh/lik/liktable"
	"sync"
)

const (
	Version = "0.2"
)

type DataSession struct {
	likapi.DataSession
	Sync		sync.Mutex
	Params		lik.Seter
}

type DataPage struct {
	likapi.DataPage
	Session    *DataSession
	Canal      string
	Variant    int
	ToPath     string
	NeedDraw   bool
	NeedUrl    bool
	NeedResize bool
	IsControl  bool
	PathAdmin  []string
	ListCanals *liktable.Table
	IdCanal    lik.IDB
	ASX, ASY   int
	ACX, ACY   int
	AMX, AMY   int
	APX, APY   int
	DirPath		string
	FilePath	string
}

type DataPager interface {
	likapi.DataPager
	GetItPage()	*DataPage
}

type DataRule struct {
	likapi.DataDrive
	ItPage	*DataPage
	ItSession *DataSession
	ResultFormat bool
	MediaTotal	int
	MediaFirst	int
	MediaCount	int
}

type DataRuler interface {
	likapi.DataDriver
	Page() *DataPage
}

func StartPage(uri string) *DataPage {
	session := &DataSession{}
	session.Uri = uri
	session.Params = lik.BuildSet()
	page := &DataPage{ Session: session }
	page.Self = page
	session.StartToPage(page)
	return page
}

func ClonePage(from *DataPage) *DataPage {
	page := &DataPage{ Session: from.Session }
	page.Self = page
	page.Canal = from.Canal
	page.ToPath = from.ToPath
	from.ContinueToPage(page)
	return page
}

func BindRule(page *DataPage) *DataRule {
	rule := &DataRule{ ItPage: page, ItSession: page.Session }
	rule.Page = page
	return rule
}

func (page *DataPage) GetItPage() *DataPage {
	return page
}

