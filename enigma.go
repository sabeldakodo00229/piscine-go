package piscine

// import "fmt"

func Enigma(a ***int, b *int, c *******int, d ****int) {
	Pointerc := *******c
	*******c = ***a
	Pointerd := ****d
	****d = Pointerc
	Pointerb := *b
	*b = Pointerd
	***a = Pointerb
}

/*
func main() {
	x := 5
	y := &x
	z := &y
	a := &z

	w := 2
	b := &w

	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	Enigma(a, b, c, d)

	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}
*/