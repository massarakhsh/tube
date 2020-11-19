// Интерфейс базы данных
package one

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/massarakhsh/lik/likbase"
)

//	Инициализация базы данных
func OpenBase(serv string, base string, user string, pass string) bool {
	if !likbase.OpenBase(serv, base, user, pass) {
		return false
	}
	InitializeCanal()
	return true
}

