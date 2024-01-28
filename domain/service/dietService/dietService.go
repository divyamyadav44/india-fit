package dietService

import (
	"india-fit/domain/model"
	"india-fit/domain/repo/DietRepo"
	"log"
)

func ListDiet(id int64) (listOfDaysAndMeals []model.ListofDaysAndMeals, err error) {

	listOfDaysAndMeals, err = DietRepo.ListDiet(id)
	if err != nil {
		log.Println("error in getting user details", err)
		return
	}
	return listOfDaysAndMeals, err
}
