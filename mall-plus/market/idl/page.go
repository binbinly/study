package idl

import (
	"encoding/json"

	"go-micro.dev/v4/logger"

	pb "common/proto/market"
	"common/util"
	"market/conf"
	"market/model"
)

//TransferAppSetting 转换页面配置数据
func TransferAppSetting(list []*model.AppSettingModel) (res []*pb.AppSetting) {
	if len(list) == 0 {
		return []*pb.AppSetting{}
	}

	for _, setting := range list {
		switch setting.Type {
		case model.AppTypeSwiper, model.AppTypeThreeAdv:
			var images []string
			if err := json.Unmarshal([]byte(setting.Data), &images); err != nil {
				logger.Warnf("[idl.app] images json unmarshal err: %v", err)
				continue
			}
			res = append(res, &pb.AppSetting{
				Type: int32(setting.Type),
				Data: &pb.AppSetting_Images{Images: &pb.SettingImages{List: buildImages(images)}},
			})
		case model.AppTypeNav:
			var navs []*pb.SettingNav
			if err := json.Unmarshal([]byte(setting.Data), &navs); err != nil {
				logger.Warnf("[idl.app] navs json unmarshal err: %v", err)
				continue
			}
			for _, nav := range navs {
				nav.Icon = util.BuildResUrl(conf.Conf.DFS, nav.Icon)
			}
			res = append(res, &pb.AppSetting{
				Type: int32(setting.Type),
				Data: &pb.AppSetting_Navs{Navs: &pb.SettingNavs{List: navs}},
			})
		case model.AppTypeOneAdv:
			ads := &pb.SettingAds{}
			if err := json.Unmarshal([]byte(setting.Data), ads); err != nil {
				logger.Warnf("[idl.app] ads json unmarshal err: %v", err)
				continue
			}
			ads.Cover = util.BuildResUrl(conf.Conf.DFS, ads.Cover)
			res = append(res, &pb.AppSetting{
				Type: int32(setting.Type),
				Data: &pb.AppSetting_Ads{Ads: ads},
			})
		case model.AppTypeProduct:
			product := &pb.SettingProduct{}
			if err := json.Unmarshal([]byte(setting.Data), product); err != nil {
				logger.Warnf("[idl.app] product json unmarshal err: %v", err)
				continue
			}
			res = append(res, &pb.AppSetting{
				Type: int32(setting.Type),
				Data: &pb.AppSetting_Product{Product: product},
			})
		}
	}
	return
}

//TransferAppNotice 转换公告数据
func TransferAppNotice(list []*model.AppNoticeModel) (res []*pb.Notice) {
	if len(list) == 0 {
		return []*pb.Notice{}
	}

	for _, notice := range list {
		res = append(res, &pb.Notice{
			Id:        notice.ID,
			Title:     notice.Title,
			Content:   notice.Content,
			CreatedAt: notice.CreatedAt,
		})
	}
	return
}

//buildImages 构建图片组路径
func buildImages(images []string) (res []string) {
	for _, image := range images {
		res = append(res, util.BuildResUrl(conf.Conf.DFS, image))
	}
	return
}
