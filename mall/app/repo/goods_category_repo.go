package repo

import (
	"context"
	"sort"

	"github.com/pkg/errors"

	"mall/app/model"
)

//GoodsCategoryTree 商品所有分类-树形结构
func (r *Repo) GoodsCategoryTree(ctx context.Context) ([]*model.GoodsCategoryTree, error) {
	list, err := r.categoryAll(ctx)
	if err != nil {
		return nil, err
	}
	data := buildData(list)
	tree := makeTreeCore(0, data)
	return tree, nil
}

//GoodsCategoryChild 获取分类下所有子分类
func (r *Repo) GoodsCategoryChild(ctx context.Context, id int) ([]int, error) {
	list, err := r.categoryAll(ctx)
	if err != nil {
		return nil, err
	}
	return makeChild(id, list), nil
}

//categoryAll 获取全部分裂
func (r *Repo) categoryAll(ctx context.Context) (list []*model.GoodsCategoryModel, err error) {
	err = r.db.WithContext(ctx).Model(&model.GoodsCategoryModel{}).Select("id,pid,name,sort").Order(model.DefaultOrderSort).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "[repo.category] get all")
	}
	return
}

//buildData pid => id => data
func buildData(list []*model.GoodsCategoryModel) map[int]map[int]*model.GoodsCategoryTree {
	var data = make(map[int]map[int]*model.GoodsCategoryTree)
	for _, v := range list {
		id := v.ID
		pid := v.PID
		if _, ok := data[pid]; !ok {
			data[pid] = make(map[int]*model.GoodsCategoryTree)
		}
		data[pid][id] = &model.GoodsCategoryTree{
			ID:    id,
			PID:   pid,
			Name:  v.Name,
			Sort:  v.Sort,
			Child: make([]*model.GoodsCategoryTree, 0),
		}
	}
	return data
}

//makeTreeCore 递归移动
func makeTreeCore(index int, data map[int]map[int]*model.GoodsCategoryTree) []*model.GoodsCategoryTree {
	tmp := make([]*model.GoodsCategoryTree, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Child = makeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	//重新排序
	sort.Sort(model.GoodsCategorySort(tmp))
	return tmp
}

//makeChild 递归所有子分类
func makeChild(id int, list []*model.GoodsCategoryModel) (ids []int) {
	for _, categoryModel := range list {
		if id == 0 { //所有子分类
			ids = append(ids, categoryModel.ID)
			continue
		}
		if categoryModel.PID == id { //递归查找
			ids = append(ids, makeChild(categoryModel.ID, list)...)
		}
	}
	if id > 0 {
		ids = append(ids, id)
	}
	return ids
}
