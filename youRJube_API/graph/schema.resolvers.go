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

func (r *mutationResolver) CreateFeedBackOnComment(ctx context.Context, input model.NewLikeComment) (*model.LikeComment, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	NewLikeComment := model.LikeComment{
		UserID:    input.UserID,
		CommentID: input.CommentID,
		Status:    input.Status,
	}
	db.Create(&NewLikeComment)
	return &NewLikeComment, nil
}

func (r *mutationResolver) CreateFeedBackOnVideo(ctx context.Context, input model.NewLikeVideo) (*model.LikeVideo, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	NewLikeVideo := model.LikeVideo{
		UserID:  input.UserID,
		VideoID: input.VideoID,
		Status:  input.Status,
	}
	db.Create(&NewLikeVideo)
	return &NewLikeVideo, nil
}

func (r *mutationResolver) CreateFeedBackOnPost(ctx context.Context, input model.NewLikePost) (*model.LikePost, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	NewLikePost := model.LikePost{
		UserID: input.UserID,
		PostID: input.PostID,
		Status: input.Status,
	}
	db.Create(&NewLikePost)
	return &NewLikePost, nil
}

func (r *mutationResolver) CreateSubscriber(ctx context.Context, input model.NewSubscriber) (*model.Subscriber, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	NewSubscriber := model.Subscriber{
		TargetID:     input.TargetID,
		SubscriberID: input.SubscriberID,
	}
	db.Create(&NewSubscriber)
	return &NewSubscriber, nil
}

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input model.NewPlaylist) (*model.Playlist, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	NewPlaylist := model.Playlist{
		UserID:      input.UserID,
		Title:       input.Title,
		Private:     input.Private,
		Description: input.Description,
		View:        input.View,
	}
	db.Create(&NewPlaylist)
	return &NewPlaylist, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	NewPost := model.Post{
		UserID:      input.UserID,
		Title:       input.Title,
		Description: input.Description,
		Media:       input.Media,
	}
	db.Create(&NewPost)
	return &NewPost, nil
}

func (r *mutationResolver) CreateSavePlaylist(ctx context.Context, playlistID string, userID string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	NewSavePlaylist := model.SavePlaylist{
		UserID:     userID,
		PlaylistID: playlistID,
	}
	db.Create(&NewSavePlaylist)
	return true, nil
}

func (r *mutationResolver) InsertVideoOnPlaylist(ctx context.Context, input model.NewPlaylistDetail) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	NewPlaylistDetail := model.PlaylistDetail{
		VideoID:    input.VideoID,
		PlaylistID: input.PlaylistID,
		Priority:   input.Priority,
	}
	db.Create(&NewPlaylistDetail)
	return true, nil
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, playlistID string, input model.NewPlaylist) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	var updatePlaylist model.Playlist
	db.Where("id = ?", playlistID).First(&updatePlaylist)
	updatePlaylist.Title = input.Title
	updatePlaylist.Description = input.Description
	updatePlaylist.View = input.View
	updatePlaylist.Private = input.Private
	db.Save(&updatePlaylist)

	return true, nil
}

func (r *mutationResolver) UpdateFeedbackOnComment(ctx context.Context, input model.NewLikeComment) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	var updatedLikeComment model.LikeComment
	db.Where("comment_id = ? AND user_id = ?", input.CommentID, input.UserID).First(&updatedLikeComment)
	updatedLikeComment.Status = input.Status
	db.Save(&updatedLikeComment)

	return true, nil
}

func (r *mutationResolver) UpdateFeedbackOnVideo(ctx context.Context, input model.NewLikeVideo) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	var updatedLikeVideo model.LikeVideo
	db.Where("video_id = ? AND user_id = ?", input.VideoID, input.UserID).First(&updatedLikeVideo)
	updatedLikeVideo.Status = input.Status
	db.Save(&updatedLikeVideo)

	return true, nil
}

func (r *mutationResolver) UpdateFeedbackOnPost(ctx context.Context, input model.NewLikePost) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	var updatedLikePost model.LikePost
	db.Where("post_id = ? AND user_id = ?", input.PostID, input.UserID).First(&updatedLikePost)
	updatedLikePost.Status = input.Status
	db.Save(&updatedLikePost)

	return true, nil
}

