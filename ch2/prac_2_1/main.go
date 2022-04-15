package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	convKind := os.Args[1]
	for _,arg := range os.Args[2:]{
		t, err:= strconv.ParseFloat(arg,64)
		if err != nil{
			fmt.Fprintf(os.Stderr,"conv: %v\n",err)
			return
		}
		if(convKind == "temp"){
			f := Fahrenheit(t)
			c := Celsius(t)
			fmt.Printf("%s = %s , %s = %s\n",f,FtoC(f),c,CtoF(c))
		} else if(convKind == "length"){
			f := Feet(t)
			m := Meter(t)
			fmt.Printf("%s = %s , %s = %s\n",f,FtoM(f),m,MtoF(m))
		} else if(convKind == "weight"){
			p := Pond(t)
			k := Kilogram(t)
			fmt.Printf("%s = %s , %s = %s\n",p,PtoK(p),k,KtoP(k))
		} else{
			fmt.Println("Cannot conv. Use temp length or weight")
			return 
		}
	}
}