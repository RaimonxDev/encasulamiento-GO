package course

import "fmt"

type course struct {
	// todos los campos estan privados
	name    string
	price   float64
	isFree  bool
	userID  []uint
	classes map[uint]string
}

//Setter and Getter

func (c *course) SetName(name string) {
	c.name = name
}
func (c *course) Name() string {
	return c.name
}

func (c *course) SetPrice(price float64) {
	c.price = price
}
func (c *course) Price() float64 {
	return c.price
}

func (c *course) SetIsFree(isFree bool) {
	c.isFree = isFree
}
func (c course) IsFree() bool {
	return c.isFree
}

func (c *course) SetUserId(userId []uint) {
	c.userID = userId
}

func (c *course) UserId() []uint {
	return c.userID
}

func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

func (c *course) Classes() map[uint]string {
	return c.classes
}

//New Function constructura buena practia utilizar el nombre new
func New(name string, price float64, isFree bool) *course {

	if price == 0 {
		price = 30.50
	}

	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

// Private Method
func (c *course) printClasses() {
	text := "Las clases son las siguientes: "
	for _, class := range c.classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
