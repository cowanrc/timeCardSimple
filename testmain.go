package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"
// )

// //Employee Structure that saves time of clockin and clockout for each employee
// type Employee struct {
// 	Name       string
// 	EmployeeId int
// 	ClockIn    time.Time
// 	ClockOut   time.Time
// 	TotalTime  string
// }

// //timeCard map that relates to struct
// var TimeCard map[string]*Employee

// var currentEmployee Employee

// func main() {
// 	TimeCard = make(map[string]*Employee)
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("\nPlease enter your name:")
// 	currentEmployee.Name, _ = reader.ReadString('\n')
// 	addEmployeeToSystem(currentEmployee.Name)
// 	// timeCard
// 	// log.Printf("Time Card name is %s ", timeCard.Name)
// 	for _, v := range TimeCard {
// 		log.Printf(v.Name)
// 		// if v.Name != currentEmployee.Name {
// 		// 	log.Printf("Welcome to our time card application, %s", currentEmployee.Name)
// 		// }
// 	}
