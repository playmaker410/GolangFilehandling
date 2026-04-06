package  main
 import ("os"
	"log"
	)

//WITH THE HELP OF  os.O_APPEND THE FILE ALWAYS UPDATE  BUT WITHOUT IT ALWAYS OVERWRITE THE EXISTING CONTENT

func main() {
createfile, err := os.OpenFile("contact.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

if err != nil {
log.Fatal("Encountering Error", err)
}
 defer createfile.Close()

writeText := "\n OSHIM ODINAKA JOSHUA- 09072324110 \n"
_, err = createfile.WriteString(writeText)

if err != nil {
log.Fatal("Encountering Eror here", err)
}

}

