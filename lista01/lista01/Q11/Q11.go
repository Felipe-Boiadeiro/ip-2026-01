11- 

package main 

 

import "fmt" 

 

var num int 

 

func main() { 

    fmt.Println("Digite um número para saber se ele é divisível por 3 e por 5: ") 

    fmt.Scan(&num) 

 

    if num%3 == 0 && num%5 == 0 { 

        fmt.Println("O NUMERO E DIVISIVEL") 

    } else { 

        fmt.Println("O NUMERO NAO E DIVISIVEL") 

 

    } 

 

} 

 