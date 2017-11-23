package entitybase

// import "fmt"

type book struct {
	Id     int    `json:id`
	Name   string `json:name`
	Author string `json:author`
}

//
// func main() {
// 	fmt.Println("vim-go")
// }
