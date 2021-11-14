package main
import "fmt"
var valor1,valor2 int
func somar(a int,b int){
  fmt.Println("A soma do primeiro valor com o do segundo valor é",a+b)
}
func main(){
  for(true){
    fmt.Printf("Digite o primeiro valor: ");
    fmt.Scan(&valor1);
    fmt.Printf("Digite o segundo valor: ");
    fmt.Scan(&valor2);
    somar(valor1,valor2);
  }
}
