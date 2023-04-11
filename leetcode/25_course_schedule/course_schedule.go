package course_schedule

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	courses := make(map[int][]int)

	for _, prerequisite := range prerequisites {
		course, prerequisite := prerequisite[0], prerequisite[1]
		courses[course] = append(courses[course], prerequisite)
	}

	var courseHasPrerequisiteCycle func(
		course int,
		courses map[int][]int,
		visited map[int]bool,
	) bool
	courseHasPrerequisiteCycle = func(
		course int,
		courses map[int][]int,
		visited map[int]bool,
	) bool {
		if visited[course] {
			return false
		}

		visited[course] = true
		for _, prerequisite := range courses[course] {
			if !courseHasPrerequisiteCycle(prerequisite, courses, visited) {
				return false
			}
		}
		visited[course] = false
		courses[course] = []int{}

		return true
	}

	for course := range courses {
		if !courseHasPrerequisiteCycle(course, courses, make(map[int]bool)) {
			return false
		}
	}

	return true
}
