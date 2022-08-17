package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

type arbitraryMapRecord struct {
	m map[string]interface{}
	s string
	d int
}

type record struct {
	name   string
	ofType interface{}
	fields []interface{}
	depth  int
}

func main() {
	var msg arbitraryMapRecord
	msg.m = exampleMap
	msg.d = 0
	schema, _ := arbitraryMapToParquetSchema(msg)
	fmt.Printf("Map to Parquet schema YAML:\n%v\n\n", schema)

	jsonFile, err := os.Open("./avro_schema.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var avroSchema arbitraryMapRecord
	var m map[string]interface{}
	json.Unmarshal([]byte(byteValue), &m)
	avroSchema.m = m
	avroSchema.d = 0
	schema, err = AvroSchemaToParquetSchema(avroSchema)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Avro schema to Parquet schema YAML:\n%v", schema)
}

func AvroSchemaToParquetSchema(msg arbitraryMapRecord) (string, error) {
	var node record
	node.name = msg.m["name"].(string)
	node.ofType = msg.m["type"]
	if _, f := msg.m["fields"]; f {
		for _, field := range msg.m["fields"].([]interface{}) {
			node.fields = append(node.fields, field.(map[string]interface{}))
		}
	}
	node.depth = msg.d
	// Assuming that Avro schema comes from OCF or Schema Registry and that actual fields are in fields[] of top-level object
	msg.s = iterateFields(node.fields, 0)
	return msg.s, nil
}

func traverseNodes(node record) string {
	switch node.ofType.(type) {
	case string:
		if len(node.fields) == 0 {
			return yamlIndent("", node.depth) + fmt.Sprintf(`{ "name":"%v","type":"%v" }`+"\n", node.name, AvroToParquetType(node.ofType.(string)))
		} else {
			if node.ofType.(string) == "record" {
				var n record
				n.name = node.name
				n.ofType = node.ofType
				n.depth = node.depth + 1
				if len(node.fields) > 0 {
					n.fields = append(n.fields, node.fields...)
				}
				return yamlIndent("", node.depth) + "name: " + node.name + "\n" + yamlIndentObj("", node.depth+1) + "optional: true\n" + yamlIndentObj("", node.depth+1) + "fields:\n" + iterateFields(n.fields, n.depth)
			}
		}
	case map[string]interface{}:
		var n record
		n.name = node.name
		n.ofType = node.ofType.(map[string]interface{})["type"]
		n.depth = node.depth + 1
		// Avro "array" field type
		if i, ok := node.ofType.(map[string]interface{})["items"]; ok {
			return yamlIndent("", node.depth) + fmt.Sprintf(`{ "name":"%v","type":"%v", repeated: true }`+"\n", node.name, AvroToParquetType(i.(string)))
		}
		// Avro "map" field type
		if i, ok := node.ofType.(map[string]interface{})["values"]; ok {
			return yamlIndent("", node.depth) + fmt.Sprintf(`{ "name":"%v","type":"MAP","valuetype":"%v" }`+"\n", node.name, AvroToParquetType(i.(string)))
		}
		// Avro "record" field type
		if _, f := node.ofType.(map[string]interface{})["fields"]; f {
			for _, field := range node.ofType.(map[string]interface{})["fields"].([]interface{}) {
				n.fields = append(n.fields, field.(map[string]interface{}))
			}
		}
		s := iterateFields(n.fields, n.depth)
		return yamlIndent("", node.depth) + "name: " + node.name + "\n" + yamlIndentObj("", node.depth+1) + "optional: true\n" + yamlIndentObj("", node.depth+1) + "fields:\n" + s
	// Avro union types
	case []interface{}:
		var unionTypes []string
		for _, ft := range node.ofType.([]interface{}) {
			switch ft.(type) {
			case string:
				if ft != "null" {
					unionTypes = append(unionTypes, ft.(string))
				}
			case []interface{}:
				var n record
				n.name = node.name
				n.ofType = node.ofType.(map[string]interface{})["type"]
				n.depth = node.depth + 1
				if _, f := node.ofType.(map[string]interface{})["fields"]; f {
					for _, field := range node.ofType.(map[string]interface{})["fields"].([]interface{}) {
						node.fields = append(node.fields, field.(map[string]interface{}))
					}
				}
				return traverseNodes(n)
			case map[string]interface{}:
				var n record
				n.name = node.name
				n.ofType = ft.(map[string]interface{})["type"]
				n.depth = node.depth + 1
				if _, f := ft.(map[string]interface{})["fields"]; f {
					for _, field := range ft.(map[string]interface{})["fields"].([]interface{}) {
						n.fields = append(n.fields, field.(map[string]interface{}))
					}
				}
				return yamlIndent("", node.depth) + "name: " + node.name + "\n" + yamlIndentObj("", node.depth+1) + "optional: true\n" + yamlIndentObj("", node.depth+1) + "fields:\n" + iterateFields(n.fields, n.depth)
			}
		}
		// Avro union type is null + one other type
		if len(unionTypes) == 1 {
			return yamlIndent("", node.depth) + fmt.Sprintf(`{ "name":"%v","type":"%v" }`+"\n", node.name, AvroToParquetType(unionTypes[0]))
		} else {
			// BYTE_ARRAY is the catchall if union type is anything beyond null + one other type
			return yamlIndent("", node.depth) + fmt.Sprintf(`{ "name":"%v","type":"BYTE_ARRAY" }`+"\n", node.name)
		}
	}
	return ""
}

func iterateFields(f []interface{}, depth int) string {
	var s string
	var n record
	for _, field := range f {
		n.name = field.(map[string]interface{})["name"].(string)
		n.ofType = field.(map[string]interface{})["type"]
		n.depth = depth

		if nf, f := field.(map[string]interface{})["fields"]; f {
			switch nf.(type) {
			case map[string]interface{}:
				for _, v := range nf.(map[string]interface{})["fields"].([]interface{}) {
					n.fields = append(n.fields, v.(map[string]interface{}))
				}
			default:
				for _, v := range nf.([]interface{}) {
					n.fields = append(n.fields, v.(map[string]interface{}))
				}
			}
		}
		s = s + traverseNodes(n)
	}
	return s
}

func yamlIndent(s string, depth int) string {
	in := "  "
	dent := "- "

	if depth == 0 {
		s = s + dent
		return s
	}

	for i := 0; i <= (depth - 1); i++ {
		s = s + in
	}

	s = s + dent
	return s
}

func yamlIndentObj(s string, d int) string {
	in := "  "

	for i := 0; i <= (d - 1); i++ {
		s = s + in
	}

	return s
}

func arbitraryMapToParquetSchema(msg arbitraryMapRecord) (string, error) {
	for k, v := range msg.m {
		switch t := v.(type) {
		case map[string]interface{}:
			msg.s = yamlIndent(msg.s, msg.d)
			msg.s = msg.s + `name: ` + k + "\n"
			msg.s = yamlIndentObj(msg.s, msg.d+1) + "optional: true\n"
			msg.s = yamlIndentObj(msg.s, msg.d+1) + "fields:\n"
			var n arbitraryMapRecord
			n.m = v.(map[string]interface{})
			n.d = msg.d + 1
			x, _ := arbitraryMapToParquetSchema(n)
			msg.s = msg.s + x
		case []interface{}:
			if len(v.([]interface{})) > 0 {
				switch f := reflect.TypeOf(v.([]interface{})[0]); f.String() {
				case "map":
					msg.s = yamlIndent(msg.s, msg.d)
					msg.s = msg.s + `name: ` + k + "\n"
					msg.s = yamlIndentObj(msg.s, msg.d+1) + "optional: true\n"
					msg.s = yamlIndentObj(msg.s, msg.d+1) + "fields:\n"
					var n arbitraryMapRecord
					n.m = v.(map[string]interface{})
					n.d = msg.d + 1
					x, _ := arbitraryMapToParquetSchema(n)
					msg.s = msg.s + x
				default:
					msg.s = yamlIndent(msg.s, msg.d)
					goType := fmt.Sprintf("%T", f)
					msg.s = msg.s + fmt.Sprintf(`{ "name":"%v","type":"%v", repeated: true }`+"\n", k, GoToParquetType(goType))
				}
			}
		case int, int32, int64, float32, float64, bool, string, nil:
			msg.s = yamlIndent(msg.s, msg.d)
			goType := fmt.Sprintf("%T", t)
			msg.s = msg.s + fmt.Sprintf(`{ "name":"%v","type":"%v" }`+"\n", k, GoToParquetType(goType))
		default:
			msg.s = yamlIndent(msg.s, msg.d)
			msg.s = msg.s + fmt.Sprintf(`{"name":"%v","type":"%v"}`+"\n", k, v)
		}
	}
	return msg.s, nil
}
