package models

type Todo struct {
	Id        int
	Text      string
	Completed bool
}

// just for demo
var TodoSlice = []*Todo{
	{
		Id:        1,
		Text:      "breakfast",
		Completed: true,
	},
	{
		Id:        2,
		Text:      "lunch",
		Completed: false,
	},
	{
		Id:        3,
		Text:      "dinner",
		Completed: false,
	},
}
