package albert

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	DriveName      string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func JsonFile() {
	a := [...]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Printf("切片的长度为%v，切片的容量为%v\n", len(b), cap(b))

	file, err := os.Open("F:/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	conf := new(Configuration)

	json.NewDecoder(file).Decode(conf)

	fmt.Println(conf.DataSourceName)
}
