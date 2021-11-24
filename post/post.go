package post

import "jandan_top_bot/types"

type Response struct {
	Status     string `json:"status"`
	Count      int    `json:"count"`
	CountTotal int    `json:"count_total"`
	Pages      int    `json:"pages"`
	Posts      []Post `json:"posts"`
}

type Comment struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Url          string     `json:"url"`
	Date         types.Time `json:"date"`
	Content      string     `json:"content"`
	Parent       int        `json:"parent"`
	VotePositive int        `json:"vote_positive"`
	VoteNegative int        `json:"vote_negative"`
	Index        int        `json:"index"`
	Author       Author     `json:"author,omitempty"`
}

type Author struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Nickname    string `json:"nickname"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type Tag struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PostCount   int    `json:"post_count"`
}

type Category struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Parent      int    `json:"parent"`
	PostCount   int    `json:"post_count"`
}

type Post struct {
	Id            int           `json:"id"`
	Type          string        `json:"type"`
	Slug          string        `json:"slug"`
	Url           string        `json:"url"`
	Status        string        `json:"status"`
	Title         string        `json:"title"`
	TitlePlain    string        `json:"title_plain"`
	Content       string        `json:"content"`
	Excerpt       string        `json:"excerpt"`
	Date          types.Time    `json:"date"`
	Modified      types.Time    `json:"modified"`
	Categories    []Category    `json:"categories"`
	Tags          []Tag         `json:"tags"`
	Author        Author        `json:"author"`
	Comments      []Comment     `json:"comments"`
	CommentsRank  []Comment     `json:"comments_rank"`
	Attachments   []interface{} `json:"attachments"`
	CommentCount  int           `json:"comment_count"`
	CommentStatus string        `json:"comment_status"`
	CustomFields  struct {
		ThumbC []string `json:"thumb_c"`
	} `json:"custom_fields"`
}
