package web_client

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "testing"
    "net/http"
)

type web_client struct{
	
}

func TestHttpGet(t *testing.T) {
    response, err := http.Get("https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1")
    if err != nil {
        t.Error(err)
        panic(err)
    }
    defer response.Body.Close()
    t.Log(response)
    fmt.Println(response.StatusCode)
    fmt.Println(response.Status)
    fmt.Println(response.Header)
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(body))
}

func TestHttpGetHeader(t *testing.T) {
    targetUrl := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"

    req, _ := http.NewRequest("GET", targetUrl, nil)

    req.Header.Add("Authorization", "xxxx")

    response, err := http.DefaultClient.Do(req)
    if err != nil {
        t.Error(err)
        panic(err)
    }
    defer response.Body.Close()
    t.Log(response)
}

func TestHttpPost(t *testing.T) {
    targetUrl := "https://blog.csdn.net/zyndev"

    payload := strings.NewReader("a=111")

    response, err := http.Post(targetUrl, "x-www-form-urlencoded", payload)

    if err != nil {
        t.Error(err)
        panic(err)
    }
    defer response.Body.Close()
    t.Log(response)
}

func TestHttpPostForm(t *testing.T) {
    targetUrl := "https://blog.csdn.net/zyndev"

    payload := url.Values{"key": {"value"}, "id": {"123"}}

    response, err := http.PostForm(targetUrl, payload)

    if err != nil {
        t.Error(err)
        panic(err)
    }
    defer response.Body.Close()
    t.Log(response)
}

func TestHttpPostJSON(t *testing.T) {
    targetUrl := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"

    payload := strings.NewReader("{\"name\":\"张瑀楠\"}")

    req, _ := http.NewRequest("POST", targetUrl, payload)

    req.Header.Add("Content-Type", "application/json")

    response, err := http.DefaultClient.Do(req)

    if err != nil {
        t.Error(err)
        panic(err)
    }

    defer response.Body.Close()
    t.Log(response)

}

func TestHttpPut(t *testing.T) {

    targetUrl := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"

    payload := strings.NewReader("{\"name\":\"张瑀楠\"}")

    req, _ := http.NewRequest("PUT", targetUrl, payload)

    req.Header.Add("Content-Type", "application/json")

    response, err := http.DefaultClient.Do(req)

    if err != nil {
        t.Error(err)
        panic(err)
    }

    defer response.Body.Close()
    t.Log(response)
}

func TestHttpPatch(t *testing.T) {
    targetUrl := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"

    payload := strings.NewReader("{\"name\":\"张瑀楠\"}")

    req, _ := http.NewRequest("PATCH", targetUrl, payload)

    req.Header.Add("Content-Type", "application/json")

    response, err := http.DefaultClient.Do(req)

    if err != nil {
        t.Error(err)
        panic(err)
    }

    defer response.Body.Close()
    t.Log(response)
}

func TestHttpDelete(t *testing.T) {

    targetUrl := "https://ddbc5ffb-c596-4f78-a99d-a6ea93bdc14f.mock.pstmn.io/user/1"

    req, _ := http.NewRequest("DELETE", targetUrl, nil)

    req.Header.Add("Authorization", "xxxx")

    response, err := http.DefaultClient.Do(req)

    if err != nil {
        t.Error(err)
        panic(err)
    }

    defer response.Body.Close()
    t.Log(response)
}