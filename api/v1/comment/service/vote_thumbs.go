package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (cs *CommentService) voteThumbs(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	isVoteOptionValid := voteOptionID == constant.VoteOptionThumbsUp ||
		voteOptionID == constant.VoteOptionThumbsDown

	if !isVoteOptionValid {
		return 0, errors.New("error vote option not valid, please choose thumbs up or thumbs down")
	}

	userCommentVote, err := cs.getUserCommentVoteByUserIDAndCommentID(ctx, dto.ReqGetUserCommentVoteByUserIDAndCommentID(req))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.Wrap(err, "error get user_comment_votes by user_id and comment_id")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return cs.handleFirstTimeUserVoteComment(ctx, req, voteOptionID)
	}

	return cs.handleNTimeUserVoteComment(ctx, req, voteOptionID, userCommentVote)
}

func (cs *CommentService) handleFirstTimeUserVoteComment(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption) (int, error) {
	err := cs.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		reqCreateUserCommentVote := req.ToReqCreateUserCommentVote()
		reqCreateUserCommentVote.VoteOptionID = voteOptionID
		_, err := cs.createUserCommentVote(ctx, reqCreateUserCommentVote)
		if err != nil {
			return errors.Wrap(err, "error create user comment vote")
		}

		isVotingThumbsUp := voteOptionID == constant.VoteOptionThumbsUp

		if isVotingThumbsUp {
			err = cs.incrementCommentThumbsUpCount(ctx, req.CommentID)
		} else {
			err = cs.incrementCommentThumbsDownCount(ctx, req.CommentID)
		}

		if err != nil {
			return errors.Wrap(err, "error increment comment_statistics")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction create user comment vote and increment comment_statistics")
	}

	return req.CommentID, nil
}

func (cs *CommentService) handleNTimeUserVoteComment(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption, userCommentVote model.UserCommentVote) (int, error) {
	if userCommentVote.VoteOptionID == voteOptionID {
		voteOptionID = constant.VoteOptionCancel
	}

	err := cs.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		_, err := cs.repo.UpdateUserCommentVoteColumnVoteOptionIDByID(
			ctx,
			dto.ReqUpdateUserCommentVoteColumnVoteOptionIDByID{
				ID:           userCommentVote.ID,
				VoteOptionID: voteOptionID,
			},
		)

		if err != nil {
			return errors.Wrap(err, "error update user_comment_votes")
		}

		err = cs.syncCommentStatistic(ctx, userCommentVote, voteOptionID, req)
		if err != nil {
			return errors.Wrap(err, "error sync comment_statistics")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction update user_comment_votes and sync comment_statistics")
	}

	return req.CommentID, nil
}

func (cs *CommentService) syncCommentStatistic(ctx context.Context, userCommentVote model.UserCommentVote, voteOptionID constant.VoteOption, req dto.ReqVoteThumbs) error {
	isPreviousVoteThumbsUp := userCommentVote.VoteOptionID == constant.VoteOptionThumbsUp
	isPreviousVoteThumbsDown := userCommentVote.VoteOptionID == constant.VoteOptionThumbsDown
	isPreviousVoteCancel := userCommentVote.VoteOptionID == constant.VoteOptionCancel

	isVotingThumbsUp := voteOptionID == constant.VoteOptionThumbsUp
	isVotingThumbsDown := voteOptionID == constant.VoteOptionThumbsDown
	isVotingThumbsCancel := voteOptionID == constant.VoteOptionCancel

	var err error

	if isPreviousVoteThumbsUp {
		if isVotingThumbsUp || isVotingThumbsCancel {
			err = cs.decrementCommentThumbsUpCount(ctx, req.CommentID)
		} else if isVotingThumbsDown {
			err = cs.decrementCommentThumbsUpCountAndIncrementThumbsDownCount(ctx, req.CommentID)
		}
	} else if isPreviousVoteThumbsDown {
		if isVotingThumbsUp {
			err = cs.incrementCommentThumbsUpCountAndDecrementThumbsDownCount(ctx, req.CommentID)
		} else if isVotingThumbsDown || isVotingThumbsCancel {
			err = cs.decrementCommentThumbsDownCount(ctx, req.CommentID)
		}
	} else if isPreviousVoteCancel {
		if isVotingThumbsUp {
			err = cs.incrementCommentThumbsUpCount(ctx, req.CommentID)
		} else if isVotingThumbsDown {
			err = cs.incrementCommentThumbsDownCount(ctx, req.CommentID)
		}
	}

	return err
}

func (cs *CommentService) decrementCommentThumbsUpCount(ctx context.Context, commentID int) error {
	_, err := cs.repo.DecrementCommentStatisticColumnThumbsUpByCommentID(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_COMMENT_THUMBS_UP)
	}
	return nil
}

func (cs *CommentService) decrementCommentThumbsUpCountAndIncrementThumbsDownCount(ctx context.Context, commentID int) error {
	_, err := cs.repo.DecrementCommentStatisticColumnThumbsUpByCommentID(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_COMMENT_THUMBS_UP)
	}
	_, err = cs.repo.IncrementCommentStatisticColumnThumbsDownByCommentID(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_COMMENT_THUMBS_DOWN)
	}
	return nil
}

func (cs *CommentService) incrementCommentThumbsUpCountAndDecrementThumbsDownCount(ctx context.Context, commentID int) error {
	_, err := cs.repo.IncrementCommentStatisticColumnThumbsUpByCommentID(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_COMMENT_THUMBS_UP)
	}
	_, err = cs.repo.DecrementCommentStatisticColumnThumbsDownByCommentID(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_COMMENT_THUMBS_DOWN)
	}
	return nil
}

func (cs *CommentService) decrementCommentThumbsDownCount(ctx context.Context, commentID int) error {
	_, err := cs.repo.DecrementCommentStatisticColumnThumbsDownByCommentID(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_COMMENT_THUMBS_DOWN)
	}
	return nil
}

func (cs *CommentService) incrementCommentThumbsUpCount(ctx context.Context, commentID int) error {
	_, err := cs.repo.IncrementCommentStatisticColumnThumbsUpByCommentID(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_COMMENT_THUMBS_UP)
	}
	return nil
}

func (cs *CommentService) incrementCommentThumbsDownCount(ctx context.Context, commentID int) error {
	_, err := cs.repo.IncrementCommentStatisticColumnThumbsDownByCommentID(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_COMMENT_THUMBS_DOWN)
	}
	return nil
}
