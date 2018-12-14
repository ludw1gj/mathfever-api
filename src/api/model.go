package api

import "html/template"

// Category holds information about a category of calculations.
type Category struct {
	Name         string        `json:"name"`         // name of the category
	Slug         string        `json:"slug"`         // slug of the category
	Description  string        `json:"description"`  // description of the category
	Calculations []Calculation `json:"calculations"` // the calculations belonging to the category
}

// Calculation holds information about a calculation
type Calculation struct {
	Name         string             `json:"name"`         // name of the calculation
	Slug         string             `json:"slug"`         // slug of the calculation
	Description  string             `json:"description"`  // description of the calculation
	InputInfo    []CalculationInput `json:"inputInfo"`    // contains the names of the inputs and their json tags
	Example      template.HTML      `json:"example"`      // example of output of calculation calculation function
	Category     string             `json:"category"`     // the name of the category the calculation belongs to
	CategorySlug string             `json:"categorySlug"` // the slug of the category the calculation belongs to
}

// CalculationInput describes an input name and input json field of a calculation.
type CalculationInput struct {
	Name string `json:"name"` // name of input eg. Side A
	Tag  string `json:"tag"`  // json field of Name eg. side_a
}
