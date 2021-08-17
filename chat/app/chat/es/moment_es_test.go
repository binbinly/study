package es

import (
	"chat/app/chat/model"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestES_Moment(t *testing.T) {
	err := es.MomentPut(context.Background(), &model.MomentModel{
		PriID:   model.PriID{ID: 1},
		UID:     model.UID{UserID: 1},
		Content: "我是人间惆怅客，知君何事泪纵横",
	})
	assert.NoError(t, err)
	err = es.MomentPut(context.Background(), &model.MomentModel{
		PriID:   model.PriID{ID: 2},
		UID:     model.UID{UserID: 2},
		Content: "我是人间惆怅客，知君何事泪纵横",
	})
	assert.NoError(t, err)
	err = es.MomentPut(context.Background(), &model.MomentModel{
		PriID:   model.PriID{ID: 3},
		UID:     model.UID{UserID: 3},
		Content: "我是人间惆怅客，知君何事泪纵横",
	})
	assert.NoError(t, err)

	ids, err := es.MomentSearch(context.Background(), 1, "我")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ids))
}

func TestES_MomentDelete(t *testing.T) {
	err := es.MomentDelete(context.Background(), "1")
	assert.NoError(t, err)
	err = es.MomentDelete(context.Background(), "2")
	assert.NoError(t, err)
	err = es.MomentDelete(context.Background(), "3")
	assert.NoError(t, err)
}
