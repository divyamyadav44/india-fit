package succesStoryService

import (
	"india-fit/domain/model"
	"india-fit/domain/repo/succesStoryRepo"
	"log"
)

func GetStory(id int64) (story model.SuccesStory, err error) {

	story, err = succesStoryRepo.GetStory(id)
	if err != nil {
		log.Println("error in getting user details", err)
		return
	}
	return story, err
}
