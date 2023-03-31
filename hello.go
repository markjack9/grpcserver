package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	trippb "server/proto/gen/go"
)

func main() {

	trip := trippb.Trip{
		Start:       "abc",
		End:         "edb",
		DurationSec: 3600,
		FeeCent:     1000,
	}
	fmt.Println(&trip)

	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", b)

	var trip2 trippb.Trip
	err = proto.Unmarshal(b, &trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)

	b, err = json.Marshal(&trip2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
