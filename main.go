package main


import (
	"fmt"
	"log"
	"net/http"

)
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path !="/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
		 
	}
	if r.Method!="GET"{
		http.Error(w,"method is not supported ",http.StatusNotFound)
	return
	}
	fmt.Fprintf(w,"hello")

}
func formHandler(w http.ResponseWriter,r *http.Request){
	if err:=r.ParseForm();err!=nil{
fmt.Fprintf(w,"ParseForm() err:%v",err)
return
	}
	fmt.Fprintf(w,"POST request successful")
	name:=r.FormValue("name")
address:=r.FormValue("address")
age:=r.FormValue("age")
fmt.Fprintf(w,"name=%s\n",name)
fmt.Fprintf(w,"address=%s\n",address)
fmt.Fprintf(w,"age=%s\n",age)


}

func main() {
	fileServer:=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Printf("starting server ata port 8080\n")
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}

	
}