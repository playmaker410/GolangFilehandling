package main 
import ("encoding/csv"
	"fmt"
	"os"
	)



func main(){

Users:="Users.csv"
file, err:= os.Open(Users)

if err != nil{
fmt.Println("Error encountered", err)
return
}

defer file.Close()

Reading:=csv.NewReader(file)

data, err := Reading.ReadAll()

if err != nil{
fmt.Println("ERROR ENCOUNTERED", err)
}

fmt.Println(data)
}
