package constvar

import "fmt"

const (
	// _userPrefix 用户令牌标识 用于单点登录
	_userPrefix = "user:token:"

	//HotSearchKey 搜索热词
	HotSearchKey = "search_hot"
)

// BuildUserTokenKey 用户令牌键
func BuildUserTokenKey(userID int) string {
	return fmt.Sprintf("%s%d", _userPrefix, userID)
}