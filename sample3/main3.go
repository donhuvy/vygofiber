package sample3

import (
	foo "encoding/json"
	"fmt"
	"time"
)

type myJSON struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"dateValue"`
	ObjectValue     *myObject `json:"objectValue"`
	NullStringValue *string   `json:"nullStringValue"`
	NullIntValue    *int      `json:"nullIntValue"`
}

type myObject struct {
	ArrayValue []int `json:"arrayValue"`
}

func Main3() {
	otherInt := 4321
	data := &myJSON{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		ObjectValue: &myObject{
			ArrayValue: []int{1, 2, 3, 4},
		},
		NullStringValue: nil,
		NullIntValue:    &otherInt,
	}
	fmt.Println(foo.Marshal(data))
	fmt.Println(data)

	type myInt struct {
		IntValue int
	}

	data2 := &myInt{IntValue: 1234}
	fmt.Println(foo.Marshal(data2))

}
