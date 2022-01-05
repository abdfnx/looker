package main

import (
    "fmt"

    "github.com/abdfnx/looker"
)

func main() {
    curlPath, err := looker.LookPath("curl")

    if err != nil {
        fmt.Println(err)
    }

	fmt.Println(curlPath)
}
