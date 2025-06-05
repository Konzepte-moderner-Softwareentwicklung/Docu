package main  
import ( "log"; "net/http" )  
func main() {  
http.Handle("/", http.FileServer(http.Dir(".")))  
log.Println("Serving on :8000...")  
log.Fatal(http.ListenAndServe(":8000", nil))  
}  
