package dto

import (
	"time"
	"timeCardSimple/app/domain/employee"
)

func Employees(v interface{}) interface{} {
	employees := v.([]*employee.Employee)

	result := make([]interface{}, len(employees))
	for i, employee := range employees {
		result[i] = Employee(employee)
	}

	return result
}

func Employee(v interface{}) interface{} {
	e := v.(*employee.Employee)

	id, _ := Transform(e.ID())

	return &struct {
		ID        interface{} `json:"id"`
		FirstName string      `json:"first_name:"`
		LastName  string      `json:"last_name"`
		Email     string      `json:"email"`
		CreatedAt time.Time   `json:"created_at"`
		UpdatedAt time.Time   `json:"updated_at"`
	}{
		id,
		e.FirstName(),
		e.LastName(),
		e.Email(),
		e.CreatedAt(),
		e.UpdatedAt(),
	}
}

// type ResponseEmployee struct {
// 	ID        i    `json:"id"`
// 	FirstName string    `json:"first_name"`
// 	LastName  string    `json:"last_name"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }
