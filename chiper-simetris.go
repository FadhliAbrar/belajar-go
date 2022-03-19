package main

import (
	_ "belajar-golang/database"
	"errors"
	"fmt"
	"strings"
)

type arrofstr struct {
	alfabet []string
}

type cryptograph struct {
	Crypt string
	key   int
}

func (slice *arrofstr) sumIndexAndKey(str string, key int) (int, error) {
	for i, v := range slice.alfabet {
		if str == v {
			return i + key, nil
		}
	}
	return 1288, errors.New("Char tidak ditemukan atau bukan alfabet")
}

func cryptEncypt(str string, key int) *cryptograph {
	myAlfabet := arrofstr{
		alfabet: []string{
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		},
	}
	mySlice := strings.Split(str, "")
	var takeAChar []int
	var takeAString []string
	for _, v := range mySlice {
		val, err := myAlfabet.sumIndexAndKey(strings.ToUpper(v), key)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			takeAChar = append(takeAChar, val)
		}
	}
	for _, v := range takeAChar {
		
		var temp int
		if v >= 26 {
			secTemp := v / 26

			if v*secTemp%26 == 0 {
				temp = v % 26
				takeAString = append(takeAString, myAlfabet.alfabet[temp])
				temp = 0
				continue
			}
			if v*secTemp%26 > 0 {
				temp = v % 26		
				takeAString = append(takeAString, myAlfabet.alfabet[temp])
				temp = 0
				continue
			}
		}
		takeAString = append(takeAString, myAlfabet.alfabet[v])
	}

	return &cryptograph{strings.Join(takeAString, ""), key}
}

func (val *cryptograph) decrypt() (string, interface{}) {
	var key int = val.key
	// menampung string
	myMapOfCrypt := strings.Split(val.Crypt, "")
	//menampung index
	var tempIdx []int
	myAlfabet := arrofstr{
		alfabet: []string{
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		},
	}
	for _, v := range myMapOfCrypt {
		for idx, val := range myAlfabet.alfabet {
			if v == val {
				tempIdx = append(tempIdx, idx)
				break
			}
		}
	}
	var encrypted []string
	// var kelipatanIndex int = key / 26
	for _, v := range tempIdx {
		tempVal := key % 26
		inject := v - tempVal
		if inject < 0 {
			encrypted = append(encrypted, myAlfabet.alfabet[inject+26])
		} else {
			encrypted = append(encrypted, myAlfabet.alfabet[inject])
		}
	}
	return strings.Join(encrypted, ""), val.Crypt
}

func main() {
	pesan := "apassword"
	key := 100
	Hasil, crypt := cryptEncypt(pesan, key).decrypt()

	fmt.Println("Private Message:", crypt)
	fmt.Println("With key:", key)
	fmt.Println("Message:", Hasil)
}
