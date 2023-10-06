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
	var userQuestionVoteID int
	var err error

	err = qs.trManager.WithTransaction(ctx, func(ctx context.Context) error {
		reqCreateUserQuestionVote := req.ToReqCreateUserQuestionVote()
		reqCreateUserQuestionVote.VoteOptionID = voteOptionID
		userQuestionVoteID, err = qs.createUserQuestionVote(ctx, reqCreateUserQuestionVote)
		if err != nil {
			return errors.Wrap(err, "error create user question vote")
		}

		err := qs.incrementQuestionStatisticOnVote(ctx, req, voteOptionID)
		if err != nil {
			return errors.Wrap(err, "error increment question_statistics")
		}

		return nil
	})
	if err != nil {
		return 0, errors.Wrap(err, "error transaction create user and increment question_statistics")
	}

	return userQuestionVoteID, nil
}

func (qs *QuestionService) incrementQuestionStatisticOnVote(ctx context.Context, req dto.ReqVoteThumbs, voteOptionID constant.VoteOption) error {
	isVotingThumbsUp := voteOptionID == constant.VoteOptionThumbsUp
	if isVotingThumbsUp {
		_, err := qs.repo.IncrementQuestionStatisticColumnThumbsUpByQuestionID(
			ctx,
			req.QuestionID,
		)
		if err != nil {
			return errors.Wrap(err, "error increment question_statistics column thumbs_up")
		}
	} else {
		_, err := qs.repo.IncrementQuestionStatisticColumnThumbsDownByQuestionID(
			ctx,
			req.QuestionID,
		)
		if err != nil {
			return errors.Wrap(err, "error increment question_statistics column thumbs_down")
		}
	}

	return nil
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

		return qs.syncQuestionStatistic(ctx, userQuestionVote, voteOptionID, req)
	})

	if err != nil {
		return 0, errors.Wrap(err, "error transaction update user_question_votes and sync question_statistics")
	}

	return userQuestionVote.ID, nil
}

func (qs *QuestionService) syncQuestionStatistic(ctx context.Context, userQuestionVote model.UserQuestionVote, voteOptionID constant.VoteOption, req dto.ReqVoteThumbs) error {
	isPreviousVoteThumbsUp := userQuestionVote.VoteOptionID == constant.VoteOptionThumbsUp
	if isPreviousVoteThumbsUp {
		isVotingThumbsDown := voteOptionID == constant.VoteOptionThumbsDown
		return qs.handlePreviousVoteThumbsUp(ctx, isVotingThumbsDown, req.QuestionID)
	}

	isVotingThumbsUp := voteOptionID == constant.VoteOptionThumbsUp
	return qs.handlePreviousVoteThumbsDown(ctx, isVotingThumbsUp, req.QuestionID)
}

func (qs *QuestionService) handlePreviousVoteThumbsUp(ctx context.Context, isVotingThumbsDown bool, questionID int) error {
	if isVotingThumbsDown {
		_, err := qs.repo.IncrementQuestionStatisticColumnThumbsDownByQuestionID(
			ctx,
			questionID,
		)
		if err != nil {
			return errors.Wrap(err, "error increment question_statistics column thumbs_down")
		}
	}

	_, err := qs.repo.DecrementQuestionStatisticColumnThumbsUpByQuestionID(
		ctx,
		questionID,
	)
	if err != nil {
		return errors.Wrap(err, "error decrement question_statistics column thumbs_up")
	}

	return nil
}

func (qs *QuestionService) handlePreviousVoteThumbsDown(ctx context.Context, isVotingThumbsUp bool, questionID int) error {
	if isVotingThumbsUp {
		_, err := qs.repo.IncrementQuestionStatisticColumnThumbsUpByQuestionID(
			ctx,
			questionID,
		)
		if err != nil {
			return errors.Wrap(err, "error increment question_statistics column thumbs_up")
		}
	}

	_, err := qs.repo.DecrementQuestionStatisticColumnThumbsDownByQuestionID(
		ctx,
		questionID,
	)
	if err != nil {
		return errors.Wrap(err, "error decrement question_statistics column thumbs_down")
	}

	return nil
}
