package types

type Tutorial struct {
	ID       int
	Title    string
	Comments []Comment
	Author   Author
}
