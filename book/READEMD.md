### goweb

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world", r.URL.Path)
}
func main() {
    http.HandleFunc("/", handler)
    
    http.ListenAndServe("localhost:8080", nil)
}

```

### 自定义请求

```go
type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "自定义请求")
}
func main() {
    handler := MyHandler{}
    http.Handle("/hello", &handler)
    http.ListenAndServe("localhost:8080", nil)
}

```