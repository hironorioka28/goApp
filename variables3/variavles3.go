package variable3

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    var (
        i int = 25
        j = 123.456
        k int
        l, m uint64 = 1, 2
        millisec int64 = 1e3
        intValue, floatValue, strValue = 1, 2.0, "Hello"
    )

    fmt.Fprintf(w, "i = %d\n", i)
    fmt.Fprintf(w, "j = %f\n", j)
    fmt.Fprintf(w, "k = %d\n", k)
    fmt.Fprintf(w, "l = %d\n", l)
    fmt.Fprintf(w, "m = %d\n", m)
    fmt.Fprintf(w, "millisec = %d\n", millisec)
    fmt.Fprintf(w, "intValue = %d\n", intValue)
    fmt.Fprintf(w, "floatValue = %f\n", floatValue)
    fmt.Fprintf(w, "strValue = %q\n", strValue)

}
