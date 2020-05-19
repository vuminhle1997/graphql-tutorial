package types

func Populate() []Tutorial {
	author := &Author{
		Name:      "Elliot Forbes",
		Tutorials: []int{1},
	}
	tutorial := Tutorial{
		ID:     1,
		Title:  "Go GraphQl Tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{
				Body: "GoLang rocks",
			},
		},
	}

	dexter := &Author{
		Name:      "Dexter Berg",
		Tutorials: []int{5},
	}
	node := Tutorial{
		ID:     2,
		Title:  "NodeJS Tutorial",
		Author: *dexter,
		Comments: []Comment{
			Comment{
				Body: "NodeJS is ez",
			},
			Comment{
				Body: "NodeJS is slow",
			},
		},
	}
	flutter := Tutorial{
		ID:     3,
		Title:  "Flutter Tutorial",
		Author: *dexter,
		Comments: []Comment{
			Comment{
				Body: "Flutter is from Google",
			},
			Comment{
				Body: "It is easy to learn",
			},
		},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	tutorials = append(tutorials, node)
	tutorials = append(tutorials, flutter)

	return tutorials
}
