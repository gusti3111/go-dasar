package main

import "fmt"

type student struct {
	Id    int
	name  string
	class int
}
type areaOfRectangle struct {
	length float64
	width  float64
}
type partOfBody struct {
	eye, nose, hand, leg int
}

//Add Method to Struct Type
func (s student) murid() string {
	fmt.Println(s.name, s.class)
	return "________"

}

func main() {
	// Creating Nested Struct
	var body = partOfBody{
		eye:  2,
		nose: 1,
		hand: 2,
		leg:  2,
	}

	// fmt.Println("berapakah jumlah mata ?")
	// created variabel eyes
	// value eyes = 2 dari body.eye
	// if you wanna print value of  eyes : *eyes
	// untuk cetak alamat memori

	var eyes *int = &body.eye
	fmt.Println("Berapa jumlah Mata ?")
	fmt.Scan(&body.eye)

	if *eyes < 2 && *eyes > 2 {
		fmt.Println("salah")

	} else {
		fmt.Println("benar")
	}
	// another way to filling struct paramater
	// kalo membungkusnya dengan sebuah variabel baru

	var newstudent student
	newstudent.name = "john doe"
	newstudent.Id = 001
	newstudent.class = 10

	newstudent.murid()

	// Creating a Struct Instance Using a Struct Literal

	area := areaOfRectangle{20, 5}
	fmt.Println(area.length * area.width)
	//

	//Struct Instantiation Using Pointer Address Operator
	// valuedalam pointer gak bolehkurang dari yang di deklarasikan di struct
	var anggotaTubuh = &partOfBody{2, 1, 2, 2}
	fmt.Println(anggotaTubuh.eye, anggotaTubuh.hand, anggotaTubuh.leg, anggotaTubuh.nose)
	//

	var newAreaRec areaOfRectangle
	newAreaRec.length = 10
	newAreaRec.width = 20.6

	newArea := newAreaRec.length * newAreaRec.width
	fmt.Println("luas : ", newArea)

	fmt.Println("nama murid barunya adalah: ", newstudent.name)
	fmt.Println("kelas si murid baru adalah ", newstudent.class)

}
