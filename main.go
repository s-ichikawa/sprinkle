package main

import (
    "math/rand"
    "time"
    "bufio"
    "os"
    "fmt"
    "strings"
    "io/ioutil"
)

const otherWord = "*"
var transforms = []string{}
func main() {

    buffer, _ := ioutil.ReadFile("other_word.txt")
    for _, line := range strings.Split(string(buffer), "\n") {
        transforms = append(transforms, line)
    }

    rand.Seed((time.Now().UTC().UnixNano()))
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        t := transforms[rand.Intn(len(transforms))]
        fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
    }
}
