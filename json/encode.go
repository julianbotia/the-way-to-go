package json

import (
	"encoding/json"
	"fmt"
)

// Message struct
type Message struct {
	Name string
	Body string
	Time int64
}

func encodeMessage() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b)) // {"Name":"Alice","Body":"Hello","Time":1294706395881547000}
}

func decodeMessage() {
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	var m Message
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m.Name) // Alice
	fmt.Println(m.Body) // Hello
	fmt.Println(m.Time) // 1294706395881547000
}

// Data struct
type Data struct {
	Name     string `json:"myName"`               // Field appears in JSON as key "myName".
	LastName string `json:"mylastname,omitempty"` // the field is omitted from the object if its value is empty,
	Age      int    `json:",omitempty"`           // the field is skipped if empty.
	Address  string `json:"-"`                    // Field is ignored by this package.
	Status   bool   `json:"-,"`                   // Field appears in JSON as key "-"
}

func encodeDatav1() {
	var d Data
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b)) //{"myName":"","-":false}
}
func encodeDatav2() {
	d := Data{
		Name:     "Juan",
		LastName: "Perez",
		Age:      50,
		Address:  "Calle 1",
		Status:   true,
	}
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b)) // {"myName":"Juan","mylastname":"Perez","Age":50,"-":true}
}
func decodeMessagev2() {
	b := []byte(`{"myName":"Juan","mylastname": "Perez","Age":20, "Address":"Calle ...","-": true, "unknownField": "..."}`)
	var d Data
	err := json.Unmarshal(b, &d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Name: ", d.Name)         // Juan
	fmt.Println("LastName: ", d.LastName) // Perez
	fmt.Println("Age: ", d.Age)           // 10
	fmt.Println("Address: ", d.Address)   //
	fmt.Println("Status: ", d.Status)     // true
}

// Response ...
type Response struct {
	Data Data
	Time string
}

func encodeWithInnerNodes() {
	d := Data{
		Name:     "Juan",
		LastName: "Perez",
		Age:      50,
		Address:  "Calle 1",
		Status:   true,
	}
	r := Response{Data: d, Time: "Now :p"}
	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b)) // {"Data":{"myName":"Juan","mylastname":"Perez","Age":50,"-":true},"Time":"Now :p"}
}
