package main

import (
	"encoding/json"
	"reflect"

	"github.com/itchyny/gojq"
)

type Pool struct {
	PoolName    string   `json:"poolname"`
	Description string   `json:"description"`
	Members     []Member `json:"members"`
}

type Member struct {
	MemberName string `json:"membername"`
	Address    string `json:"address"`
	State      string `json:"state"`
}

// function to load json file and parse all pools
func loadJsonPools(firename string) ([]Pool, error) {
	var pools []Pool
	//read file at once
	body, err := ReadFile(filename)
	//process json
	query, err := gojq.Parse(`.list.form.fields[] | .data[]`)
	if err != nil {
		return nil, err
	}
	var input map[string]interface{}
	json.Unmarshal(body, &input)

	iter := query.Run(input) // or query.RunWithContext

	return pools, nil
}

func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func iter2slice2(iter gojq.Iter) ([]interface{}, error) {

	slice := make([]interface{}, 0)

	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return nil, err
		}

		if isNilFixed(v) {

		} else {

			slice = append(slice, v)
		}

		//fmt.Printf("%#v\n", v)
	}

	// ByServerName implements sort.Interface based on the ByServerName field.

	//    sort.Sort(slice,
	// 	func(i, j int) bool {
	// 			return string(slice[i]) < string(slice[j]);
	// 		  }
	// 	  )

	return slice, nil
}
