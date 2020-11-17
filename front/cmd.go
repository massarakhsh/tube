package front

func (rule *DataRule) CmdExec(zone string, cmd string) {
	if cmd == "admin" {
		rule.ItPage.IsControl = !rule.ItPage.IsControl
		rule.PageRedraw()
	}
}
