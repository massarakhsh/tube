package front

func (rule *DataRule) CmdExec(zone string, cmd string) {
	if cmd == "admin" {
		rule.ItPage.IsAdmin = !rule.ItPage.IsAdmin
		rule.PageRedraw()
	}
}
