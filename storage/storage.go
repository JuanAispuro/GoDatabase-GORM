package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// crearemos el singleton
// Variables Globales
var (
	db   *gorm.DB //Tipo puntero
	once sync.Once
)

// Driver de storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// -------------------------------- DAO --------------------------------
// Crea la conexión con la base de datos.
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// -------------------------------- PostgresDB --------------------------------

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", "postgres://postgres:Destructor11@localhost:5432/Godb?sslmode=disable")
		if err != nil {
			log.Fatalf("No pudimos abrir la base de datos: %v", err)
		}
		fmt.Println("Conectado a la base de datos postgres")
	}) //Lo que este aqui adentro solo se ejecutara una vez
}

// -------------------------------- MySQL --------------------------------
func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", "root:Destructor11@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("No pudimos abrir la base de datos: %v", err)
		}
		fmt.Println("Conectado a la MySQL")
	}) //Lo que este aqui adentro solo se ejecutara una vez
}

// Pool retorna una unica instancia de db
func Pool() *gorm.DB {
	return db
}
