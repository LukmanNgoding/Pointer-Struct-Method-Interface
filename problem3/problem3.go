package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 { // menghasilkan hasil rata-rata dar 6 skor
	// your code here
	var hasil float64
	for i := 0; i < 6; i++ {
		hasil += float64(s.score[i])

	}
	//rumus rata-rata
	total := hasil / float64(len(s.score))

	return total
}

func (s Student) Min() (min int, name string) {
	// your code here
	min = s.score[0]         // min ini data pertama dari array [0]
	name = s.name[0]         // nama ini data pertama dari array [0]
	for i := 0; i < 6; i++ { // mengulang sebanyak 6 kali
		if s.score[i] < min { // jika array s.score 1 kurang dari min maka
			min = s.score[i] // min adalah s.score [i]
			name = s.name[i] // nama adalah s.name[i]
		}

	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	// your code here
	max = s.score[0]         // max ini data pertama dari array [0]
	name = s.name[0]         // nama ini data pertama dari array [0]
	for i := 0; i < 6; i++ { // mengulang sebanyak 6 kali
		if s.score[i] > max { // jika array s.score 1 kurang dari min maka
			max = s.score[i] // max adalah s.score [i]
			name = s.name[i] // nama adalah s.name[i]
		}

	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 1; i < 7; i++ { // 6 nama dan 6 siswa input
		var name string
		fmt.Print("Input " + fmt.Sprint(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name) // berapa di dalam array pada struk student
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score) // berapa di dalam array pada struk student
	} //a = = {[nama1, nama2, nama3,...][nilai1, nilai2, ...]}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
