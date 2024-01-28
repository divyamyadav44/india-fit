package succesStoryRepo

import "india-fit/domain/model"

func GetStory(id int64) (story model.SuccesStory, err error) {

	story = model.SuccesStory{
		Id:          1,
		SuccesStory: "this is a succes story",
		ImageLink:   "https://gfbdb//link",
	}
	return
}
