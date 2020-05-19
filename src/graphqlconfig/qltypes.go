package graphqlconfig

import (
	"github.com/graphql-go/graphql"
)

var CommentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var AuthorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var TutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Author": &graphql.Field{
				Type: AuthorType,
			},
			"Comments": &graphql.Field{
				Type: graphql.NewList(CommentType),
			},
		},
	},
)
