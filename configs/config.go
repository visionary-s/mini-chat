package configs

var AppJsonConfig = []byte(`
{
  "app": {
    "port": "8081",
    "cookie_key": "5174vydhibd60c7345mtdgt",
    "serve_type": "GoServe"
  },
  "mysql": {
    "dsn": "root:root@tcp(127.0.0.1:3306)/mini_chat?charset=utf8mb4&parseTime=True&loc=Local"
  }
}
`)
