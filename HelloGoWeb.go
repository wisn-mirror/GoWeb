package main
import(
	"io"
	"log"
	"net/http"
)
func HelloGoWebHttp() {
	http.HandleFunc("/",HelloHttp)
	error := http.ListenAndServe(":8000", nil)
	if error!=nil{
		log.Fatal(error)
	}
}
func HelloHttp( responseWriter http.ResponseWriter,request *http.Request){
	io.WriteString(responseWriter,"hello go web ! 		-wisn")
}