package main

import "fmt"

func main() {
    input := "Do not give opnions or advice unless you are asked"
    rev := Reverse(input)
    doubleRev := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q\n", rev)
    fmt.Printf("reversed again: %q\n", doubleRev)
}

func Reverse(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i<len(b)/2; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}
