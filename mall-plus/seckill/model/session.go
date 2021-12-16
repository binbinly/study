package model

//SessionModel 秒杀场次
type SessionModel struct {
	ID      int64       `json:"id"`       //id
	Name    string      `json:"name"`     //场次名
	StartAt int64       `json:"start_at"` //开始时间
	EndAt   int64       `json:"end_at"`   //结束时间
	Skus    []*SkuModel `json:"skus"`     //秒杀商品
}
