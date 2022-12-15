package main

import (
	"fmt"
	"reflect"
)

func tokenize(file []byte) ([]Token, error) {
	var out []Token

	for i := 0; i < len(file); i++ {
		// Main && Func Handling
		if len(file) > i+2 {
			fmt.Println("fa")
			if reflect.DeepEqual(string(file[i:i+2]), "fn") {
				fmt.Println("fb")
				newToken := Token{Type: 2, Values: map[string]string{}}
				i += 2
				for file[i] == byte(' ') {
					i++
				}
				var name, parameters, innards = "", "", ""
				for file[i] != byte(' ') && file[i] != byte('(') {
					name += string(file[i])
					i++
					fmt.Println("fc")
				}
				for file[i] != byte('(') {
					i++
					fmt.Println("fd")
				}
				for file[i] != byte(')') {
					parameters += string(file[i])
					i++
					fmt.Println("fe")
				}
				parameters += string(file[i])
				i++

				for file[i] != byte('{') {
					i++
					fmt.Println("ff")
				}
				for file[i] != byte('}') {
					innards += string(file[i])
					i++
					fmt.Println("fg")
				}
				innards += string(file[i])
				i++
				newToken.Values["params"] = parameters
				newToken.Values["name"] = name
				newToken.Values["innards"] = innards
				if name == "main" {
					newToken.Type = 9
				}
				out = append(out, newToken)
			}
		}

		// Var Handling
		if len(file) > i+4 {
			fmt.Println("va")
			if reflect.DeepEqual(string(file[i:i+4]), "var ") {
				fmt.Println("vb")
				newToken := Token{Type: 1, Values: map[string]string{}}
				i += 4
				var name, varType, value string
				for file[i] == byte(' ') {
					fmt.Println("vc")
					i++
				}
				for file[i] != byte(' ') && file[i] != byte(':') && file[i] != byte('=') {
					name += string(file[i])
					i++
					fmt.Println("vd")
				}
				if file[i] == byte(':') {
					i++
					fmt.Println("ve")
					for file[i] == byte(' ') {
						i++
					}
					for file[i] != byte(' ') && file[i] != byte('=') {
						fmt.Println("vf")
						varType += string(file[i])
						i++
					}
				} else {
					varType = "any"
				}
				for file[i] != byte('=') {
					i++
					fmt.Println("vg")
				}
				i++
				for file[i] == byte(' ') {
					i++
					fmt.Println("vh")
				}
				for i < len(file) {
					if file[i] != byte(' ') {
						value += string(file[i])
						i++
						fmt.Println("vi")
					}
				}
				newToken.Values["name"] = name
				newToken.Values["type"] = varType
				newToken.Values["value"] = value
				out = append(out, newToken)
			}
		}
	}

	return out, nil
}
