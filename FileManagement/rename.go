package main 
import ("os"
	"fmt"
	)



func main(){
 err:= os.Rename("home.txt", "Contact.txt")

if err != nil{
fmt.Println(err)
}

}
