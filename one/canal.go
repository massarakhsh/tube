package one

import (
	"github.com/jinzhu/gorm"
	"github.com/massarakhsh/lik"
)

const KeyCanal = "canal"

type Canal struct {
	One						//	Общий объект
	Name        string		//	Наименование
	Code       	string		//	Код
	Variant    	int			//	Вариант
	Generate   	int			//	Генерация
	FormatId	int		`gorm:"index"`	//	ID формата
	Source0Id	int		`gorm:"index"`	//	ID источника
	Source1Id	int		`gorm:"index"`	//	ID источника
	Source2Id	int		`gorm:"index"`	//	ID источника
	Source3Id	int		`gorm:"index"`	//	ID источника
}

//	Инициализация таблицы
func InitializeCanal() {
	if !ODB.HasTable(KeyCanal) {
		DBCanal().CreateTable(&Canal{})
	} else {
		DBCanal().AutoMigrate(&Canal{})
	}
}

//	Позиционирование интерфейса
func DBCanal() *gorm.DB {
	return ODB.Table(KeyCanal)
}

//	Получить объект
func GetCanal(id lik.IDB) (Canal,bool) {
	it := Canal{}
	ok := it.read(KeyCanal, id, &it)
	return it,ok
}

//	Новый объект
func NewCanal(datas... interface{}) (Canal,bool) {
	it := Canal{}
	ok := it.create(KeyCanal, &it)
	if ok && len(datas) > 0 {
		ok = it.Update(datas...)
	}
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

//	Сохранить
func (it *Canal) Save() bool {
	return it.save(KeyCanal, it)
}

//	Изменить
func (it *Canal) Update(datas... interface{}) bool {
	return it.update(it, datas)
}

//	Удалить
func (it *Canal) Delete() {
	it.delete(KeyCanal, it)
}

func GetCanalName(name string, variant int) (Canal, bool) {
	var canals []Canal
	if DBCanal().Where("Code=? AND Variant=?", name, variant).Find(&canals); len(canals) > 0 {
		return canals[0], true
	}
	return Canal{}, false
}

