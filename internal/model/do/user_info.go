// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta   `orm:"table:user_info, do:true"`
	UserId   interface{} //
	Nickname interface{} //
	Avatar   interface{} //
	Type     interface{} //
	Gender   interface{} //
}
