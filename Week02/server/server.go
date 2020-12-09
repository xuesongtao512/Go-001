// @Desc:
// @CreateTime: 2020/11/29
package server

import (
    "my_go/Week02/dao"
)



// get user id
func GetUserId() (int, error) {
    return dao.SelectUserId()
}