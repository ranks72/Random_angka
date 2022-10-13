package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status `json:"status"`
}

type MSG struct {
	Msg_water  string
	Data_water int
	Msg_wind   string
	Data_wind  int
}

var message = MSG{}

func updateData() {
	for {

		var data = Data{Status: Status{}}
		//dataMin := 1
		//dataMax := 30

		data.Status.Water = rand.Intn(30)

		data.Status.Wind = rand.Intn(30)

		b, err := json.MarshalIndent(&data, "", " ")

		if err != nil {
			log.Fatalln("error while marshalling json data  =>", err.Error())
		}

		err = ioutil.WriteFile("data.json", b, 0644)

		if err != nil {
			log.Fatalln("error while writing value to data.json file  =>", err.Error())
		}

		if data.Status.Water < 5 {
			message.Msg_water = "Status Aman"
			message.Data_water = data.Status.Water

		} else if data.Status.Water >= 6 && data.Status.Water <= 8 {
			//fmt.Println("air siaga")
			message.Msg_water = "Status Siaga"
			message.Data_water = data.Status.Water
		} else {
			//fmt.Println("air bahaya")
			message.Msg_water = "Status Bahaya"
			message.Data_water = data.Status.Water
		}

		if data.Status.Wind < 5 {
			//fmt.Println("angin aman")
			message.Msg_wind = "Status Aman"
			message.Data_wind = data.Status.Wind
		} else if data.Status.Wind >= 6 && data.Status.Wind <= 8 {
			//fmt.Println("angin siaga")
			message.Msg_wind = "Status Siaga"
			message.Data_wind = data.Status.Wind
		} else {
			//fmt.Println("angin bahaya")
			message.Msg_wind = "Status Bahaya"
			message.Data_wind = data.Status.Wind
		}

		fmt.Println("menggungu 5 detik")

		time.Sleep(time.Second * 5)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go updateData()

	http.HandleFunc("/tugas3", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("index.html")

		var data = Data{Status: Status{}}

		b, err := ioutil.ReadFile("data.json")

		if err != nil {
			fmt.Fprint(w, "error")
			return
		}

		err = json.Unmarshal(b, &data)

		err = tpl.ExecuteTemplate(w, "index.html", message)

	})

	http.ListenAndServe(":8080", nil)
}
