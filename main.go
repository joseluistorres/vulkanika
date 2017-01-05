package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "os"
)

func main() {
    arg := os.Args[1]
    url := "https://"+os.Getenv("RYVER_ACCOUNT")+".ryver.com/api/1/odata.svc/forums("+os.Getenv("FORUM_ID")+")/Chat.PostMessage()"

    m := map[string]interface{}{
      "body": "deploying to web-dev the latest from @" +arg}
    mJson, _ := json.Marshal(m)
    contentReader := bytes.NewReader(mJson)

    req, err := http.NewRequest("POST", url, contentReader)
    req.Header.Set("Content-Type", "application/json")
    req.SetBasicAuth(os.Getenv("RYVER_USR"),os.Getenv("RYVER_KEY"))
    req.Header.Set("Accept", "*/*")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
