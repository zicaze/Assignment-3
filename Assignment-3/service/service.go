package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

type Angka struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Alert struct {
	Status Angka `json:"status"`
}

func GetService() Alert {
	jsonFile, err := os.Open("file.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	fileBytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	var data Alert

	err = json.Unmarshal(fileBytes, &data)

	if err != nil {
		panic(err)
	}

	return data
}

func Random(a, b int) int {
	return a + rand.Intn(b)
}

func Updatejs(st Alert) {

	const File = "./file.json"

	st.Status.Water = Random(1, 100)
	st.Status.Wind = Random(1, 100)

	disasterByte, err := json.Marshal(st)

	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(File, disasterByte, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", st)
}
