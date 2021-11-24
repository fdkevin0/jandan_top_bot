package comment

import (
	"github.com/parnurzeal/gorequest"
	"jandan_top_bot/types"
)

func GetCommentOfPost(postID string) (response Response, errs []error) {
	_, _, errs = gorequest.New().
		//https://jandan.net/tucao/all/5080728 == https://jandan.net/tucao/5080728
		Get("https://jandan.net/tucao/all/" + postID).
		EndStruct(&response)
	return response, errs
}

type Response struct {
	Code        int       `json:"code"`
	HotComments []Comment `json:"hot_tucao"`
	Comments    []Comment `json:"tucao"`
}

type Comment struct {
	CommentID       int         `json:"comment_ID"`
	CommentPostID   int         `json:"comment_post_ID"`
	CommentAuthor   string      `json:"comment_author"`
	CommentDate     types.Time  `json:"comment_date"`
	CommentDateInt  int         `json:"comment_date_int"`
	CommentContent  string      `json:"comment_content"`
	CommentParent   int         `json:"comment_parent"`
	CommentReplyID  int         `json:"comment_reply_ID"`
	IsJandanUser    int         `json:"is_jandan_user"`
	IsTipUser       int         `json:"is_tip_user"`
	VotePositive    int         `json:"vote_positive"`
	VoteNegative    int         `json:"vote_negative"`
	SubCommentCount int         `json:"sub_comment_count"`
	Images          interface{} `json:"images"`
}
