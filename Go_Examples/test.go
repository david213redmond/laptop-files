package main
import "fmt"
func main() {
	for {
		var val int
		_, err := fmt.Scan(&val)	
		if {
			break
		}
		if err != nil{
			fmt.Println(err)
			break
		}
		//fmt.Println(val)
	}
}