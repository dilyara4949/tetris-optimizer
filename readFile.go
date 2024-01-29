package main

import (
    // "fmt"
    "errors"
    // "fmt"
    "os"
    "strings"
    "unicode"
)

func readFile() ([]string, error) {
    arg := os.Args[1:]
    if len(arg) != 1 {
        return nil, errors.New("argument size not correct")
    }
    dat, _ := os.ReadFile(arg[0])

    input := string(dat)
    sep := func(c rune) bool {
        return unicode.IsSpace(c) || c == 13 || c == 10
    }

    blocks := strings.FieldsFunc(input, sep)

    // for , tet := range blocks {
    //     for _, s := range tet {

    //         fmt.Print(string(s))
    //     }
    // }

    return blocks, nil
}