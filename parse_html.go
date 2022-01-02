package main

import (
    "fmt"
    "log"
    "strings"
    "golang.org/x/net/html"
)

func main() {
    s := `<html>
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
    `

    doc, err := html.Parse(strings.NewReader(s))
    if err != nil {
        log.Fatal(err)
    }

    var f func(*html.Node)

    f = func(n *html.Node) {

        if n.Type == html.ElementNode {
            fmt.Printf("node %v \n", n.Data);

            for _, a := range n.Attr {
                fmt.Printf("  %v = %v \n", a.Key, a.Val)
            }
        }


        if n.Type == html.TextNode {
            text := strings.TrimSpace(n.Data)
            fmt.Printf("text %v \n", text);
        }

        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
}