package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)
	var results []interface{}
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(in, &n)
		var jsonBytes []byte
		for j := 0; j < n; j++ {
			line, _ := in.ReadBytes('\n')
			jsonBytes = append(jsonBytes, line...)
		}
		var data interface{}
		json.Unmarshal(jsonBytes, &data)
		results = append(results, prettify(data))
	}
	prettyJson, _ := json.Marshal(results)
	fmt.Fprintln(out, string(prettyJson))
}

func prettify(data interface{}) interface{} {
	switch v := data.(type) {
	case []interface{}:
		nonEmptyItems := make([]interface{}, 0)
		for _, item := range v {
			prettyItem := prettify(item)
			if prettyItem != nil {
				nonEmptyItems = append(nonEmptyItems, prettyItem)
			}
		}
		if len(nonEmptyItems) == 0 {
			return nil
		}
		return nonEmptyItems
	case map[string]interface{}:
		nonEmptyDict := make(map[string]interface{})
		for key, item := range v {
			prettyItem := prettify(item)
			if prettyItem != nil {
				nonEmptyDict[key] = prettyItem
			}
		}
		if len(nonEmptyDict) == 0 {
			return nil
		}
		return nonEmptyDict
	default:
		return v
	}
}
