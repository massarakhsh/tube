package front

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likbase"
)

var (
	DB	likbase.JsonBaser
	TableParam  *likbase.ItTable
	TableCanal  *likbase.ItTable
	TableSource  *likbase.ItTable
	ListTables   []*likbase.ItTable
)

func GoIt(serv string, base string, user string, pass string) {
	DB = likbase.OpenJsonBase(serv, base, user, pass)
	TableParam = DB.BuildTable("param", "Параметры")
	TableCanal = DB.BuildTable("canal", "Каналы")
	TableSource = DB.BuildTable("source", "Источники")
	ListTables   = []*likbase.ItTable{
		TableParam,
		TableCanal,
		TableSource,
	}
	LoadListTables()
}

func GetTable(part string) *likbase.ItTable {
	for _, table := range ListTables {
		if table.Part == part {
			return table
		}
	}
	return nil
}

func GetElm(part string, id lik.IDB) *likbase.ItElm {
	var elm *likbase.ItElm
	if table := GetTable(part); table != nil {
		elm = table.GetElm(id)
	}
	return elm
}

func DeleteElm(part string, id lik.IDB) {
	if table := GetTable(part); table != nil {
		table.DeleteElm(id)
	}
}

func LoadListTables() {
	for _, table := range ListTables {
		table.LoadElms()
	}
}

func PurgeListTables() {
	for _, table := range ListTables {
		table.Purge()
	}
}

func FindCanal(name string) *likbase.ItElm {
	var canal *likbase.ItElm
	if name != "" {
		for _, elm := range TableCanal.Elms {
			if elm.GetString("idc") == name {
				canal = elm
				break
			}
		}
	}
	return canal
}

