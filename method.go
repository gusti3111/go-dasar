package main

import "fmt"

type student struct {
	id    int
	name  string
	class int
}

// creating Method hai

func (s student) hai() {
	fmt.Println("Hallo saya ", s.name)
	if s.class == 12 && s.name == "george" {
		fmt.Println("ya nama saya", s.name, "saya kelas", s.class)
	} else if s.class != 12 {
		fmt.Println("ya saya dari kelas", s.class, "nama saya ", s.name)

	} else {
		fmt.Println("saya bukan ", s.name, "dan saya bukan dari kelas", s.class)
	}

}

// creating method jurusan
// type jurusan struct{

// }

func main() {
	var s1 = student{name: "george"}
	fmt.Println("Nama kamu siapa ?")
	fmt.Scanln(&s1.name)

	fmt.Println("Kamu kelas berapa ?")
	fmt.Scanln(&s1.class)

	s1.hai()

}