func (r *mutationResolver) UpdateSubscriber(ctx context.Context, input model.NewSubscriber) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	var updatedSubscriber model.Subscriber
	db.Where("target_id = ? AND subscriber_id = ?", input.TargetID, input.SubscriberID).First(&updatedSubscriber)
	updatedSubscriber.Notification = input.Notification
	db.Save(&updatedSubscriber)

	return true, nil
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, playlistID string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	if err := db.Where("id = ?", playlistID).Delete(&r.playlist).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DeleteSavePlaylist(ctx context.Context, playlistID string, userID string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	if err := db.Where("playlist_id = ? AND user_id = ? ", playlistID, userID).Delete(&r.savePlaylist).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DeleteVideoOnPlaylist(ctx context.Context, playlistID string, videoID string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	if err := db.Where("playlist_id = ? AND video_id = ? ", playlistID, videoID).Delete(&r.playlistDetail).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DeleteFeedbackOnComment(ctx context.Context, userID string, commentID string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	if err := db.Where("comment_id = ? AND user_id = ?", commentID, userID).Delete(&r.likeComment).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DeleteFeedbackOnVideo(ctx context.Context, userID string, videoID string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	if err := db.Where("video_id = ? AND user_id = ?", videoID, userID).Delete(&r.likeVideo).Error; gorm.IsRecordNotFoundError(err) {
		return false, nil
	}

	return true, nil
}

func (r *mutationResolver) DeleteFeedbackOnPost(ctx context.Context, userID string, postID string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	if err := db.Where("comment_id = ? AND user_id = ?", postID, userID).Delete(&r.likePost).Error; gorm.IsRecordNotFoundError(err) {
		return false, nil
	}

	return true, nil
}

func (r *mutationResolver) DeleteSubscriber(ctx context.Context, targetID string, subscriberID string) (bool, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	if err := db.Where("target_id = ? AND subscriber_id = ?", targetID, subscriberID).Delete(&r.subscriber).Error; gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	return true, nil
}

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
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return false, err
	}
	defer db.Close()
	if err := db.Where("id = ? ", id).Delete(&r.video).Error; gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	return true, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	newComment := model.Comment{
		UserID:   input.UserID,
		VideoID:  input.VideoID,
		ParentID: input.ParentID,
		ReplyTo:  input.ReplyTo,
		Comment:  input.Comment,
	}
	db.Create(&newComment)
	return &newComment, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *playlistDetailResolver) Videos(ctx context.Context, obj *model.PlaylistDetail) (*model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Preload("User.Subscribers").Preload("Location").Preload("LikeVideo").Find(&r.videos)
	return r.videos, nil
}

func (r *queryResolver) FindUser(ctx context.Context, email string) ([]*model.User, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()

	if err := db.Preload("SavePlaylists.User").Preload("SavePlaylists.Videos.User").Preload("Subscribers.Target.Subscribers").Where("email = ?", email).First(&r.users).Error; gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return r.users, nil
}

func (r *queryResolver) FindUserByID(ctx context.Context, id string) ([]*model.User, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()

	if err := db.Preload("SavePlaylists.User").Preload("SavePlaylists.Videos.User").Preload("Subscribers.Target.Subscribers.Target").Where("id = ?", id).First(&r.users).Error; gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return r.users, nil
}

func (r *queryResolver) FindPostByUserID(ctx context.Context, userID string) ([]*model.Post, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Preload("LikePost").Preload("User").Where("user_id = ? ", userID).Order("created_at desc").Find(&r.posts)

	return r.posts, nil
}

func (r *queryResolver) FindPlaylistByUserID(ctx context.Context, id string) ([]*model.Playlist, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	if id != "" {
		db.Preload("User").Preload("Videos.User").Where("user_id = ? ", id).Find(&r.playlists)
	} else {
		db.Preload("User").Preload("Videos.User").Find(&r.playlists)
	}

	return r.playlists, nil
}

func (r *queryResolver) GetVideoPriorityByPlaylistID(ctx context.Context, playlistID string) ([]*model.PlaylistDetail, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Preload("Video.User").Where("playlist_id = ? ", playlistID).Order("priority asc").First(&r.playlistDetails)

	return r.playlistDetails, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Preload("Subscribers.Target").Find(&r.users)
	return r.users, nil
}

func (r *queryResolver) GetLikesByCommentID(ctx context.Context, id string) ([]*model.LikeComment, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Where("comment_id = ?", id).Find(&r.likeComments)
	return r.likeComments, nil
}

func (r *queryResolver) CommentsByVideoID(ctx context.Context, id string) ([]*model.Comment, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Preload("User").Preload("LikeComment").Where("video_id = ?", id).Find(&r.comments)
	return r.comments, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// PlaylistDetail returns generated.PlaylistDetailResolver implementation.
func (r *Resolver) PlaylistDetail() generated.PlaylistDetailResolver {
	return &playlistDetailResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type playlistDetailResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *postResolver) Title(ctx context.Context, obj *model.Post) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

type postResolver struct{ *Resolver }

func (r *playlistResolver) SavePlaylist(ctx context.Context, obj *model.Playlist) ([]*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

type playlistResolver struct{ *Resolver }

func (r *subscriberResolver) Users(ctx context.Context, obj *model.Subscriber) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

type subscriberResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

func (r *userResolver) Subsribers(ctx context.Context, obj *model.User) ([]*model.Subscriber, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *commentResolver) LikeComment(ctx context.Context, obj *model.Comment) (*model.LikeComment, error) {
	panic(fmt.Errorf("not implemented"))
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) LikeComments(ctx context.Context, obj *model.Comment) ([]*model.LikeComment, error) {
	panic(fmt.Errorf("not implemented"))
}
