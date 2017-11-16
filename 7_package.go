package main

import "hello/packages"

func main() {
	println(packages.GetSmallVar())

	println(packages.LargeConst)
}