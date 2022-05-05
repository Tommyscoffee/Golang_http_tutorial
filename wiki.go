package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) { //複数の返り値、
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil { //エラー処理
		return nil, err //第一返り値はnilを返す
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")} //P構造体に値を代入
	p1.save()                                                              //ここでファイル作成
	p2, _ := loadPage("TestPage")                                          //ファイル内に構造体内のメンバ変数を入れている。
	fmt.Println(string(p2.Body))
}
