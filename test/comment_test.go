package test

import (
	"jandan_top_bot/comment"
	"reflect"
	"testing"
)

func TestGetCommentOfPost(t *testing.T) {
	type args struct {
		postID string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse comment.Response
		wantErrs     []error
	}{
		{
			name: "test1",
			args: args{
				postID: "5080728",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotErrs := comment.GetCommentOfPost(tt.args.postID)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("GetCommentOfPost() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("GetCommentOfPost() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}
