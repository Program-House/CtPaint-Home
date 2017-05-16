package router

import (
    "io"
    "io/ioutil"
    "net/http"
    . "./admin/sign-in"
)

func SetUp() {
    http.HandleFunc("/api/admin/signin", SignIn)
    http.HandleFunc("/", Home)
}

func Home(res http.ResponseWriter, req *http.Request) {

    switch req.Method {
    case "GET" :
        res.Header().Set("Content-Type", "text/html; charset=UTF-8")

        webpage, err := ioutil.ReadFile("./apps/test.html")

        if err == nil {
            io.WriteString(res, string(webpage))  
        } else {
            io.WriteString(res, "Error!")
        }
    }
}
