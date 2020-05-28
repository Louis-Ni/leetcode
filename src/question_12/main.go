package main

import "fmt"
func main() {
	str := intToRoman(809)
	fmt.Println(str)
}
func intToRoman(num int) string {
	one := []string{"","I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"} //1,2,3,4,5,6,7,8,9
	ten := []string{"","X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"} //10,20,30,40,50,60,70,80,90
	hundred := []string{"","C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"} //100,200,300,400,500,600,700,800,900
	thousand := []string{"","M", "MM", "MMM"}//1000,2000,3000
	if num < 10 {
		return one[num]
	}else if num >= 10 && num <= 99 {
		t := num / 10
		o := num % 10
		return ten[t] + one [o]
	}else if num >= 100 && num <= 999{
		h := num / 100
		t := num / 10 % 10
		o := num % 10
		return hundred[h ] + ten [t] + one [o]
	}else {
		th := num / 1000
		h := num / 100 % 10
		t := num / 10 % 10
		o := num % 10
		return thousand[th] + hundred[h] + ten[t] + one [o]
	}
}