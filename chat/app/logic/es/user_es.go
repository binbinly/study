package es

import (
	"context"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"chat/app/logic/model"
)

const indexUser = "user"

//IUser elastic 用户操作接口
type IUser interface {
	UserPut(ctx context.Context, user *model.UserModel) (err error)
	UserSearch(ctx context.Context, keyword string) (users []*model.UserEs, err error)
	UserDelete(ctx context.Context, id string) (err error)
	UserUpdate(ctx context.Context, id uint32, data map[string]interface{}) (err error)
}

//UserPut 写入用户数据
func (e *ES) UserPut(ctx context.Context, user *model.UserModel) (err error) {
	u := model.UserEs{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Phone:    strconv.FormatInt(user.Phone, 10),
	}
	_, err = e.client.Index().Index(e.userIndex()).
		Id(strconv.Itoa(int(user.ID))).BodyJson(u).Do(ctx)
	if err != nil {
		return errors.Wrapf(err, "[es.user] put id:%v", user.ID)
	}
	return nil
}

//UserSearch 搜索用户
func (e *ES) UserSearch(ctx context.Context, keyword string) (users []*model.UserEs, err error) {
	q := elastic.NewQueryStringQuery(keyword + "*").Field("username").Field("nickname").Field("phone")
	res, err := e.client.Search(e.userIndex()).Query(q).Do(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[es.user] search keyword:%v", keyword)
	}
	var typ model.UserEs
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		u := item.(model.UserEs)
		users = append(users, &u)
	}
	return users, nil
}

//UserDelete 删除用户
func (e *ES) UserDelete(ctx context.Context, id string) (err error) {
	_, err = e.client.Delete().Index(e.userIndex()).Id(id).Do(ctx)
	if err != nil {
		return errors.Wrapf(err, "[es.user] delete id:%v", id)
	}
	return nil
}

//UserUpdate 修改用户
func (e *ES) UserUpdate(ctx context.Context, id uint32, data map[string]interface{}) (err error) {
	_, err = e.client.Update().Index(e.userIndex()).Id(strconv.Itoa(int(id))).Doc(data).Do(ctx)
	if err != nil {
		return errors.Wrapf(err, "[es.user] update id:%v", id)
	}
	return nil
}

func (e *ES) userIndex() string {
	return e.indexPrefix + "_" + indexUser
}
