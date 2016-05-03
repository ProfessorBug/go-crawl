package main

import (
  "fmt"
  "flag"
  "os"
  "net/http"
  "io/ioutil"
  "golang.org/x/net/html"
  "strings"
  "log"
)

type person struct {
  name string
  age  int
}


type urlStats struct {
    visited  bool
    count    int
    status   int
}

var urlMap = map[string]urlStats{}

func main() {

  // Deal with flags
  // ----
  // valid flags:
  // -url "www.golang.org"
  // etc output type and filename etc
  urlPtr := flag.String("url", "http://www.github.com", "URL to crawl")

  flag.Parse()
  url := *urlPtr

  finishLogic := false






 /* uS := urlStats{visited: false, count: 1}

  urlMap["lalallalallala"] = uS

  uS2 := urlStats{visited: false, count: 14}

  addURL("github.com", urlMap)
  urlMap["github.com"] = uS2
*/

  addURL("BBB.com")
  addURL("BBB.com")
  addURL("github.com")



  //main loop
  for !finishLogic {
    getUrl(url)

//    getUrl("http://google.com")
    finishLogic = true
  }

  fmt.Println("Starting point:", *urlPtr)

  fmt.Println(len(os.Args))
  writeConsole()


  fmt.Println(urlMap)
}

func getUrl(url string) {

    // URL stats add 1 to the count. uS is 0 when the url hasn't been found
    uS := urlMap[url]
    uS.visited = true
    urlMap[url] = uS

    response, err := http.Get(url)

    if err != nil {
       fmt.Println("-----------")
        fmt.Printf("%s", err)
        //os.Exit(1)
    } else {
      fmt.Println("-----------")
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }

    //    fmt.Printf("%s\n", string(contents))


        s := string(contents)
        doc, err := html.Parse(strings.NewReader(s))
        if err != nil {
          log.Fatal(err)
        }
        var f func(*html.Node)
        f = func(n *html.Node) {
          if n.Type == html.ElementNode && n.Data == "a" {
            for _, a := range n.Attr {
              if a.Key == "href" {
                link := a.Val
                if(len(a.Val)>1){
                  firstOne := a.Val[0:1]
                  firstTwo := a.Val[0:2]
                  if(firstTwo == "//") {
                    link = "http://" + a.Val[2:len(a.Val)]
                    //fmt.Println("=-=-=-=-=-=-=-=-")
                    //fmt.Println(a.Val)
                  }else{
                    if(firstOne == "/") {
                      //link = url + a.Val[1:len(a.Val)]
                      link = url + a.Val
                    }
                  }
                }


                fmt.Println(link)
                addURL(a.Val)
                break
              }
            }
          }
          for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
          }
        }
        f(doc)



    }
}

func writeConsole() {
  fmt.Println("---------------: Crawl Level: " + "1" + " :---------------")

  fmt.Println("www.example.com/1 " + " 404 " + "Links found: " + " 4 " + " Referenced: 23 times")
  fmt.Println("www.example.com/2 " + " 404 " + "Links found: " + " 1 " + " Referenced: 2  times")
  fmt.Println("www.example.com/3 " + " 404 " + "Links found: " + " 0 " + " Referenced: 25 times")

  fmt.Println("---------------: Crawl Level: " + "2" + " :---------------")

  fmt.Println("in thingy")



}

func writeXls() {

}


func addURL(url string) {

  // URL stats add 1 to the count. uS is 0 when the url hasn't been found
  uS := urlMap[url]
  uS.count += 1
  urlMap[url] = uS

}

