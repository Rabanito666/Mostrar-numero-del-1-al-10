package numeros

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"

)



func Tabla(){
	convertir:= bufio.NewReader(os.Stdin)
	var pito int

	for{
	fmt.Println("ingrese un numero")
	numero, _ := convertir.ReadString('\n')

	numero = strings.TrimSpace(numero)
	tabla, err:= strconv.Atoi(numero)

	if err != nil{
		fmt.Println("esto no es un numero bruto")
		continue
	}
	
	pito = tabla
	break
	}
	fmt.Println("tabla de multiplicar de",pito)
	for i:= 1; i<=10; i++{
		fmt.Println(pito,"x",i,"=",pito*i)
}
	



}
