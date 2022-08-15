package main
import "fmt"
import "log"
import "net/http"
func main() {
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to my home page\n")
	fmt.Println("Endpoint hit: homePage")

}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to about page\n")
	fmt.Println("Endpoint hit: aboutMe")

}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutMe)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
