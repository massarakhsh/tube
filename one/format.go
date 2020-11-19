package one

import (
	"github.com/jinzhu/gorm"
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likbase"
)

const KeyFormat = "format"

type Format struct {
	likbase.One        //	Общий объект
	Name        string //	Наименование
	Code        string //	Код
}

//	Инициализация таблицы
func InitializeFormat() {
	if !likbase.ODB.HasTable(KeyFormat) {
		DBFormat().CreateTable(&Format{})
	} else {
		DBFormat().AutoMigrate(&Format{})
	}
}

//	Позиционирование интерфейса
func DBFormat() *gorm.DB {
	return likbase.ODB.Table(KeyFormat)
}

//	Получить объект
func GetFormat(id lik.IDB) (Format,bool) {
	it := Format{}
	ok := likbase.Read(id, &it)
	return it,ok
}

//	Новый объект
func NewFormat(datas... interface{}) (Format,bool) {
	it := Format{}
	ok := likbase.Update(&it, datas...)
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

//	Создать
func (it *Format) Create(datas... interface{}) bool {
	it.Id = 0
	return likbase.Update(it, datas)
}

//	Изменить
func (it *Format) Update(datas... interface{}) bool {
	return likbase.Update(it, datas)
}

//	Удалить
func (it *Format) Delete() {
	likbase.Delete(it)
}

