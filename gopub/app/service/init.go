package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gopub/app/entity"
	"os"
	"fmt"
)

var (
	o 	orm.Ormer
)

func Init(){
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	tablePrefix := beego.AppConfig.String("db.prefix")
	if dbPort == ""{
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	//if timezone != "" {
	//	dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	//}
	orm.RegisterDataBase("default","mysql",dsn)
	orm.RegisterModelWithPrefix(tablePrefix,
		new(entity.Action),
		new(entity.Env),
		new(entity.EnvServer),
		//new(entity.MailTpl),
		//new(entity.Perm),
		new(entity.Project),
		new(entity.Role),
		//new(entity.RolePerm),
		new(entity.Server),
		new(entity.Task),
		new(entity.TaskReview),
		new(entity.User),
		//new(entity.UserRole),
	)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm()
	orm.RunCommand()

	os.Mkdir(GetProjectsBasePath(),0755)
	os.Mkdir(GetTasksBasePath(),0755)



}

// 所有项目根目录
func GetProjectsBasePath() string {
	return fmt.Sprintf(beego.AppConfig.String("data_dir") + "/projects")
}

// 任务单根目录
func GetTasksBasePath() string {
	return fmt.Sprintf(beego.AppConfig.String("data_dir") + "/tasks")
}