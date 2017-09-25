// package main

// import (
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// type Person struct {
// 	Name  string
// 	Phone string
// }

// func main() {
// 	session, err := mgo.Dial("192.168.0.61:27015")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("GENESEEQ").C("SAMPLE_RECEIVE_ITEM")
// 	type TheGroup struct {
// 		Id         bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
// 		SampleCode string        `bson:"SAMPLE_CODE"`
// 	}
// 	var results []TheGroup
// 	var labeCodeList = []string{"P17030349223-YH006-28", "B17030349222"}
// 	// tempMap := map[string][]string{"$in": labeCodeList}
// 	tmpMatch := bson.M{"LAB_CODE": bson.M{"$in": labeCodeList}}
// 	tmpLookup := bson.M{"from": "DIC_SAMPLE_TYPE", "localField": "DIC_SAMPLE_TYPE", "foreignField": "ID", "as": "sample_advanced_info"}

// 	pipeline := []bson.M{
// 		bson.M{"$match": tmpMatch},
// 		bson.M{"$lookup": tmpLookup}}
// 	println(pipeline)
// 	pipe := c.Pipe(pipeline)
// 	pipe.All(&results)
// 	for _, data := range results {
// 		println(data.SampleCode)
// 	}

// 	// err := pipe.All(&results)
// 	// for iter.Next(&result) {
// 	// 	fmt.Printf("Result: %v\n", result)
// 	// }
// 	// if err := json.NewEncoder(w).Encode(results); err != nil {
// 	// 	panic(err)
// 	// }

// }
package main

import (
	"context"
	"fmt"
	"time"
)

func childFunc(cont context.Context, num *int) {
	ctx, _ := context.WithCancel(cont)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("child Done : ", ctx.Err())
			return
		}
	}
}

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("parent Done : ", ctx.Err())
					return // returning not to leak the goroutine
				case dst <- n:
					n++
					go childFunc(ctx, &n)
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	for n := range gen(ctx) {
		fmt.Println(n)
		if n >= 30000 {
			break
		}
	}
	cancel()
	time.Sleep(5 * time.Second)
}
