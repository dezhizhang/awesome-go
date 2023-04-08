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
### 插入和获取数据库
```go
func AddUser() error {
	sqlStr := "insert into users(id,name,password,email) values (?,?,?,?)"
	result, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		log.Printf("预编译失败%s", err)
	}

	_, err = result.Exec("12312", "刘德华", "123456", "1541609448@qq.com")
	if err != nil {
		log.Printf("添加失败%s", err)
	}
	return nil
}

func GetUserQuery() (model.User, error) {
	sqlStr := "select * from users where id = ?"
	row := dao.DB.QueryRow(sqlStr, "456")

	user := model.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	if err != nil {
		log.Printf("扫描失败%s", err)
	}
	return user, nil
}

func GetUsersQuery() ([]model.User, error) {
	sqlStr := "select * from users"
	rows, err := dao.DB.Query(sqlStr)
	if err != nil {
		log.Printf("查询失败%s", err)
	}
	var user model.User
	var users []model.User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			log.Printf("扫描失败%s", err)
		}
		users = append(users, user)
	}
	return users, nil
}

```
### 获取表单的值
```go
func handler(w http.ResponseWriter, r *http.Request) {

	length := r.ContentLength
	bytes := make([]byte, length)
	_, err := r.Body.Read(bytes)
	if err != nil {
		log.Printf("读取失败%s", err)
	}
	fmt.Fprintln(w, string(bytes))
}

func main() {
	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}
}
```
### 获取post参数
```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "请求参数", r.PostFormValue("username"))
}

func main() {
    http.HandleFunc("/hello", handler)
    err := http.ListenAndServe("localhost:8080", nil)
    if err != nil {
        log.Printf("启动服务失败%s", err)
    }
}

```
### 返回json数据
```go
func handler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	user := model.User{
		Id:       "123",
		Name:     "小明",
		Password: "123456",
		Email:    "1541609448@qq.com",
	}
	marshal, err := json.Marshal(user)
	if err != nil {
		log.Printf("序列化失败%s", err)
	}

	w.Write(marshal)
}

func main() {
	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}
}

```
### 加载模板
```go
func handler(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Printf("加载模板失败%s", err)
	}
	files.Execute(w, nil)
}

func main() {
	http.HandleFunc("/hello", handler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}

}

```