package model

import (
  "fmt"
)


type Student struct {
    No int64
    Name string
}


func (stu *Student) Study(subject string) {

  fmt.Printf("student:%d learning subject: %s", stu.No, subject)
  fmt.Println()
 
}
