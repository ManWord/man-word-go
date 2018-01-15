package utils

/**
session可以保存一定的信息，这样可以避免繁琐的登录数据携带问题
具体保存哪些东西可以根据业务需要来做
 */

//用户信息
type UserInfo struct {
	Username string
	Password string
}

//购物车信息
type CartInfo struct {
	//	TODO add cart info
	OrderId string //example
}

//session会话
type SessionInfo struct {
	UserInfo
	CartInfo
}
