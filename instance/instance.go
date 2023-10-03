package instance

import (
	"sync"

	"todo_graphql/config"
	"todo_graphql/logger"

	// for postgres
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type instance struct {
	db     orm.Ormer
}

var singleton = &instance{}
var once sync.Once

// Init initializes the instance
func Init() {
	postgresConfig := config.Postgres()
	once.Do(func() {

		logger.Log.Info("Connecting to postgres...", postgresConfig.ConnURL())
		
		orm.Debug = true
		
		err := orm.RegisterDriver("postgres", orm.DRPostgres)
		if err != nil {
			logger.Log.Fatal(err)
		}
		err = orm.RegisterDataBase("default", "postgres", postgresConfig.ConnURL())
		if err != nil {
			logger.Log.Fatal(err)
		}
		singleton.db = orm.NewOrm()
		singleton.db.Using("default")

		logger.Log.Info("Connected to postgres successfully...")
	})

}


// DB returns the database object
func DB() orm.Ormer {
	return singleton.db
}

// Destroy closes the connections & cleans up the instance
func Destroy() error {
	return nil
}