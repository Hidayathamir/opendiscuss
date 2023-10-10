package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/pkg/errors"
)

func (cs *CommentService) GetCommentListByAnswerID(ctx context.Context, answerID int) ([]dto.CommentHighlight, error) {
	if answerID == 0 {
		return nil, errors.New("answer id can not be empty")
	}

	comments, err := cs.repo.GetCommentListByAnswerID(ctx, answerID)
	if err != nil {
		return nil, errors.Wrap(err, "error get comment list by answer id")
	}

	for _, comment := range comments {
		comment.ThumbsRate = comment.ThumbsUp - comment.ThumbsDown
	}

	return comments, nil
}

func (cs *CommentService) GetSubCommentListByCommentID(ctx context.Context, commentID int) ([]dto.CommentHighlight, error) {
	if commentID == 0 {
		return nil, errors.New("comment id can not be empty")
	}

	subComments, err := cs.repo.GetSubCommentListByCommentID(ctx, commentID)
	if err != nil {
		return nil, errors.Wrap(err, "error get sub comment list by comment id")
	}

	for _, subComment := range subComments {
		subComment.ThumbsRate = subComment.ThumbsUp - subComment.ThumbsDown
	}

	return subComments, nil
}
