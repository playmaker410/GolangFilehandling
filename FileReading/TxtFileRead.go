package main
import("fmt"
		"os")



func main()  {
	file:= "home.txt"
txt, err := os.ReadFile(file)

if err != nil{
	fmt.Println("Erro is encountered",)
	return
}
fmt.Println(string(txt))
	
}