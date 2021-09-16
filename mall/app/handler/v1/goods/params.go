package goods

//SearchParams 搜索参数
type SearchParams struct {
	Cid     int    `json:"cid" binding:"omitempty,numeric"`                       //分类ID
	T       int    `json:"t" binding:"omitempty,numeric"`                         //1=最新，2=最热
	Keyword string `json:"keyword" binding:"omitempty,max=30"`                    //搜索关键词
	Price   string `json:"price" binding:"omitempty,max=30"`                      //价格 0,100
	Field   string `json:"field" binding:"omitempty,oneof=sort price sale_count"` //排序字段
	Order   string `json:"order" binding:"omitempty,oneof=asc desc"`              //排序方式
}
