package main


import (
    "fmt"
    "log"
    "net/http"
	"os/exec"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/opt/test/executor")
	stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }
	res := string(stdout)
    fmt.Println(res)
    fmt.Fprintf(w, res)
}


func main() {
    http.HandleFunc("/", helloHandler) // Update this line of code


    fmt.Printf("Starting server at port 80\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}