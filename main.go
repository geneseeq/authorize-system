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
