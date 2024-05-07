package main

func main() {

}

func countStudents(students []int, sandwiches []int) int {
	counter := 0
	for len(sandwiches) > 0 {
		if counter == len(students) {
			return len(students)
		}
		if sandwiches[0] == students[0] {
			sandwiches = sandwiches[1:]
			students = students[1:]
			counter = 0
		} else {
			temp := students[0]
			students = students[1:]
			students = append(students, temp)
			counter++
		}
	}
	return 0
}
