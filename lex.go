package main

import (
	"fmt"
	"reflect"
)

func tokenizeFile(file []byte) ([]Token, error) {
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
				var name, parameters, innards, returnType string
				for file[i] != byte(' ') && file[i] != byte('(') {
					name += string(file[i])
					i++
					fmt.Println("fc")
				}
				for file[i] == byte(' ') {
					i++
				}
				for file[i] != byte(' ') && file[i] != byte('(') {
					returnType += string(file[i])
					i++
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
				newToken.Values["returnType"] = returnType
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

	fmt.Println("file done")
	return out, nil
}

func tokenizeFunc(function string) (out []Token, err error) {

	for i := 0; i < len(function); i++ {

		// Var Handling
		if len(function) > i+4 {
			fmt.Println("vha")
			if reflect.DeepEqual(function[i:i+4], "var ") {
				fmt.Println("vb")
				newToken := Token{Type: 1, Values: map[string]string{}}
				i += 4
				var name, varType, value string
				for function[i] == byte(' ') {
					fmt.Println("vhc")
					i++
				}
				for function[i] != byte(' ') && function[i] != byte(':') && function[i] != byte('=') {
					name += string(function[i])
					i++
					fmt.Println("vhd")
				}
				if function[i] == byte(':') {
					i++
					fmt.Println("vhe")
					for function[i] == byte(' ') {
						i++
					}
					for function[i] != byte(' ') && function[i] != byte('=') {
						fmt.Println("vhf")
						varType += string(function[i])
						i++
					}
				} else {
					varType = "any"
				}
				for function[i] != byte('=') {
					i++
					fmt.Println("vhg")
				}
				i++
				for function[i] == byte(' ') {
					i++
					fmt.Println("vhh")
				}
				for i < len(function) {
					if function[i] != byte(' ') {
						value += string(function[i])
						i++
						fmt.Println("vhi")
					}
				}
				newToken.Values["name"] = name
				newToken.Values["type"] = varType
				newToken.Values["value"] = value
				out = append(out, newToken)
			}
		}

		// Func Call Handling
		if function[i] == byte('(') {
			newToken := Token{Type: 3, Values: map[string]string{}}
			var funcName, params string
			i--
			for function[i] == byte(' ') {
				i--
			}
			for function[i] != byte(' ') {
				funcName = string(function[i]) + funcName
				i--
			}
			for function[i] != byte('(') {
				i++
			}
			for function[i] != byte(')') {
				params += string(function[i])
				i++
			}
			params += ")"
			i++
			newToken.Values["name"] = funcName
			newToken.Values["params"] = params
			out = append(out, newToken)
		}
	}

	return
}
