package main
import(
	"io"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/",helloHttp)
	error := http.ListenAndServe(":8000", nil)
	if error!=nil{
		log.Fatal(error)
	}
}
func helloHttp( responseWriter http.ResponseWriter,request *http.Request){
	io.WriteString(responseWriter,"hello go web ! 		-wisn")

}
