package main

import (
    // "fmt"
    "fmt"
    // "os"
)

func main() {

    // read argument and file, rewrite blocks to array
    blocks, err := readFile()
    if err != nil {
        fmt.Println(err)
        return
    }

    // checking blocks format
    goodBlock := checkBlocks(blocks)
    if goodBlock != nil {
        fmt.Println(goodBlock.Error())
        return
    }
    // checking tets format
    for i := 0; i < len(blocks); i+=4 {
        block := blocks[i:i+4]
        err = checkTets(block)
        if err != nil {
            fmt.Println(err)
            return
        }
    }


    // get each tet
    Tets := [][][]rune{}
    for i := 0; i < len(blocks); i+=4 {
        block := blocks[i:i+4]
        tet := getTet(block)
        Tets = append(Tets, tet)
    }

    FieldSize := getMinFiledSize(Tets)

    Field := buildField(FieldSize)
    placeTets(Tets, Field, FieldSize, 0)

    // for , tet := range allTets {
    //     for , s := range tet {
    //         for , l := range s {
    //             fmt.Print(string(l))
    //         }
    //         fmt.Println()
    //     }
    //     fmt.Println()
    // }
    // for , tet := range blocks {
    //     for _, s := range tet {

    //         fmt.Print(string(s))
    //     }
    //     fmt.Println()
    // }

}