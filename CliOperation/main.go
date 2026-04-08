package main
import ("os"
		"fmt"
		//"io"
		 //"log"
		 )



func FileRead() {
			filepath:= "data.txt"
			file, err:=os.ReadFile(filepath)

			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(string(file))
		}

func WriteFile(){
	var text string
	fmt.Println("Enter your words here")
	fmt.Scanln(&text)

	err:= os.WriteFile("dato.txt", []byte(text), 0644)

	if err != nil{
		fmt.Println("Error Encountered", err)
		return
	}

	fmt.Println("File Writing Successsfully done")
}

func FileAppend()  {
	var filename string
	var content string

	fmt.Println("Add to exsiting file")
	fmt.Scanln(&filename)

data, err :=os.OpenFile("data.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
if err != nil{
	fmt.Println("Error Encountered", err)
	return
}	

defer data.Close()
_,err = data.WriteString(content + "\n")
if err != nil{
	fmt.Println("Error Writing", err)
	return
}
fmt.Println("Content Appended succesfully")
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
	fmt.Println("FILE READC")
	FileRead()
	return

 case 2:
	WriteFile()
	return
 case 3:
	FileAppend()
	return


default:
	fmt.Println("invalid choice")	
		 }
		}
	}