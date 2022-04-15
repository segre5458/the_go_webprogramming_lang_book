package main

import "fmt"

type Pond float64
type Kilogram float64

func PtoK(p Pond) Kilogram{return Kilogram(p/2.2046)}
func KtoP(k Kilogram) Pond{return Pond(k*2.2046)}

func (p Pond) String() string {return fmt.Sprintf("%gPond",p)}
func (k Kilogram) String() string {return fmt.Sprintf("%gKg",k)}