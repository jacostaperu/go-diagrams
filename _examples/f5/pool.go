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
func LoadJsonPools(filename string) ([]Pool, error) {
	var pools []Pool
	//read file at once
	body, err := ReadFile(filename)
	//process json

	//falta poner aqui la query que transforma en el formato requerido
	// revisar el getPool.sh para obtener la query
	stringQuery := `to_entries[] | .["pool"]=.key  |
        .["description"]=.value.description | .["members"]=[
            .value.members | to_entries[]|
             .["name"]=.key|
             .["address"]=.value.address|
             .["state"]=.value.state|
             del(.key,.value)
            ] |
        del(.key,.value)`
	query, err := gojq.Parse(stringQuery)
	if err != nil {
		return nil, err
	}
	var input map[string]interface{}
	json.Unmarshal(body, &input)

	iter := query.Run(input) // or query.RunWithContext

	jsonPools, err := iter2slice2(iter)
	if err != nil {
		return nil, err
	}

	//body, err = json.Marshal(nodes)
	var pools []Pool
	mbytes, err := json.Marshal(jsonPools)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(mbytes, &pools)

	return pool, err

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
