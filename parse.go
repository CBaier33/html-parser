package htmlParser

import (
    "io"
    "strings"
    "golang.org/x/net/html"
)

// Link in an HTML document.
type Link struct {
    Href string
    Text string
}

func Parse(r io.Reader) ([]Link, error) {
    doc, err := html.Parse(r)
    if err != nil{
        return nil, err
    }
    // dfs(doc, "")
    nodes := linkNodes(doc)
    var links []Link
    for _, node := range nodes{
        links = append(links, buildLink(node))
    }
    return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
    if n.Type == html.ElementNode && n.Data == "a" {
        return []*html.Node{n}
    }
    var ret []*html.Node
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret = append(ret, linkNodes(c)...)
    }
    return ret
}

func buildLink(n *html.Node) Link {
    var ret Link
    for _, attr := range n.Attr {
        if attr.Key == "href" {
            ret.Href = attr.Val
            break
        }
    }
    ret.Text = text(n)
    return ret
}

func text(n *html.Node) string {
    if n.Type == html.TextNode {
        return n.Data
    }
    if n.Type != html.ElementNode {
        return ""
    }
    var ret string
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret += text(c)
    }
    return strings.Join(strings.Fields(ret), " ")
}


// func dfs(n *html.Node, padding string) {
//     msg := n.Data
//     if n.Type == html.ElementNode {
//         msg = "<" + msg + ">"
//     }
//     fmt.Println(padding, msg)
//     for c := n.FirstChild; c != nil; c = c.NextSibling {
//         dfs(c, padding + "  ")
//     }
// }
