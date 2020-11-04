package main

import (
  "database/sql"
  "fmt"
  "os"
  _ "github.com/lib/pq"
  "net/http"
  "encoding/json"
)

func main() {
  http.HandleFunc("/", getAll)

  PORT := getenv("PORT", "9100")
  http.ListenAndServe(":" + PORT, nil)
}

func getenv(key, fallback string) string {
  if value, ok := os.LookupEnv(key); ok {
     return value
  }
  return fallback
}

type User struct {
  Id        int
  Age       int
  FirstName string
  Username  string
  Email     string
}

func getAll(w http.ResponseWriter, r *http.Request) {
  // DB connection
  psqlInfo := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    getenv("DB_HOST", "app_db"),
    getenv("DB_PORT", "5432"),
    getenv("DB_USER", "dbUser"),
    getenv("DB_PASSWORD", "dbPassword"),
    getenv("DB_NAME", "dbName"),
  )

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("DB has been successfully connected!")

  // request
  fmt.Println("Cool, we received request")

  sql := "SELECT * FROM users;"
  rows, err := db.Query(sql)
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  users := make([]User, 0)

  for rows.Next() {
    var user User
    err = rows.Scan(&user.Id, &user.Age, &user.FirstName, &user.Username, &user.Email)
    if err != nil {
      panic(err)
    }
    users = append(users, user)
  }

  // get any error encountered during iteration
  err = rows.Err()
  if err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(users)
}
