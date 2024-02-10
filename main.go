package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	//array
	var array1 [4]int = [4]int{1, 2, 3, 4}
	array2 := [4]int{1, 2, 3, 4}
	//if array1 tidak sama dengan null atau array2 tidak sama dengan null
	if array1 != [4]int{} || array2 != [4]int{} {
		fmt.Println("bisa")
	}
	i := array2
	fmt.Printf("cobak %v \n", i)

	//slice
	var slice1 []int = []int{1, 2, 5}
	var slice2 []int = []int{5, 6, 7, 8}
	slice1 = append(slice1, slice2...)
	var slice3 []int = make([]int, 3, 10)
	fmt.Println(slice1)
	fmt.Println(len(slice3), cap(slice3))
	fmt.Println(slice1[len(slice1)-1]) //nilai terakhir

	//map
	var map1 = make(map[string]int)
	map1 = map[string]int{"Adam": 25, "Munaroh": 15}
	fmt.Println(map1)
	var age, isThereSisil = map1["Sisil"]
	if isThereSisil {
		fmt.Printf("Umur sisil adalah %v \n", age)
	} else {
		fmt.Println("Tidak ada Sisil")
	}
	for name := range map1 {
		fmt.Printf("Nama: %v \n", name)
	}
}

func tesDataType() {
	var angka1 int = 10
	var angka2 int = 2
	angka3 := 10.5
	var string3, string4 string = "Hello", "World"
	var1, var2 := "qwerty", 2
	fmt.Println(angka1 / angka2)
	fmt.Println(angka1 % angka2)
	fmt.Println(angka3 / float64(angka1))
	fmt.Println(string3 + " " + string4 + " " + fmt.Sprint(angka1))
	fmt.Println(var1, var2)
	var string1 string = `Hello
there`
	var string2 string = "Hello thereΩ"
	fmt.Println(string1)
	fmt.Println(utf8.RuneCountInString(string2))
}

func tesFunction() {
	var var5, var6, err = meng(10, 2)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Hasilnya adalah %v dan %v", var5, var6)
	}

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case var6 == 0:
		fmt.Println("okegas")
	default:
		fmt.Printf("The results are %v and %v", var5, var6)
	}

	switch var6 {
	case 10:
		fmt.Println("10a")
	case 11:
		fmt.Println("11a")
	default:
		fmt.Println(var6)
	}
}
func meng(pembilang int, pembagi int) (int, int, error) {
	var err error
	if pembagi == 0 {
		err = errors.New("gabisa bang")
		return 0, 0, err
	}
	var5 := pembilang / pembagi
	var6 := pembilang % pembagi
	return var5, var6, err
}
