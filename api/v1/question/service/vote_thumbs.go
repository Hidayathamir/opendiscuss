package service

import (
	"context"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (qs *QuestionService) voteThumbs(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption) (int, error) {
	if err := req.Validate(); err != nil {
		return 0, errors.Wrap(err, constant.ERR_REQ_VALIDATE)
	}

	isVoteOptionValid := voteOptionID == constant.VoteOptionThumbsUp ||
		voteOptionID == constant.VoteOptionThumbsDown

	if !isVoteOptionValid {
		return 0, errors.New("error vote option not valid, please choose thumbs up or thumbs down")
	}

	userQuestionVote, err := qs.getUserQuestionVoteByUserIDAndQuestionID(ctx, dto.ReqGetUserQuestionVoteByUserIDAndQuestionID(req))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.Wrap(err, "error get user_question_votes by user_id and question_id")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return qs.handleFirstTimeUserVoteQuestion(ctx, req, voteOptionID)
	}

	return qs.handleNTimeUserVoteQuestion(ctx, req, voteOptionID, userQuestionVote)
}

func (qs *QuestionService) handleFirstTimeUserVoteQuestion(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption) (int, error) {
	err := qs.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		reqCreateUserQuestionVote := req.ToReqCreateUserQuestionVote()
		reqCreateUserQuestionVote.VoteOptionID = voteOptionID
		_, err := qs.createUserQuestionVote(ctx, reqCreateUserQuestionVote)
		if err != nil {
			return errors.Wrap(err, "error create user question vote")
		}

		isVotingThumbsUp := voteOptionID == constant.VoteOptionThumbsUp

		if isVotingThumbsUp {
			err = qs.incrementQuestionThumbsUpCount(ctx, req.QuestionID)
		} else {
			err = qs.incrementQuestionThumbsDownCount(ctx, req.QuestionID)
		}

		if err != nil {
			return errors.Wrap(err, "error increment question_statistics")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction create user question vote and increment question_statistics")
	}

	return req.QuestionID, nil
}

func (qs *QuestionService) handleNTimeUserVoteQuestion(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption, userQuestionVote model.UserQuestionVote) (int, error) {
	if userQuestionVote.VoteOptionID == voteOptionID {
		voteOptionID = constant.VoteOptionCancel
	}

	err := qs.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		_, err := qs.repo.UpdateUserQuestionVoteColumnVoteOptionIDByID(
			ctx,
			dto.ReqUpdateUserQuestionVoteColumnVoteOptionIDByID{
				ID:           userQuestionVote.ID,
				VoteOptionID: voteOptionID,
			},
		)

		if err != nil {
			return errors.Wrap(err, "error update user_question_votes")
		}

		err = qs.syncQuestionStatistic(ctx, userQuestionVote, voteOptionID, req)
		if err != nil {
			return errors.Wrap(err, "error sync question_statistics")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction update user_question_votes and sync question_statistics")
	}

	return req.QuestionID, nil
}

func (qs *QuestionService) syncQuestionStatistic(ctx context.Context, userQuestionVote model.UserQuestionVote, voteOptionID constant.VoteOption, req dto.ReqVoteThumbs) error {
	isPreviousVoteThumbsUp := userQuestionVote.VoteOptionID == constant.VoteOptionThumbsUp
	isPreviousVoteThumbsDown := userQuestionVote.VoteOptionID == constant.VoteOptionThumbsDown
	isPreviousVoteCancel := userQuestionVote.VoteOptionID == constant.VoteOptionCancel

	isVotingThumbsUp := voteOptionID == constant.VoteOptionThumbsUp
	isVotingThumbsDown := voteOptionID == constant.VoteOptionThumbsDown
	isVotingThumbsCancel := voteOptionID == constant.VoteOptionCancel

	var err error

	if isPreviousVoteThumbsUp {
		if isVotingThumbsUp || isVotingThumbsCancel {
			err = qs.decrementQuestionThumbsUpCount(ctx, req.QuestionID)
		} else if isVotingThumbsDown {
			err = qs.decrementQuestionThumbsUpCountAndIncrementThumbsDownCount(ctx, req.QuestionID)
		}
	} else if isPreviousVoteThumbsDown {
		if isVotingThumbsUp {
			err = qs.incrementQuestionThumbsUpCountAndDecrementThumbsDownCount(ctx, req.QuestionID)
		} else if isVotingThumbsDown || isVotingThumbsCancel {
			err = qs.decrementQuestionThumbsDownCount(ctx, req.QuestionID)
		}
	} else if isPreviousVoteCancel {
		if isVotingThumbsUp {
			err = qs.incrementQuestionThumbsUpCount(ctx, req.QuestionID)
		} else if isVotingThumbsDown {
			err = qs.incrementQuestionThumbsDownCount(ctx, req.QuestionID)
		}
	}

	return err
}

func (qs *QuestionService) decrementQuestionThumbsUpCount(ctx context.Context, questionID int) error {
	_, err := qs.repo.DecrementQuestionStatisticColumnThumbsUpByQuestionID(ctx, questionID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_QUESTION_THUMBS_UP)
	}
	return nil
}

func (qs *QuestionService) decrementQuestionThumbsUpCountAndIncrementThumbsDownCount(ctx context.Context, questionID int) error {
	_, err := qs.repo.DecrementQuestionStatisticColumnThumbsUpByQuestionID(ctx, questionID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_QUESTION_THUMBS_UP)
	}
	_, err = qs.repo.IncrementQuestionStatisticColumnThumbsDownByQuestionID(ctx, questionID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_QUESTION_THUMBS_DOWN)
	}
	return nil
}

func (qs *QuestionService) incrementQuestionThumbsUpCountAndDecrementThumbsDownCount(ctx context.Context, questionID int) error {
	_, err := qs.repo.IncrementQuestionStatisticColumnThumbsUpByQuestionID(ctx, questionID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_QUESTION_THUMBS_UP)
	}
	_, err = qs.repo.DecrementQuestionStatisticColumnThumbsDownByQuestionID(ctx, questionID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_QUESTION_THUMBS_DOWN)
	}
	return nil
}

func (qs *QuestionService) decrementQuestionThumbsDownCount(ctx context.Context, questionID int) error {
	_, err := qs.repo.DecrementQuestionStatisticColumnThumbsDownByQuestionID(ctx, questionID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_DECREMENT_QUESTION_THUMBS_DOWN)
	}
	return nil
}

func (qs *QuestionService) incrementQuestionThumbsUpCount(ctx context.Context, questionID int) error {
	_, err := qs.repo.IncrementQuestionStatisticColumnThumbsUpByQuestionID(ctx, questionID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_QUESTION_THUMBS_UP)
	}
	return nil
}

func (qs *QuestionService) incrementQuestionThumbsDownCount(ctx context.Context, questionID int) error {
	_, err := qs.repo.IncrementQuestionStatisticColumnThumbsDownByQuestionID(ctx, questionID)
	if err != nil {
		return errors.Wrap(err, constant.ERR_INCREMENT_QUESTION_THUMBS_DOWN)
	}
	return nil
}
