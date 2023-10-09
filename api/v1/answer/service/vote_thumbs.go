package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (as *AnswerService) voteThumbs(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	isVoteOptionValid := voteOptionID == constant.VoteOptionThumbsUp ||
		voteOptionID == constant.VoteOptionThumbsDown

	if !isVoteOptionValid {
		return 0, errors.New("error vote option not valid, please choose thumbs up or thumbs down")
	}

	userAnswerVote, err := as.getUserAnswerVoteByUserIDAndAnswerID(ctx, dto.ReqGetUserAnswerVoteByUserIDAndAnswerID(req))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.Wrap(err, "error get user_answer_votes by user_id and answer_id")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return as.handleFirstTimeUserVoteAnswer(ctx, req, voteOptionID)
	}

	return as.handleNTimeUserVoteAnswer(ctx, req, voteOptionID, userAnswerVote)
}

func (as *AnswerService) handleFirstTimeUserVoteAnswer(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption) (int, error) {
	err := as.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		reqCreateUserAnswerVote := req.ToReqCreateUserAnswerVote()
		reqCreateUserAnswerVote.VoteOptionID = voteOptionID
		_, err := as.createUserAnswerVote(ctx, reqCreateUserAnswerVote)
		if err != nil {
			return errors.Wrap(err, "error create user answer vote")
		}

		isVotingThumbsUp := voteOptionID == constant.VoteOptionThumbsUp

		if isVotingThumbsUp {
			err = as.incrementAnswerThumbsUpCount(ctx, req.AnswerID)
		} else {
			err = as.incrementAnswerThumbsDownCount(ctx, req.AnswerID)
		}

		if err != nil {
			return errors.Wrap(err, "error increment answer_statistics")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction create user answer vote and increment answer_statistics")
	}

	return req.AnswerID, nil
}

func (as *AnswerService) handleNTimeUserVoteAnswer(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption, userAnswerVote model.UserAnswerVote) (int, error) {
	if userAnswerVote.VoteOptionID == voteOptionID {
		voteOptionID = constant.VoteOptionCancel
	}

	err := as.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		_, err := as.repo.UpdateUserAnswerVoteColumnVoteOptionIDByID(
			ctx,
			dto.ReqUpdateUserAnswerVoteColumnVoteOptionIDByID{
				ID:           userAnswerVote.ID,
				VoteOptionID: voteOptionID,
			},
		)

		if err != nil {
			return errors.Wrap(err, "error update user_answer_votes")
		}

		err = as.syncAnswerStatistic(ctx, userAnswerVote, voteOptionID, req)
		if err != nil {
			return errors.Wrap(err, "error sync answer_statistics")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction update user_answer_votes and sync answer_statistics")
	}

	return req.AnswerID, nil
}

func (as *AnswerService) syncAnswerStatistic(ctx context.Context, userAnswerVote model.UserAnswerVote, voteOptionID constant.VoteOption, req dto.ReqVoteThumbs) error {
	isPreviousVoteThumbsUp := userAnswerVote.VoteOptionID == constant.VoteOptionThumbsUp
	isPreviousVoteThumbsDown := userAnswerVote.VoteOptionID == constant.VoteOptionThumbsDown
	isPreviousVoteCancel := userAnswerVote.VoteOptionID == constant.VoteOptionCancel

	isVotingThumbsUp := voteOptionID == constant.VoteOptionThumbsUp
	isVotingThumbsDown := voteOptionID == constant.VoteOptionThumbsDown
	isVotingThumbsCancel := voteOptionID == constant.VoteOptionCancel

	var err error

	if isPreviousVoteThumbsUp {
		if isVotingThumbsUp || isVotingThumbsCancel {
			err = as.decrementAnswerThumbsUpCount(ctx, req.AnswerID)
		} else if isVotingThumbsDown {
			err = as.decrementAnswerThumbsUpCountAndIncrementThumbsDownCount(ctx, req.AnswerID)
		}
	} else if isPreviousVoteThumbsDown {
		if isVotingThumbsUp {
			err = as.incrementAnswerThumbsUpCountAndDecrementThumbsDownCount(ctx, req.AnswerID)
		} else if isVotingThumbsDown || isVotingThumbsCancel {
			err = as.decrementAnswerThumbsDownCount(ctx, req.AnswerID)
		}
	} else if isPreviousVoteCancel {
		if isVotingThumbsUp {
			err = as.incrementAnswerThumbsUpCount(ctx, req.AnswerID)
		} else if isVotingThumbsDown {
			err = as.incrementAnswerThumbsDownCount(ctx, req.AnswerID)
		}
	}

	return err
}

func (as *AnswerService) decrementAnswerThumbsUpCount(ctx context.Context, answerID int) error {
	_, err := as.repo.DecrementAnswerStatisticColumnThumbsUpByAnswerID(ctx, answerID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_ANSWER_THUMBS_UP)
	}
	return nil
}

func (as *AnswerService) decrementAnswerThumbsUpCountAndIncrementThumbsDownCount(ctx context.Context, answerID int) error {
	_, err := as.repo.DecrementAnswerStatisticColumnThumbsUpByAnswerID(ctx, answerID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_ANSWER_THUMBS_UP)
	}
	_, err = as.repo.IncrementAnswerStatisticColumnThumbsDownByAnswerID(ctx, answerID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_ANSWER_THUMBS_DOWN)
	}
	return nil
}

func (as *AnswerService) incrementAnswerThumbsUpCountAndDecrementThumbsDownCount(ctx context.Context, answerID int) error {
	_, err := as.repo.IncrementAnswerStatisticColumnThumbsUpByAnswerID(ctx, answerID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_ANSWER_THUMBS_UP)
	}
	_, err = as.repo.DecrementAnswerStatisticColumnThumbsDownByAnswerID(ctx, answerID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_ANSWER_THUMBS_DOWN)
	}
	return nil
}

func (as *AnswerService) decrementAnswerThumbsDownCount(ctx context.Context, answerID int) error {
	_, err := as.repo.DecrementAnswerStatisticColumnThumbsDownByAnswerID(ctx, answerID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_ANSWER_THUMBS_DOWN)
	}
	return nil
}

func (as *AnswerService) incrementAnswerThumbsUpCount(ctx context.Context, answerID int) error {
	_, err := as.repo.IncrementAnswerStatisticColumnThumbsUpByAnswerID(ctx, answerID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_ANSWER_THUMBS_UP)
	}
	return nil
}

func (as *AnswerService) incrementAnswerThumbsDownCount(ctx context.Context, answerID int) error {
	_, err := as.repo.IncrementAnswerStatisticColumnThumbsDownByAnswerID(ctx, answerID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_ANSWER_THUMBS_DOWN)
	}
	return nil
}
