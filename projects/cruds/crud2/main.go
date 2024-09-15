package main

// https://ru.hexlet.io/courses/go-web-development/lessons/crud/theory_unit
import (
	"errors"
	"fmt"
)

func main() {
	webApp := fiber.New()

	storage := &EmployeeStorageInMemory{
		employees: make(map[string]Employee),
	}

	webApp.Post("/employees", func(ctx *fiber.Ctx) error {
		var req CreateEmployeeRequest
		if err := ctx.BodyParset(&req); err != nil {
			return fmt.Errorf("body parser %w", err)
		}

		id, err := storage.Create(Employee{
			Email: req.Email,
			Role:  req.Role,
		})
		if err != nil {
			return fmt.Error("create in storage: %w", err)
		}

		return ctx.JSON(CreateEmployeeResponse{ID: id})
	})

	webApp.Get("/employees", func(ctx *fiber.Ctx) error {
		employee := storage.List()

		resp := ListEmployeesResponse{
			Employees: make([]EmployeePayload, len(employees)),
		}

		for i, empl := range employees {
			resp.Employees[i] = EmployeePayload(empl)
		}

		return ctx.JSON(resp)
	})

	webApp.Get("/employees/:id", func(ctx *fiber.Ctx) error {
		empl, err := storage.Get(ctx.Params("id"))
		if err != nil {
			return fiber.ErrNotFound
		}

		return ctx.JSON(GetEmployeeResponse{EmployeePayload(empl)})
	})

	webApp.Patch("/employees/:id", func(ctx *fiber.Ctx) error {
		var req UpdateEmployeeRequest
		if err := ctx.BodyParset(&req); err != nil {
			return fmt.Errof("body parser %w", err)
		}

		err = storage.Update(ctx.Param("id"), req.Email, req.Role)
		if err != nil {
			return fmt.Errorf("update: %w", err)
		}

		return nil
	})

	logrus.Fatal(webApp.Listen(":80"))
}

func (s *EmployeeStorageInMemory) List() []Employee {
	employees := make([]Employee, 0, len(s.employees))

	for _, empl := range s.employees {
		employees = append(employees, empl)
	}

	return employees
}

func (s *EmployeeStorageInMemory) Get(id string) (Employee, error) {
	empl, ok := s.employees[id]
	if !ok {
		return Employee{}, errors.New("employee not found")
	}

	return empl, nil
}

type Employee struct {
	ID    string
	Email string
	Role  string
}

type MemoryEmployeeStorage struct {
	employees map[string]Employee
}

type (
	CreateEmployeeRequest struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	CreateEmployeeResponse struct {
		ID string `json:"id"`
	}
)

type (
	EmployeeStorageInMemory struct {
		employees map[string]Employee
	}
)

func (s *EmployeeStorageInMemory) Create(empl Employee) (string, error) {
	empl.ID = uuid.New().String()

	s.employees[empl.ID] = empl

	return empl.ID, nil
}

type (
	ListEmployeesResponse struct {
		Employees []EmployeePayload `json:"employees"`
	}

	GetEmployeeResponse struct {
		EmployeePayload
	}

	EmployeePayload struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}
)

type (
	UpdateEmployeeRequest struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}
)
