package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	database "yourJube_API/db"
	"yourJube_API/graph/generated"
	"yourJube_API/graph/model"

	"github.com/jinzhu/gorm"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	NewUser := model.User{
		Name:              input.Name,
		Email:             input.Email,
		ProfilePict:       input.ProfilePict,
		ChannelBackground: input.ChannelBackground,
		Premium:           input.Premium,
	}
	db.Create(&NewUser)
	return &NewUser, nil
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input model.NewVideo) (*model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Set("gorm:auto_preload", true).Find(&r.videos)
	return r.videos, nil
}

func (r *queryResolver) FindUser(ctx context.Context, email string) (*model.User, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	if err := db.Where("email = ?", email).First(&r.user).Error; gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return &r.user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Find(&r.users)
	return r.users, nil
}

func (r *queryResolver) CommentsByVideoID(ctx context.Context, id string) ([]*model.Comment, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Set("gorm:auto_preload", true).Find(&r.comments)
	return r.comments, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *commentResolver) CreatedAt(ctx context.Context, obj *model.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *commentResolver) UpdatedAt(ctx context.Context, obj *model.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *locationResolver) CreatedAt(ctx context.Context, obj *model.Location) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *locationResolver) UpdatedAt(ctx context.Context, obj *model.Location) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *videoResolver) CreatedAt(ctx context.Context, obj *model.Video) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *videoResolver) UpdatedAt(ctx context.Context, obj *model.Video) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

type commentResolver struct{ *Resolver }
type locationResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type videoResolver struct{ *Resolver }
