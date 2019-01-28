package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	//omitempty 値がからの場合には出力しない設定
	Name     string   `json:"name,omitempty"` //空の場合
	Age      int     `json:"age,omitempty"` //0の場合
	Nicknames []string `json:"nicknames"` //空の場合
}

//json.MarshalはMarshalJSON()でカスタマイズすることができる
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct{
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}
//同様にjson.Unmarshalをカスタマイズ可能
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

func main() {
	b := []byte(`{"name":"mike", "age":20, "nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil{
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))

}