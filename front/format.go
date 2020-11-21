package front

type Format struct {
	Code	string
	Name	string
	Count	int
}

var listFormats = []Format {
	{ "1", "Одно окно", 1},
	{ "2", "Два окна", 2},
	{ "12", "Одно слева и два справа", 3},
	{ "21", "Два слева и одно справа", 3},
	{ "4", "Четыре окна", 4},
}

func GetListFormats() []Format {
	return listFormats
}

func GetFormat(code string) *Format {
	for nf := 0; nf < len(listFormats); nf++ {
		if code == listFormats[nf].Code {
			return &listFormats[nf]
		}
	}
	return nil
}