package maps

import (
	"fmt"
)

func main()  {
	m:= map[string]string {
		"name":"ccmouse",
		"course":"golang",
		"site":"immoc",
		"quality":"notbad",
	}
	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string] int // m3==nil
	fmt.Println(m,m2,m3)

	fmt.Println("Traversing map")
	for k,v := range m {
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	courseName,ok := m["couse"]
	fmt.Println(courseName,ok)
	if courseName,ok := m["coursde"]; ok{
		fmt.Println(courseName)
	}else {
		fmt.Println("key does not exist.")
	}

	fmt.Println("Deleting values")
	name,ok := m["name"]
	fmt.Println(name,ok)

	delete(m,"name")
	name,ok = m["name"]
	fmt.Println(name,ok)
}