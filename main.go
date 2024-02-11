package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {

}

func tesStruct() {
	var mawar kosEkslusif = kosEkslusif{800, true, penyewaKos{"naura"}}
	mawar.biaya = 1000
	var melati kosMurahTanpaKTP = kosMurahTanpaKTP{isian: false, biaya: 200}
	var lily kosMurah = kosMurah{200, false, penyewaKos{"Dea"}}
	var kosMelati = struct {
		biaya   int
		isian   bool
		penyewa penyewaKos
	}{200, false, penyewaKos{"amia"}}
	fmt.Println(mawar.biaya, melati.biaya, mawar.nama, lily.penyewa.nama, kosMelati.penyewa.nama)
	fmt.Println(mawar.biayaBulanan(6))
	fmt.Println(cekBudget(lily, 2, 400))
}

func cekBudget(k kos, month int, budget int) bool {
	if budget < k.biayaBulanan(month) {
		return false
	} else {
		return true
	}

}

type kos interface {
	biayaBulanan(month int) int
}

type kosEkslusif struct {
	biaya int
	isian bool
	penyewaKos
}

func (k kosEkslusif) biayaBulanan(month int) int {
	return k.biaya * month
}

type kosMurah struct {
	biaya   int
	isian   bool
	penyewa penyewaKos
}

func (k kosMurah) biayaBulanan(month int) int {
	return k.biaya * month
}

type kosMurahTanpaKTP struct {
	biaya int
	isian bool
}

func (k kosMurahTanpaKTP) biayaBulanan(month int) int {
	return k.biaya * month
}

type penyewaKos struct {
	nama string
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

func tesString() {

	fmt.Println("without stringbuilder")
	var stringSlices = []string{"c", "o", "b", "a"}
	var stringSliced = ""
	for i := len(stringSlices) - 1; i >= 0; i-- {
		fmt.Println(stringSlices[i])
		stringSliced += stringSlices[i]
		fmt.Println(stringSliced)
	}
	fmt.Println("")
	fmt.Println("stringbuilder")
	var strBuild strings.Builder
	for i := range stringSlices {
		strBuild.WriteString(stringSlices[i])
	}
	var strBuilt = strBuild.String()
	fmt.Println(strBuilt)
	var string1 = "untuk kamu"
	strBuild.Reset()
	strBuild.WriteString(string1)
	strBuilt2 := strings.Split(strBuild.String(), " ")
	strBuild.Reset()
	for i := range strBuilt2 {
		strBuild.WriteString(strBuilt2[i])
	}
	strBuilt3 := strBuild.String()
	fmt.Printf("%v", strBuilt3)
}

func tesArraySliceMap() time.Duration {
	t0 := time.Now()
	fmt.Println(t0)
	//array
	var array1 [4]int = [4]int{1, 2, 3, 4}
	array2 := [4]int{100, 90, 80, 70}
	//if array1 tidak sama dengan null atau array2 tidak sama dengan null
	if array1 != [4]int{} || array2 != [4]int{} {
		fmt.Println("bisa")
	}
	i := array2
	fmt.Printf("cobak %v \n", i)
	for i := len(array2) - 1; i >= 0; i-- {
		fmt.Printf("* %v * \n", array2[i])
	}

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
	return time.Since(t0)
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
