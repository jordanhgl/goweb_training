package model

import "testing"
import "fmt"

func TestStudy(t *testing.T) {
  var stu = Student{11, "hgl"}
  stu.Study("english")
  fmt.Println("hello san you say this")

  t.Log("test end")
}
