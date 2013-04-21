package main

import (
  "code.google.com/p/goweb/goweb"
  "fmt"
  "io/ioutil"
)

func main() {
  pageRead, err := ioutil.ReadFile("osman.html")
  if err != nil {
    panic("Error reading file")
  }
  pageReadString := string(pageRead)

  goweb.MapFunc("/", func(c *goweb.Context) {
    fmt.Fprintf(c.ResponseWriter, pageReadString)
  })

  goweb.ListenAndServe(":80")
}
