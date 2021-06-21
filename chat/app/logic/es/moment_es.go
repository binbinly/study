package es

import (
	"context"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"chat/app/logic/model"
)

const indexMoment = "moment"

//IMoment elastic 朋友圈操作接口
type IMoment interface {
	MomentPut(ctx context.Context, user *model.MomentModel) (err error)
	MomentSearch(ctx context.Context, userId uint32, keyword string) (users []*model.MomentEs, err error)
	MomentDelete(ctx context.Context, id string) (err error)
	MomentUpdate(ctx context.Context, id uint32, data map[string]interface{}) (err error)
}

//MomentPut 写入朋友圈数据
func (e *ES) MomentPut(ctx context.Context, moment *model.MomentModel) (err error) {
	m := model.MomentEs{
		ID:      moment.ID,
		UserID:  moment.UserID,
		Content: moment.Content,
	}
	_, err = e.client.Index().Index(e.momentIndex()).
		Id(strconv.Itoa(int(moment.ID))).BodyJson(m).Do(ctx)
	if err != nil {
		return errors.Wrapf(err, "[es.moment] put id:%v", moment.ID)
	}
	return nil
}

//MomentSearch 搜索我发布的朋友圈
func (e *ES) MomentSearch(ctx context.Context, userId uint32, keyword string) (moments []*model.MomentEs, err error) {
	q := elastic.NewBoolQuery()
	q.Must(elastic.NewQueryStringQuery("content:"+keyword),
		elastic.NewMatchQuery("user_id", userId))
	res, err := e.client.Search(e.momentIndex()).Query(q).Do(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[es.moment] search keyword:%v", keyword)
	}
	var typ model.MomentEs
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		u := item.(model.MomentEs)
		moments = append(moments, &u)
	}
	return moments, nil
}

//MomentDelete 删除用户
func (e *ES) MomentDelete(ctx context.Context, id string) (err error) {
	_, err = e.client.Delete().Index(e.momentIndex()).Id(id).Do(ctx)
	if err != nil {
		return errors.Wrapf(err, "[es.moment] delete id:%v", id)
	}
	return nil
}

//MomentUpdate 修改用户
func (e *ES) MomentUpdate(ctx context.Context, id uint32, data map[string]interface{}) (err error) {
	_, err = e.client.Update().Index(e.momentIndex()).Id(strconv.Itoa(int(id))).Doc(data).Do(ctx)
	if err != nil {
		return errors.Wrapf(err, "[es.moment] update id:%v", id)
	}
	return nil
}

func (e *ES) momentIndex() string {
	return e.indexPrefix + "_" + indexMoment
}
