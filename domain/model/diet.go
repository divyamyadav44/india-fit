package model

import "time"

type ListofDaysAndMeals struct {
	Id                 int64     `json:"id"`
	CurrentDateAndTime time.Time `json:"current_date_and_time"`
	Meals              []Meal    `json:"meals"`
}

type Meal struct {
	Id                int64     `json:"id"`
	UserId            int64     `json:"user_id"`
	MealTime          time.Time `json:"meal_time"`
	MealName          string    `json:"meal_name"`
	Indredient        []string  `json:"indredient"`
	Steps             string    `json:"steps"`
	CreatedByDietcian int64     `json:"created_by_dietcian"`
}
