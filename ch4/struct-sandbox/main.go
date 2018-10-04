package main

import "time"

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func main() {

}
