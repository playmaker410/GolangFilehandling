package main
import("bufio"
		"fmt"
		"os"
		)

func Printlastno(lines []string, num int)[]string{
	var result []string
	for i:= len(lines) - num; i < len(lines); i++{
		result = append(result, lines[i])
	}
	return result
}




func main(){
	file:= "home.txt"
	info,err := os.Open(file)
	
	if err != nil {
		fmt.Println("Error encountered", err)
		return
	}
defer info.Close()

var lines []string

scanner:= bufio.NewScanner(info)

for scanner.Scan(){
lines = append(lines, scanner.Text())	
}

if err := scanner.Err(); err != nil{
	fmt.Println("error encountered",err)
}

lastthree := Printlastno(lines, 3)

for _, line := range lastthree{
	fmt.Println(line)
	fmt.Println("__________________")
}


}