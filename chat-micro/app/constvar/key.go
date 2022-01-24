package constvar

import "fmt"

const (
	// _onlinePrefix 在线key前缀
	_onlinePrefix = "user:online:"
	// _userPrefix 用户令牌标识 用于单点登录
	_userPrefix = "user:token:"
	//_historyPrefix 离线消息前缀
	_historyPrefix = "history:message:%d"
)

// BuildHistoryKey 历史消息键
func BuildHistoryKey(userID uint32) string {
	return fmt.Sprintf(_historyPrefix, userID)
}

// BuildOnlineKey 用户在线键
func BuildOnlineKey(userID uint32) string {
	return fmt.Sprintf("%s%d", _onlinePrefix, userID)
}

// BuildUserTokenKey 用户令牌键
func BuildUserTokenKey(userID uint32) string {
	return fmt.Sprintf("%s%d", _userPrefix, userID)
}
