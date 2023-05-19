package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}

	fmt.Println(courseName)

	// Add value
	courseName = append(courseName, "C", "C#", "HTML", "CSS", "JavaScript")
	fmt.Println(courseName)

	// Delete value
	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
