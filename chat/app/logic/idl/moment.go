package idl

import "chat/app/logic/model"

//TransferMomentInput 朋友圈对外转化结构
type TransferMomentInput struct {
	Moments     []*model.MomentModel
	Users       []*model.UserModel
	LikeList    map[uint32]*[]uint32
	CommentList map[uint32]*[]*model.MomentCommentModel
}

// TransferMomentList 组装数据并输出
// 对外暴露的moment结构，都应该经过此结构进行转换
func TransferMomentList(input *TransferMomentInput) []*model.MomentList {
	um := usersToMap(input.Users)
	likes := likesToMap(input.LikeList, um)
	comments := commentsToMap(input.CommentList, um)
	list := make([]*model.MomentList, 0)
	for _, moment := range input.Moments {
		if user, ok := um[moment.UserID]; ok {
			m := &model.MomentList{
				ID:        moment.ID,
				Content:   moment.Content,
				Image:     moment.Image,
				Video:     moment.Video,
				Location:  moment.Location,
				Type:      moment.Type,
				CreatedAt: moment.CreatedAt,
				User:      user,
				Likes:     make([]*model.User, 0),
				Comments:  make([]*model.Comment, 0),
			}
			if l, o := likes[moment.ID]; o {
				m.Likes = l
			}
			if c, o := comments[moment.ID]; o {
				m.Comments = c
			}
			list = append(list, m)
		}
	}
	return list
}

// 点赞列表数据格式化为map
func likesToMap(likes map[uint32]*[]uint32, users map[uint32]*model.UserBase) map[uint32][]*model.User {
	ml := make(map[uint32][]*model.User)
	if len(likes) == 0 {
		return ml
	}
	for mid, like := range likes {
		for _, uid := range *like {
			if user, ok := users[uid]; ok {
				u := &model.User{
					ID:   user.ID,
					Name: user.Name,
				}
				if _, o := ml[mid]; o {
					ml[mid] = append(ml[mid], u)
				} else {
					ml[mid] = []*model.User{u}
				}
			}
		}
	}
	return ml
}

// 评论列表格式化为map
func commentsToMap(mComments map[uint32]*[]*model.MomentCommentModel, users map[uint32]*model.UserBase) map[uint32][]*model.Comment {
	ml := make(map[uint32][]*model.Comment)
	if len(mComments) == 0 {
		return ml
	}
	for _, comments := range mComments {
		for _, comment := range *comments {
			if user, ok := users[comment.UserID]; ok {
				ct := &model.Comment{
					Content: comment.Content,
					User: &model.User{
						ID:   user.ID,
						Name: user.Name,
					},
				}
				if comment.ReplyID > 0 { // 格式化回复者
					if u, o := users[comment.ReplyID]; o {
						ct.Reply = &model.User{
							ID:   u.ID,
							Name: u.Name,
						}
					}
				}
				if _, o := ml[comment.MomentID]; o {
					ml[comment.MomentID] = append(ml[comment.MomentID], ct)
				} else {
					ml[comment.MomentID] = []*model.Comment{ct}
				}
			}
		}
	}
	return ml
}
