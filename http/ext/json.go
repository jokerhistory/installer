package json

import (
	//缓存IO
	"encoding/json"
	"fmt" //io 工具包
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func Json2String() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	//							Output:
	//							{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}

}
