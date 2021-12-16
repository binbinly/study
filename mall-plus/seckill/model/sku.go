package model

//SkuModel 秒杀商品
type SkuModel struct {
	ID            int64  `json:"id"`             //商品id
	Price         int    `json:"price"`          //秒杀价格
	Count         int    `json:"count"`          //秒杀数量
	Limit         int64  `json:"limit"`          //个人限购
	OriginalPrice int    `json:"original_price"` //原价
	StartAt       int64  `json:"start_at"`       //开始时间
	EndAt         int64  `json:"end_at"`         //结束时间
	Title         string `json:"title"`          //标题
	Cover         string `json:"cover"`          //封面
	Key           string `json:"key"`            //加密key
}
