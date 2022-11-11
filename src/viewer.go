package main


import (
    "fmt"
    "log"
    "net/http"
 "os/exec"
)
var conCount int
func helloHandler(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("/opt/test/executor")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }
if conCount > 0 {
 res := string(stdout)
    fmt.Println(res)
    fmt.Fprintf(w, res)
    if r.URL.Path == "/" {
    conCount -= 1
    }
    } else {
        w.WriteHeader(500)
    }
}


func main() {
    conCount = 10
    http.HandleFunc("/", helloHandler) // Update this line of code


    fmt.Printf("Starting server at port 80\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}