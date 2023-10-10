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

	for i := range comments {
		comments[i].ThumbsRate = comments[i].ThumbsUp - comments[i].ThumbsDown
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

	for i := range subComments {
		subComments[i].ThumbsRate = subComments[i].ThumbsUp - subComments[i].ThumbsDown
	}

	return subComments, nil
}
