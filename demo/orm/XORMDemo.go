package main

import (
 _ "github.com/go-sql-driver/mysql"
 "github.com/go-xorm/xorm"
	"fmt"
	"time"
	"justgo/util"
)

var engine *xorm.Engine

type Userinfo struct {
	Uid   int64 `xorm:"pk autoincr"`
	UserName string  `xorm:"varchar(64) notnull unique 'username'"`
	DepartName string  `xorm:"varchar(64) notnull 'departName'"`
	Created time.Time `xorm:"created"`
}

func main() {
	// init engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root@/godemo?charset=utf8")

	engine.ShowSQL(true)

	if err != nil {
		fmt.Println(err)
	}

	// insert
	user := new(Userinfo)
	user.UserName = util.RandString(5)
	user.DepartName = "IT"
	affected, err := engine.Insert(user)
	fmt.Println(affected, user.Uid)
	if err != nil {
		fmt.Println(err)
	}

	// query
	var queryUser Userinfo
	engine.Id(user.Uid).Get(&queryUser)

	fmt.Printf("%#v", queryUser)

	var users []Userinfo
	engine.Sql("select * from userinfo").Find(&users)

	for _, u := range users {
		fmt.Printf("%#v\n", u)
	}

	// delete
	engine.Delete(user)

}
