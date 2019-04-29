package dbtest

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "time"
)

const (
    USERNAME = "root"
    PASSWORD = "mysqlroot23"
    NETWORK  = "tcp"
    SERVER   = "localhost"
    PORT     = 3306
    DATABASE = "hgl_biz"
)

type User struct {
  Id int64    `json:"user_id"`
  Name string `json:"user_name"`
  Age int     `json:"age"`
}

var Db *sql.DB

func ConnDb() {
  dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
  DB,err := sql.Open("mysql",dsn)
  if err != nil {
      fmt.Printf("Open mysql failed,err:%v\n",err)
      return
  }
  DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
  //DB.SetMaxOpenConns(100）//设置最大连接数
  DB.SetMaxIdleConns(16) //设置闲置连接数
  fmt.Println("conn db success")
  Db = DB
}

func InsertRecord(db *sql.DB, user User) (affectRow int64) {
  query :=  fmt.Sprintf("insert into users values (%d, '%s', %d)", user.Id, user.Name, user.Age)
  ret, err := db.Exec(query)
  if err != nil {
    fmt.Println("catch error when insert new record", err.Error())
    return
  }

  affectRow, _ = ret.RowsAffected()
  fmt.Println("affect row is ", affectRow)
  return
}
