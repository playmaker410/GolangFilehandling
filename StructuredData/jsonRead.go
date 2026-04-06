package main 
import("encoding/json"
	"fmt"
	"os"
		)



type  JsonConfig struct{
DBHost 		string  `json:"database_host"`
DBPort 		int 	`json:"database_port"`
DBUser 		string 	`json:"database_username"`
DBPassword 	string  `json:"database_password"`
ServerPort 	int 	`json:"server_port"`
ServerDebug 	bool  	`json:"server_debug"`
ServerTimeout 	int 	`json:"server_timeout"`
}


func main(){
jsonFile:="config.json"

var config JsonConfig
 data, err := os.Open(jsonFile)

if err != nil{

fmt.Println("Error encounterded", err)
return
}

defer data.Close()


Readjson := json.NewDecoder(data)

err = Readjson.Decode(&config)

if err != nil{
fmt.Println("ERROR", err)
}

fmt.Println(config)
}


