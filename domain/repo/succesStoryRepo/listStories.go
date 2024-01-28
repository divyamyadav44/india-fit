package succesStoryRepo

import "india-fit/domain/model"

func ListStories(count int64) (stories []model.SuccesStory, err error) {

	stories = []model.SuccesStory{
		{
			Id:          1,
			SuccesStory: "this is a 1 succes story",
			ImageLink:   "https://gfbdb//link",
		},
		{
			Id:          2,
			SuccesStory: "this is a 2 succes story",
			ImageLink:   "https://gfbdb//link",
		},
		{
			Id:          3,
			SuccesStory: "this is a 3 succes story",
			ImageLink:   "https://gfbdb//link",
		},
		{
			Id:          4,
			SuccesStory: "this is a 4 succes story",
			ImageLink:   "https://gfbdb//link",
		},
		{
			Id:          5,
			SuccesStory: "this is a 5 succes story",
			ImageLink:   "https://gfbdb//link",
		},
	}

	return
}
