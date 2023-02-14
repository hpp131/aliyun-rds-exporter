package main

func main() {
	s1 := []int{11, 22, 33}
	s2 := []int{44, 55, 66}
	s1 = append(s1, s2...)

}
