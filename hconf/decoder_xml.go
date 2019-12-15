package hconf

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
)

type XmlDecoder struct{}

const (
	DataElement = iota
	StartElement
	EndElement
)

type TokenInfo struct {
	Key string
	Mod int
}

func (d *XmlDecoder) Decode(buf []byte) (Storage, error) {
	decoder := xml.NewDecoder(bytes.NewReader(buf))

	var tokens []*TokenInfo
	t, err := decoder.Token()
	for ; err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			// 忽略开始元素前面的数据
			if len(tokens) > 0 && tokens[len(tokens)-1].Mod == DataElement {
				tokens = tokens[:len(tokens)-1]
			}
			tokens = append(tokens, &TokenInfo{
				Key: token.Name.Local,
				Mod: StartElement,
			})
		case xml.EndElement:
			tokens = append(tokens, &TokenInfo{
				Key: token.Name.Local,
				Mod: EndElement,
			})
		case xml.CharData:
			// 忽略没有跟在开始元素后面的数据
			if len(tokens) > 0 && tokens[len(tokens)-1].Mod != StartElement {
				continue
			}
			tokens = append(tokens, &TokenInfo{
				Key: string(token),
				Mod: DataElement,
			})
		}
	}
	if err != io.EOF {
		return nil, err
	}

	var keys []string
	var nodes []map[string]interface{}
	nodes = append(nodes, map[string]interface{}{})
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Mod {
		case StartElement:
			if i+1 == len(tokens) {
				return nil, fmt.Errorf("miss data element")
			}
			if tokens[i+1].Mod == DataElement {
				if i+2 == len(tokens) || tokens[i+2].Mod != EndElement || tokens[i+2].Key != tokens[i].Key {
					return nil, fmt.Errorf("element not match [%v]", tokens[i].Key)
				}
				node := nodes[len(nodes)-1]
				key := tokens[i].Key
				val := tokens[i+1].Key
				old, ok := node[key]
				if !ok {
					node[key] = val
				} else {
					if v, ok := old.([]interface{}); ok {
						node[key] = append(v, val)
					} else {
						node[key] = []interface{}{old, val}
					}
				}
				i += 2
			} else if tokens[i+1].Mod == EndElement {
				i++
			} else {
				node := nodes[len(nodes)-1]
				key := tokens[i].Key
				old, ok := node[key]
				val := map[string]interface{}{}
				if !ok {
					node[key] = val
				} else {
					if v, ok := old.([]interface{}); ok {
						node[key] = append(v, val)
					} else {
						node[key] = []interface{}{old, val}
					}
				}
				nodes = append(nodes, val)
				keys = append(keys, tokens[i].Key)
			}
		case EndElement:
			if len(keys) == 0 || keys[len(keys)-1] != tokens[i].Key {
				return nil, fmt.Errorf("end elements not match [%v]", tokens[i].Key)
			}
			node := nodes[len(nodes)-1]
			key := keys[len(keys)-1]

			// 如果对象只包含一个对象数组，将该对象改为数组
			if len(node) == 1 {
				for _, v := range node {
					if _, ok := v.([]interface{}); ok {
						nodes[len(nodes)-2][key] = v
					}
				}
			}

			nodes = nodes[:len(nodes)-1]
			keys = keys[:len(keys)-1]
		case DataElement:
			return nil, fmt.Errorf("parse error near [%v]", tokens[i].Key)
		}
	}

	if len(nodes) != 1 {
		return nil, fmt.Errorf("parse xml failed")
	}
	var data interface{}
	node := nodes[0]
	data = node
	if len(node) == 1 {
		for _, v := range node {
			fmt.Println(reflect.TypeOf(v))
			if _, ok := v.([]interface{}); ok {
				data = v
			}
		}
	}

	return NewInterfaceStorage(data), nil
}
