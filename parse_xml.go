// Note: this illustrates ValuesForKey() and ValuesForPath() methods

package main

import (
    "fmt"
    "github.com/clbanning/mxj/v2"
    "encoding/json"
    "log"
)

// var xmldata = []byte(`
//    <books>
//       <book seq="1">
//          <author>William H. Gaddis</author>
//          <title>The Recognitions</title>
//          <review>One of the great seminal American novels of the 20th century.</review>
//       </book>
//       <book seq="2">
//          <author>Austin Tappan Wright</author>
//          <title>Islandia</title>
//          <review>An example of earlier 20th century American utopian fiction.</review>
//       </book>
//       <book seq="3">
//          <author>John Hawkes</author>
//          <title>The Beetle Leg</title>
//          <review>A lyrical novel about the construction of Ft. Peck Dam in Montana.</review>
//       </book>
//       <book seq="4">
//          <author>T.E. Porter</author>
//          <title>King's Day</title>
//          <review>A magical novella.</review>
//       </book>
//    </books>
// `)

var xmldata = []byte(`
    <html>
        <div class="container">
            <h1>Title Here</h1>
            <p class="bold">
                One of the great seminal American novels of the
                <a href="/20th-century.html">20th century</a>.
            </p>
        </div>
        <p>Hello!</p>
        <div class="container">
            <h1>Title Here</h1>
            <p class="bold">
                One of the great seminal American novels of the
                <a href="/20th-century.html">20th century</a>.
            </p>
        </div>
    </html>
`)

func main() {

    m, err := mxj.NewMapXml(xmldata)
    if err != nil {
        log.Fatal("err:", err.Error())
    }

    b, err := json.MarshalIndent(m, "", "  ")
    if err != nil { }

    fmt.Print(string(b) + "\n")


    v, _ := m.ValuesForPath("*")
    fmt.Println("path: * len(v):", len(v))

    for _, vv := range v {
        fmt.Printf("\t%+v\n", vv)
    }


    // v, _ := m.ValuesForKey("div")
    // fmt.Println("path: div; len(v):", len(v))
    // fmt.Printf("\t%+v\n", v)

    // v, _ = m.ValuesForPath("books.book")
    // fmt.Println("path: books.book; len(v):", len(v))
    // for _, vv := range v {
    // 	fmt.Printf("\t%+v\n", vv)
    // }

    // v, _ = m.ValuesForPath("books.*")
    // fmt.Println("path: books.*; len(v):", len(v))
    // for _, vv := range v {
    // 	fmt.Printf("\t%+v\n", vv)
    // }

    // v, _ = m.ValuesForPath("books.*.title")
    // fmt.Println("path: books.*.title len(v):", len(v))
    // for _, vv := range v {
    // 	fmt.Printf("\t%+v\n", vv)
    // }

    // v, _ = m.ValuesForPath("books.*.*")
    // fmt.Println("path: books.*.*; len(v):", len(v))
    // for _, vv := range v {
    // 	fmt.Printf("\t%+v\n", vv)
    // }
}