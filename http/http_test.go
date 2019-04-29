package http

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "testing"
  // "strings"
)

const H_SERVER = "http://47.92.156.69:2611/hellogo"

func httpGet() {
  resp, err := http.Get("http://47.92.156.69:2611/hellogo")
  if err != nil {
    fmt.Println("catch err", err)
    return
  }

  defer resp.Body.Close()
  body, err :=  ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
}

func clientHttp() {
  req, err := http.NewRequest("GET", H_SERVER, nil)



}


func TestHttp(t *testing.T) {
  httpGet()

}
