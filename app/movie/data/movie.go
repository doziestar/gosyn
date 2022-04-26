package data

// import "github.com/doziestar/gosyn/movie/data"

type Movie struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Year     int64     `json:"year"`
	Director *Director `json:"director"`
}

// type Director struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name"`
// }
