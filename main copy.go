// // // // Run codeCopy code
// package main

// import (
// 	"log"
// 	"sort"
// )

// type Employee struct {
// 	Name     string
// 	Pos      int
// 	Children []Employee
// }

// type User struct {
// 	Name   string `json:"name"`
// 	Type   string `json:"type"`
// 	Age    int    `json:"Age"`
// 	Social Social `json:"social"`
// }

// type Social struct {
// 	Facebook string `json:"facebook"`
// 	Twitter  string `json:"twitter"`
// }

// type Bird struct {
// 	Species     string
// 	Description string
// }

// // func main() {
// // 	// an instance of our User struct
// // 	// social := Social{Facebook: "https://facebook.com", Twitter: "https://twitter.com"}
// // 	// user := User{Name: "LanKa", Type: "Author", Age: 25, Social: social}

// // 	// byteArray, err := json.Marshal(user)
// // 	// if err != nil {
// // 	// 	fmt.Println(err)
// // 	// }

// // 	// fmt.Println(string(byteArray))

// // 	// birdJson := `[{"species":"pigeon", "description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
// // 	// var birds []Bird

// // 	// json.Unmarshal([]byte(birdJson), &birds)
// // 	// fmt.Printf("Birds : %+v", birds)

// // }

// // func main() {

// //////////
// // bytes := []byte(data)
// // e := &Employee{}

// // err := json.Unmarshal(bytes, e)
// // if err != nil {
// // 	log.Fatal(err)
// // }
// // log.Println(e.Name)
// // ///
// // printChild_Name(e)
// // ///
// // sortEmp(e)
// // ///
// // log.Println("-------")
// // printChild_Name(e)

// // }

// func printChild_Name(e *Employee) {
// 	for i := 0; i < len(e.Children); i++ {
// 		log.Println("User Name: " + e.Children[i].Name)
// 		printChild_Name(&e.Children[i])
// 	}
// }

// func sortEmp(e *Employee) {
// 	sort.Slice(e.Children, func(i, j2 int) bool {
// 		return e.Children[i].Name < e.Children[j2].Name
// 	})
// }
