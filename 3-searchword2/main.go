package main

import (
	"fmt"
	"strings"
)

func main() {
	aa:=[]string{"Batiar","Afas", "Ari", "Afasari", "Batiar afas", "Afasari Batiara",
	"Rahma Mulia"}
	bb:="Bat"
	//pri := countsearch(aa,b)
	//fmt.Println(pri)
	//pril := listsearch(aa,b)
	//for _, r := range pril{
	//	fmt.Println(r)
	//}
	a,b := wordsearch(aa,bb)
	fmt.Println(a)
	for _, r := range b{
		fmt.Println(r)
	}

}

func wordsearch(aa []string, b string) (int,[]string){
	var list []string
	var count int
	for _, r := range aa{
		ret := strings.Split(r," ")
		for _, ret2 := range ret{
			if strings.Contains(ret2, b) {
				list = append(list, r)
				count++
			}
		}
	}
	return count,list
}
//func countsearch(aa []string, b string) int{
//	var count int
//		for _, r := range aa{
//			ret := strings.Split(r," ")
//			for _, ret2 := range ret{
//				if strings.Contains(ret2, b) {
//					count++
//				}
//			}
//		}
//	return count
//}
//
//func listsearch(aa []string, b string) []string{
//	var list []string
//	for _, r := range aa{
//		ret := strings.Split(r," ")
//		for _, ret2 := range ret{
//			if strings.Contains(ret2, b) {
//				list = append(list, r)
//			}
//		}
//	}
//	return list
//}