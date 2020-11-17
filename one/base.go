// Интерфейс базы данных
package one

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//	Дескриптор базы данных
var ODB *gorm.DB

//	Инициализация базы данных
func OpenBase(serv string, base string, user string, pass string) bool {
	args := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, serv, base)
	if ODB,_ = gorm.Open("mysql", args); ODB == nil {
		return false
	}
	InitializeFormat()
	InitializeCanal()
	InitializeSource()
	return true
}

