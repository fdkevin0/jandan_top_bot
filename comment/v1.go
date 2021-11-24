package comment

import (
	"github.com/parnurzeal/gorequest"
	"jandan_top_bot/types"
	"time"
)

func GetCommentsV1ByPostID(postID string) (response ResponseV1, errs []error) {
	_, _, errs = gorequest.New().
		Get("https://api.jandan.net/api/v1/tucao/list/" + postID).
		EndStruct(&response)
	return response, errs
}

type ResponseV1 struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Hot  []V1 `json:"hot"`
		List []V1 `json:"list"`
	} `json:"data"`
}

type V1 struct {
	ID           int             `json:"id"`
	PostID       int             `json:"post_id"`
	Author       string          `json:"author"`
	AuthorType   int             `json:"author_type"`
	Date         time.Time       `json:"date"`
	DateUnix     types.Timestamp `json:"date_unix"`
	Content      string          `json:"content"`
	UserID       int             `json:"user_id"`
	VotePositive int             `json:"vote_positive"`
	VoteNegative int             `json:"vote_negative"`
	Images       []string        `json:"images"`
	AtComments   []AtCommentV1   `json:"at_comments,omitempty"`
}

type AtCommentV1 struct {
	AtAuthor    string `json:"at_author"`
	AtCommentID string `json:"at_comment_id"`
}
