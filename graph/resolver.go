package graph

import (
	database "github.com/gcesario203/fullcycle-go-graphq-/internal/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
