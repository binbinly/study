package repo

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"mall/app/model"
	"mall/pkg/log"
	"mall/pkg/redis"
)

//ICart 购物车接口
type ICart interface {
	GetCartByID(ctx context.Context, userID int, id string) (*model.CartModel, error)
	GetCartsByIds(ctx context.Context, userID int, ids []string) ([]*model.CartModel, error)
	AddCart(ctx context.Context, userID int, cart *model.CartModel) error
	EditCart(ctx context.Context, userID int, id string, cart *model.CartModel) error
	DelCart(ctx context.Context, userID int, id []string) error
	EmptyCart(ctx context.Context, userID int) error
	CartList(ctx context.Context, userID int) ([]*model.CartModel, error)
}

//AddCart 添加购物车
func (r *Repo) AddCart(ctx context.Context, userID int, cart *model.CartModel) error {
	data, err := json.Marshal(cart)
	if err != nil {
		return errors.Wrapf(err, "[repo.cart] json marshal")
	}
	return redis.Client.HSet(ctx, model.BuildCartKey(userID), cart.ID, data).Err()
}

//EditCart 更新购物车
func (r *Repo) EditCart(ctx context.Context, userID int, id string, cart *model.CartModel) error {
	data, err := json.Marshal(cart)
	if err != nil {
		return errors.Wrapf(err, "[repo.cart] json marshal")
	}
	return redis.Client.HSet(ctx, model.BuildCartKey(userID), id, data).Err()
}

//GetCartByID 获取购物车数据
func (r *Repo) GetCartByID(ctx context.Context, userID int, id string) (*model.CartModel, error) {
	data, err := redis.Client.HGet(ctx, model.BuildCartKey(userID), id).Result()
	if err != nil {
		return nil, errors.Wrap(err, "[repo.cart] hget db")
	}
	cart := &model.CartModel{}
	err = json.Unmarshal([]byte(data), cart)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.cart] json unmarshal")
	}
	return cart, nil
}

//GetCartsByIds 批量获取购物车数据
func (r *Repo) GetCartsByIds(ctx context.Context, userID int, ids []string) ([]*model.CartModel, error) {
	data, err := redis.Client.HMGet(ctx, model.BuildCartKey(userID), ids...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "[repo.cart] hmget db")
	}
	var carts []*model.CartModel
	for _, datum := range data {
		if datum == nil {
			continue
		}
		cart := &model.CartModel{}
		err = json.Unmarshal([]byte(datum.(string)), cart)
		if err != nil {
			log.Warnf("[repo.cart] json.unmarshal err: %v", err)
			continue
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

//DelCart 移除购物车
func (r *Repo) DelCart(ctx context.Context, userID int, ids []string) error {
	return redis.Client.HDel(ctx, model.BuildCartKey(userID), ids...).Err()
}

//EmptyCart 清空购物车
func (r *Repo) EmptyCart(ctx context.Context, userID int) error {
	return redis.Client.Del(ctx, model.BuildCartKey(userID)).Err()
}

//CartList 我的购物车
func (r *Repo) CartList(ctx context.Context, userID int) ([]*model.CartModel, error) {
	data, err := redis.Client.HGetAll(ctx, model.BuildCartKey(userID)).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.cart] json marshal")
	}
	carts := make([]*model.CartModel, 0, len(data))
	for _, v := range data {
		cart := &model.CartModel{}
		err = json.Unmarshal([]byte(v), cart)
		if err != nil {
			log.Warnf("[repo.cart] list json unmarshal err:%v", err)
			continue
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
