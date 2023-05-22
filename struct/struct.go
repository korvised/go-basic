package main

import "fmt"

type employee struct {
	id   string
	name string
	tel  string
}

func main() {
	staff := employee{
		id:   "101",
		name: "Korvised",
		tel:  "20 98 866 447",
	}

	fmt.Println("employee =", staff)

	//! struct with array
	employeeList := [3]employee{}
	employeeList[0] = employee{
		id:   "101",
		name: "Long",
		tel:  "20 98 866 447",
	}

	employeeList[1] = employee{
		id:   "102",
		name: "Tim",
		tel:  "20 98 866 447",
	}

	employeeList[2] = employee{
		id:   "103",
		name: "Kim",
		tel:  "20 98 866 447",
	}

	fmt.Println("employeeList =", employeeList)

	//! struct with slice
	employees := []employee{}

	employee1 := employee{
		id:   "101",
		name: "Long",
		tel:  "20 98 866 447",
	}

	employee2 := employee{
		id:   "102",
		name: "Tim",
		tel:  "20 98 866 447",
	}

	employee3 := employee{
		id:   "103",
		name: "Kim",
		tel:  "20 98 866 447",
	}

	employees = append(employeeList[:], employee1)
	employees = append(employeeList[:], employee2)
	employees = append(employeeList[:], employee3)

	fmt.Println("employees =", employees)
}
