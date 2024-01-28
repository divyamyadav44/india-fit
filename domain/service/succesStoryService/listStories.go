package succesStoryService

import (
	"india-fit/domain/model"
	"india-fit/domain/repo/succesStoryRepo"
	"log"
)

func ListStories(count int64) (story []model.SuccesStory, err error) {

	story, err = succesStoryRepo.ListStories(count)
	if err != nil {
		log.Println("error in getting list of success stories", err)
		return
	}
	return story, err
}
