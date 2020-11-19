package front

type Format struct {
	Code	string
	Name	string
}

var listFormats = []Format {
	{ "1", "Одно окно"},
	{ "4", "Четыре окна"},
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