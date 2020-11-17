package one

import (
	"github.com/massarakhsh/lik"
	"github.com/jinzhu/gorm"
)

const KeySource = "source"

type Source struct {
	One						//	Общий объект
	Name        string		//	Наименование
	Proto      	string		//	Прототип
	Path      	string		//	Путь
}

//	Инициализация таблицы
func InitializeSource() {
	if !ODB.HasTable(KeySource) {
		DBSource().CreateTable(&Source{})
	} else {
		DBSource().AutoMigrate(&Source{})
	}
}

//	Позиционирование интерфейса
func DBSource() *gorm.DB {
	return ODB.Table(KeySource)
}

//	Получить объект
func GetSource(id lik.IDB) (Source,bool) {
	it := Source{}
	ok := it.read(KeySource, id, &it)
	return it,ok
}

//	Новый объект
func NewSource(datas... interface{}) (Source,bool) {
	it := Source{}
	ok := it.create(KeySource, it)
	if ok && len(datas) > 0 {
		ok = it.Update(datas...)
	}
	return it,ok
}

//	Выбрать объекты
func SelectSource(query interface{}, args... interface{}) []Source {
	var sources []Source
	if query != nil {
		DBSource().Where(query, args...).Find(&sources)
	} else {
		DBSource().Find(&sources)
	}
	return sources
}

//	Получить таблицу
func (it *Source) Table() string {
	return KeySource
}

//	Сохранить
func (it *Source) Save() bool {
	return it.save(KeySource, it)
}

//	Изменить
func (it *Source) Update(datas... interface{}) bool {
	return it.update(it, datas)
}

//	Удалить
func (it *Source) Delete() {
	it.delete(KeySource, it)
}

