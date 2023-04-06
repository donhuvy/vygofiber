package main

import "fmt"

func main2() {
	const name, dept = "Geek for foo", "CS"
	ee := fmt.Errorf("anh yeu em %q %q", name, dept)
	fmt.Println(ee.Error())
}
