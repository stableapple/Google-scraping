package main
import (
  "fmt"
  "os"
  "flag"
  "log"
  

  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
  
)
  func main(){

    var user string
    fmt.Println("Enter username:")
    fmt.Scanln(&user)
    config := oauth1.NewConfig("consumerKey", "consumerSecret")
    token := oauth1.NewToken("accessToken", "accessSecret")
    httpClient := config.Client(oauth1.NoContext, token)
    client := twitter.NewClient(httpClient)
    just := flag.String("",user,"")   
    flag.Parse()
    d1, err := os.Create("Number of followers of"+user+",txt")
    if err != nil {
       log.Fatal(err)
    list, resp, err := client.Followers.List(&twitter.FollowerListParams{ScreenName: *just})
    var num int
    num = 0
    fmt.Println(resp, err)
    for _, list := range list.Users {
    num++
    }

     if err != nil {
      log.Fatal(err)
    }
    d1.WriteString("\n" + list.ScreenName)
    d1.WriteString("\n")
    d1.WriteString(fmt.Sprintf("%d\n",num))
    d1.Close()
    }
}

