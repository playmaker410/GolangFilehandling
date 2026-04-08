package main
import ("os"
	"fmt"
	//"io"
	 //"log"
	"bufio"
	"strings"
	"encoding/json"
	"encoding/csv"
)

//READING FILE

func FileRead() {
			filepath:= "data.txt"
			file, err:=os.ReadFile(filepath)

			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(string(file))
		}


//CREATE AND WRITE INSIDE THE FILE
func WriteFile(){


	file, err:= os.Create("data.txt")

	if err != nil{
		fmt.Println("Error Encountered", err)
		return
	}
	defer file.Close()


	//var text string
	//fmt.Println("Enter your words here")
	//fmt.Scanln(&text)
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter your words here")
	text, err := reader.ReadString('\n')

	text = strings.TrimSpace(text)
	_,err = file.WriteString(text)
	if err != nil{
	fmt.Println(err)
			}
//Note: scaln accept only one value
//data:= (string(text))
//	_, err = file.WriteString(data)
//	if err != nil{
//	fmt.Println(err)
//}
	fmt.Println("File Writing Successsfully done")
}


//APPEND TEXT TO AN EXISTING FILE WITHOUTH CLEANING THE EXISTING FILE 

func FileAppend()  {

data, err :=os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
if err != nil{
	fmt.Println("Error Encountered", err)
	return
}

Append:= bufio.NewReader(os.Stdin)
fmt.Println("Add text to document")
text, err := Append.ReadString('\n')

text = strings.TrimSpace(text)
_, err = data.WriteString(text)
if err != nil {
fmt.Println(err)
}

fmt.Println("Content Appended succesfully")
}


//READ JSON



type  Info struct{
    DBHost        string `json:"database_host"`
    DBPort        int    `json:"database_port"`
    DBUsername    string `json:"database_username"`
    DBPassword    string `json:"database_password"`
    ServerPort    int    `json:"server_port"`
    ServerDebug   bool   `json:"server_debug"`
    ServerTimeout int    `json:"server_timeout"`
}


func ReadJson() {
filepath:= "info.json"
var  Reg Info
file, err:= os.Open(filepath)

if err != nil{
fmt.Println(err)
}
defer file.Close()

reader:= json.NewDecoder(file)
err = reader.Decode(&Reg)

if err != nil {
fmt.Println("ERROR ENCOUNTERD",err)
}

fmt.Println(Reg)
}


//Read CSV

func ReadCsv(){
filepath:= "csvRead.csv"

file, err := os.Open(filepath)

if err != nil{
fmt.Println(err)
}

defer file.Close()
Read:= csv.NewReader(file)
record, err := Read.ReadAll()

if err != nil {
fmt.Println(err)
}
fmt.Println(record)

}


func CreateDir() {
var text string
fmt.Println("create dir")
fmt.Scanln(&text)
err:=os.Mkdir(text, 0755)

if err != nil {
fmt.Println(err)
}

fmt.Println("Success")
}


func ListFiles(){

var text string
fmt.Println("List dir")
fmt.Scanln(&text)

dir, err:=os.ReadDir(text)

if err != nil{
fmt.Println(err)
}

for _, path := range dir {
fmt.Println(path.Name())
}
fmt.Println("success")

}


func Delete() {
var text string
fmt.Println("Enter file to Delet")
fmt.Scanln(&text)

err:=os.Remove(text)
if err != nil {
fmt.Println(err)
} else{
fmt.Println("Deleted Succesfully")
}


}






















func main()  {

var Choice int
		for{
			fmt.Println("1. Read file")
			fmt.Println("2. Write file")
			fmt.Println("3. Append file")
			fmt.Println("4. Read JSON")
			fmt.Println("5. Read CSV")
			fmt.Println("6. Create Directory")
			fmt.Println("7. List Files")
			fmt.Println("8. Delete Files")
			fmt.Println("9. Exit")
			fmt.Scanln(&Choice)


 switch Choice {
 case 1:
	FileRead()
	return

 case 2:
	WriteFile()
	return
 case 3:
	FileAppend()
	return
 case 4:
	ReadJson()
	return
case 5:
	 ReadCsv()
	return
case 6: 
	 CreateDir()
	return
case 7:
	 ListFiles()
	return
case 8:
	Delete()
	return
default:
	fmt.Println("invalid choice")	
		 }
		}
	}
