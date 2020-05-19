package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"graphql-tutorial/src/graphqlconfig"
	"graphql-tutorial/src/types"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	tutorials := types.Populate()
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"tutorial": &graphql.Field{
			Type:        graphqlconfig.TutorialType,
			Description: "Get tutorial by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"].(int)
				if ok {
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							return tutorial, nil
						}
					}
				}
				return nil, errors.New("404 Tutorial not found")
			},
		},

		"list": &graphql.Field{
			Type:        graphql.NewList(graphqlconfig.TutorialType),
			Description: "Get List of Tutorials",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// data query to fetch
	queryList := `
		{
			list {
				ID
				Title
				Comments {
					Body
				}
				Author {
					Name
					Tutorials
				}
			}
		}
	`
	queryTutorial_1 := `
		{
			tutorial(id:1) {
				ID
				Title
				Author {
					ID
					Name
					Tutorials 
				}
				Comments {
					Body
				}
			}
		}
	`

	paramsList := graphql.Params{
		Schema:        schema,
		RequestString: queryList,
	}
	r := graphql.Do(paramsList)

	if len(r.Errors) > 0 {
		log.Fatalf("failed to operate")
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	paramsTutorial_1 := graphql.Params{
		Schema:        schema,
		RequestString: queryTutorial_1,
	}
	resolveQuery := graphql.Do(paramsTutorial_1)

	if len(resolveQuery.Errors) > 0 {
		log.Fatalf("failed to operate")
	}
	rJSON, _ = json.Marshal(resolveQuery)
	fmt.Printf("%s \n", rJSON)

	queryWrongAssertion := `
		{
			tutorial(id:100) {
				ID
				Title
				Comments {
					Body
				}
				Author {
					ID
					Name
					Tutorials
				}
			}
		}
	`

	requestParamsWrong := graphql.Params{
		Schema:        schema,
		RequestString: queryWrongAssertion,
	}

	resParamsWrong := graphql.Do(requestParamsWrong)

	if resParamsWrong.HasErrors() == true {
		fmt.Println("404, not found")
	} else {
		rJSON, _ = json.Marshal(resParamsWrong)
		fmt.Printf("%s \n", rJSON)
	}
}
