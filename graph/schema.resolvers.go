package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen-echo-sample/graph/generated"
	"gqlgen-echo-sample/graph/model"
	"time"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	task := model.Task{
		Title:     input.Title,
		Note:      input.Note,
		Completed: 0,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
	r.DB.Create(&task)

	return &task, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	account := model.Account{
		Name: input.Name,
		// Password:  password,
		Password:  input.Password,
		Completed: 0,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
	r.DB.Create(&account)

	return &account, nil
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	tasks := []*model.Task{}

	r.DB.Find(&tasks)

	return tasks, nil
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*model.Account, error) {
	accounts := []*model.Account{}

	r.DB.Find(&accounts)

	return accounts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
