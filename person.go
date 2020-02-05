package main

// Person struct stores person age
type Person struct {
	Age uint8
}

// TwiceOlderExists looks is there Person age
// in slice that is at least exactly as old as other persons
// Assimprotic compexility of this function is O(n)
func TwiceOlderExists(persons []Person) bool {

	ageMap := make(map[uint8]uint8)

	for _, person := range persons {
		ageMap[person.Age] = 0
	}

	for _, person := range persons {
		_, exists := ageMap[person.Age/2]

		if exists {
			return true
		}

	}

	return false
}

// MinTwiceOlderExists looks is there Person age
// in slice that is at least twice as old as other persons
// Assimprotic compexility of this function is O(n)
func MinTwiceOlderExists(persons []Person) bool {

	if len(persons) < 2 {
		return false
	}

	maxYearsIndex := 0

	// find largest number in slice
	for index, person := range persons {
		if person.Age > persons[maxYearsIndex].Age {
			maxYearsIndex = index
		}
	}

	for index, person := range persons {
		// if any number in slice doubled is bigger then max. num
		// there is no number that is at least twice as big as other numbers
		if index != maxYearsIndex && persons[maxYearsIndex].Age < person.Age*2 {
			return false
		}
	}

	// otherwise return true
	return true
}
