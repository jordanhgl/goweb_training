package datatype
import (
  "fmt"
  "testing"
)


type MapString map[string]string


func (mapstr MapString) selfPrint() {
  for k, v := range mapstr {
    fmt.Printf("k = %s and v = %s", k, v)
    fmt.Println()
  }
}

func TestMap(t *testing.T) {
  var amap = make(map[string]int)
  amap["key1"] = 10
  fmt.Println("amap as follow:", amap)

  if val, ok := amap["key1"]; ok {
    fmt.Println("found val match key1", val)
  } else {
    fmt.Println("key does not exist!")
  }

  var testMap = make(MapString)
  testMap["aaa"] = "aaa123"
  testMap["bbb"] = "bbb123"

  testMap.selfPrint()


  t.Log(amap)
}
