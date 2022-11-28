package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	// score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	// your code here
	var code rune
	for i := 0; i < len(s.name); i++ {
		if s.name[i] >= 'a' && s.name[i] <= 'z' {
			code = rune((((s.name[i] - 'a') + 3) % 26) + 'a')
			nameEncode += string(code)
		}
	}
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	// your code here
	for _, code := range s.nameEncode {
		if code >= 'a' && code <= 'z' {
			code = (((code-'a')-3)%26 + 'a')
		}
		if code < 'a' {
			code = (code + 26)
		} else if code > 'z' {
			code = (code - 26)
		}
		nameDecode += string(code)
	}
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
