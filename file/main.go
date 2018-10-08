package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "os"
)

func main() {
    bs, err := ioutil.ReadFile("test.txt")
    if err != nil {
        return
    }
    str := string(bs)
    fmt.Println(str)
    create(str)
    ls("./")
    walk("/")
}
func create(s string) {
    file, err := os.Create("test1.txt")
    if err != nil {
        // здесь перехватывается ошибка
        return
    }
    defer file.Close()

    file.WriteString(s)
}
func ls(s string){
    dir, err:=os.Open(s)
    if err != nil {
        return
    }
    defer dir.Close()

    fileInfos, err := dir.Readdir(-1)
    if err != nil {
        return
    }
    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }
}
func walk(s string){
    filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
        fmt.Println(path)
        return nil
    })  
}