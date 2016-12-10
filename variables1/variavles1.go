package variable1

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    var v int = 25

    fmt.Fprintf(w, "vに設定されていいるのは%dです。", v)
}
