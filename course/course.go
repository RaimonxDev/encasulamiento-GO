package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserID  []uint
	Classes map[uint]string
}

//Podemos llamar los desde otros archivos
// Deben de pertenecer a el mismo package
func (c Course) PrintClasses() {
	text := "Las clases son las siguientes: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

/* Para poder actualizar un campo debemos de recibir un puntero de la interface a trabajar */
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
