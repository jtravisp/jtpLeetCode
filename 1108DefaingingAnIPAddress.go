package main

import (
    "strings"
)

func defangIPaddr(address string) string {
    // Replace all occurrences of "." with "[.]" using strings.ReplaceAll
    defangedAddress := strings.ReplaceAll(address, ".", "[.]")
    return defangedAddress
}

func main() {
    // Test cases
    address1 := "1.1.1.1"
    address2 := "255.100.50.0"

    defanged1 := defangIPaddr(address1)
    defanged2 := defangIPaddr(address2)

    println("Input address 1:", address1)
    println("Defanged address 1:", defanged1)

    println("Input address 2:", address2)
    println("Defanged address 2:", defanged2)
}
