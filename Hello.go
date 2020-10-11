package main
import "fmt"
import "time"

func loop(){
	for i:=0;i<5;i++{
		time.Sleep(2*time.Second)
		fmt.Println(i)
	}
}

func main(){
	
	fmt.Println("Murali is back")
	go loop()
	go loop()
	fmt.Scanln()

}
