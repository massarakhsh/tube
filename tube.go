package main

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likapi"
	"github.com/massarakhsh/tube/front"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	Port int = 80
	Serv string = "localhost"
	Base string = "tube"
	User string = "tube"
	Pass string = "tube"
)

func getArgs() bool {
	args, ok := lik.GetArgs(os.Args[1:])
	if val := args.GetInt("port"); val > 0 {
		Port = val
	}
	if val := args.GetString("serv"); val != "" {
		Serv = val
	}
	if val := args.GetString("base"); val != "" {
		Base = val
	}
	if val := args.GetString("user"); val != "" {
		User = val
	}
	if val := args.GetString("pass"); val != "" {
		Pass = val
	}
	if len(Base) <= 0 {
		fmt.Println("Base name must be present")
		ok = false
	}
	if !ok {
		fmt.Println("Usage: tsan [-key val | --key=val]...")
		fmt.Println("port    - port value (80)")
		fmt.Println("serv    - Database server")
		fmt.Println("base    - Database name")
		fmt.Println("user    - Database user")
		fmt.Println("pass    - Database pass")
	}
	return ok
}

func router(w http.ResponseWriter, r *http.Request) {
	if lik.RegExCompare(r.RequestURI, "\\.(js|css|htm|html|ico|gif|png|jpg|mp3|mp4|mpg|avi)") {
		likapi.ProbeRouteFile(w, r, r.RequestURI)
		return
	}
	var page *front.DataPage
	if sp := lik.StrToInt(likapi.GetParm(r, "_sp")); sp > 0 {
		if pager := likapi.FindPage(sp); pager != nil {
			page = pager.(front.DataPager).GetItPage()
		}
	}
	if page == nil {
		page = front.StartPage("")
	} else if lik.StrToInt(likapi.GetParm(r, "_tp")) > 0 {
		page = front.ClonePage(page)
		page.NeedUrl = true
	}
	rule := front.BindRule(page)
	rule.LoadRequest(r)
	if rule.IsShift("front") {
		rc, json := rule.BuildFront()
		likapi.RouteCookies(w, rule.GetAllCookies())
		likapi.RouteJson(w, rc, json, rule.ResultFormat)
	} else {
		rc, html := rule.PageHtml()
		likapi.RouteCookies(w, rule.GetAllCookies())
		likapi.RouteHtml(w, rc, html.ToString())
	}
}

func main() {
	//if host, _ := os.Hostname(); host == "Shaman" {
	//	Serv = "192.168.234.62"
	//}
	if !getArgs() {
		return
	}
	front.GoIt(Serv, Base, User, Pass)

	http.HandleFunc("/", router)
	if err := http.ListenAndServe(":"+fmt.Sprint(Port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
