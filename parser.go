package main

func parse(tokens []Token) (string, []string) {
	var out string
	var imports []string

	for _, token := range tokens {
		switch token.Type {
		// Func Declaration Handler
		case 2:
			val, ok := typeTable[token.Values["returnType"]]
			if !ok {
				panic("INVALID DATA TYPE")
			}
			out += "func " + token.Values["name"] + " " + val
			var function = token.Values["innards"]
			out += function[:1]
			function = function[1:]
			tokens, err := tokenizeFunc(function[:len(function)-1])
			if err != nil {
				panic(err)
			}
			outPlus, imps := parse(tokens)
			out += outPlus
			imports = merge(imports, imps)
			out += "}"
		// Func Call Handler
		case 3:
			val, ok := funcTransTable[token.Values["name"]]
			if ok {
				out += " " + val + token.Values["params"] + ";"
			} else {
				val = importFuncTransTable[token.Values["name"]]
				imp := importTable[token.Values["name"]]
				if !contains(imports, imp) {
					imports = append(imports, imp)
				}
				out += " " + val + token.Values["params"] + ";\n"
			}

		// Main Handler
		case 9:
			out += "func main ()"
			var function = token.Values["innards"]
			out += function[:1]
			function = function[1:]
			tokens, err := tokenizeFunc(function[:len(function)-1])
			if err != nil {
				panic(err)
			}
			outPlus, imps := parse(tokens)
			out += outPlus
			imports = merge(imports, imps)
			out += "}"
		}

	}

	return out, imports
}
