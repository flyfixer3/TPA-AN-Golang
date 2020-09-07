package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"yourJube_API/graph/model"
)

//Resolver -> ini Static DB
type Resolver struct {
	videos          []*model.Video
	video           model.Video
	user            *model.User
	users           []*model.User
	comments        []*model.Comment
	comment         model.Comment
	likeComments    []*model.LikeComment
	likeComment     model.LikeComment
	likeVideos      []*model.LikeVideo
	likeVideo       model.LikeVideo
	likePosts       []*model.LikePost
	likePost        model.LikePost
	subscribers     []*model.Subscriber
	subscriber      model.Subscriber
	playlists       []*model.Playlist
	playlist        model.Playlist
	playlistDetails []*model.PlaylistDetail
	playlistDetail  model.PlaylistDetail
	savePlaylist    model.SavePlaylist
	posts           []*model.Post
}
