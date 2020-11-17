package one

import (
	"github.com/massarakhsh/lik"
	"github.com/jinzhu/gorm"
)

const KeyFormat = "format"

type Format struct {
	One						//	Общий объект
	Name        string		//	Наименование
	Code       	string		//	Код
}

//	Инициализация таблицы
func InitializeFormat() {
	if !ODB.HasTable(KeyFormat) {
		DBFormat().CreateTable(&Format{})
	} else {
		DBFormat().AutoMigrate(&Format{})
	}
}

//	Позиционирование интерфейса
func DBFormat() *gorm.DB {
	return ODB.Table(KeyFormat)
}

//	Получить объект
func GetFormat(id lik.IDB) (Format,bool) {
	it := Format{}
	ok := it.read(KeyFormat, id, &it)
	return it,ok
}

//	Новый объект
func NewFormat(datas... interface{}) (Format,bool) {
	it := Format{}
	ok := it.create(KeyFormat, it)
	if ok && len(datas) > 0 {
		ok = it.Update(datas...)
	}
	return it,ok
}

//	Выбрать объекты
func SelectFormat(query interface{}, args... interface{}) []Format {
	var formats []Format
	if query != nil {
		DBFormat().Where(query, args...).Find(&formats)
	} else {
		DBFormat().Find(&formats)
	}
	return formats
}

//	Получить таблицу
func (it *Format) Table() string {
	return KeyFormat
}

//	Сохранить
func (it *Format) Save() bool {
	return it.save(KeyFormat, it)
}

//	Изменить
func (it *Format) Update(datas... interface{}) bool {
	return it.update(it, datas)
}

//	Удалить
func (it *Format) Delete() {
	it.delete(KeyFormat, it)
}

