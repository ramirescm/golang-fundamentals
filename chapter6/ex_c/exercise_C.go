package main

import "fmt"

var xpto interface{}

func main() {
	x := 10
	switch {
	case x > 5:
		fmt.Println("> 5 val x ", x)
		break
	case x >= 10:
		fmt.Println(">= 10 val x ", x)
		break
	case x > 15:
		fmt.Println(" > 15 val x ", x)
		break
	default:
		fmt.Println("default val x ", x)
		break
	}

	name := "dell"
	switch name {
	case "hp":
		fmt.Println("hp")
	case "dell":
		fmt.Println("dell")
	case "mac":
		fmt.Println("mac")
	default:
		fmt.Println("not found")
	}

	balance := 0
	switch balance {
	case 0, 1:
		fmt.Println("balance 0 ou 1")
	case 3, 6:
		fmt.Println("balance 3 ou 6")
	default:
		fmt.Println("not found balance")
	}

	expression := 2
	switch {
	case (expression == 2), (expression == 4):
		fmt.Println("expression 2 ou 4 and fallthrough")
		fallthrough
	case (expression == 1), (expression == 3):
		fmt.Println("expression 1 ou 3")
	default:
		fmt.Println("not found expression")
	}

	xpto = "teste"
	switch xpto.(type) {
	case int:
		fmt.Println("o tipo é int")
	case bool:
		fmt.Println("o tipo é bool")
	case float64:
		fmt.Println("o tipo é float64")
	case string:
		fmt.Println("o tipo é string")
	default:
		fmt.Println("tipo nao encontrado")
	}

	switch aaa := 1; {
	case aaa == 1:
		fmt.Println("o aaa é 1")
	case aaa == 2:
		fmt.Println("o aaa é 2")
	default:
		fmt.Println("tipo nao encontrado")
	}
}
