package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	map_data := map[string]interface{}{
		"name":"zhiliao",
		"website":[]map[string]interface{}{
			{
				"name":"知了官网",
				"url":"zlkt.net",
			},
			{
				"name":"知了官网",
				"url":"zhiliaoketang.net",
			},
		},
	}


	json_ret,err := json.MarshalIndent(map_data,""," ")
	if err != nil{
		return
	}
	fmt.Println(string(json_ret))     // 是切片得转string
}
