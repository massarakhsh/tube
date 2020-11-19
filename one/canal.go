package one

import (
	"github.com/jinzhu/gorm"
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likbase"
)

const KeyCanal = "canal"

type Canal struct {
	likbase.One        //	Общий объект
	Name        string //	Наименование
	Code        string //	Код
	Variant     int    //	Вариант
	Generate    int    //	Генерация
	Format    	string	 //	Формат
	Source0   	string   //	ID источника
	Source1   	string   //	ID источника
	Source2   	string   //	ID источника
	Source3   	string   //	ID источника
}

//	Инициализация таблицы
func InitializeCanal() {
	if !likbase.ODB.HasTable(KeyCanal) {
		DBCanal().CreateTable(&Canal{})
	} else {
		DBCanal().AutoMigrate(&Canal{})
	}
}

//	Позиционирование интерфейса
func DBCanal() *gorm.DB {
	return likbase.ODB.Table(KeyCanal)
}

//	Получить объект
func GetCanal(id lik.IDB) (Canal,bool) {
	it := Canal{}
	ok := likbase.Read(id, &it)
	return it,ok
}

//	Новый объект
func NewCanal(datas... interface{}) (Canal,bool) {
	it := Canal{}
	ok := likbase.Update(&it, datas...)
	return it,ok
}

//	Выбрать объекты
func SelectCanal(query interface{}, args... interface{}) []Canal {
	var canals []Canal
	if query != nil {
		DBCanal().Where(query, args...).Find(&canals)
	} else {
		DBCanal().Find(&canals)
	}
	return canals
}

//	Получить таблицу
func (it *Canal) Table() string {
	return KeyCanal
}

//	Создать
func (it *Canal) Create(datas... interface{}) bool {
	it.Id = 0
	return likbase.Update(it, datas...)
}

//	Изменить
func (it *Canal) Update(datas... interface{}) bool {
	return likbase.Update(it, datas...)
}

//	Удалить
func (it *Canal) Delete() {
	likbase.Delete(it)
}

func GetCanalName(name string, variant int) (Canal, bool) {
	var canals []Canal
	if DBCanal().Where("Code=? AND Variant=?", name, variant).Find(&canals); len(canals) > 0 {
		return canals[0], true
	}
	return Canal{}, false
}

