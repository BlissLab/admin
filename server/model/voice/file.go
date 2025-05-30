
// 自动生成模板File
package voice
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 文件 结构体  File
type File struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`  //文件名
  Duration  *int `json:"duration" form:"duration" gorm:"default:0;comment:时长，毫秒;column:duration;" binding:"required"`  //时长
  UserId  *int `json:"userId" form:"userId" gorm:"index;column:user_id;" binding:"required"`  //用户ID
  Status  *string `json:"status" form:"status" gorm:"column:status;" binding:"required"`  //状态
  Path  *string `json:"path" form:"path" gorm:"comment:文件存储路径;column:path;"`  //存储路径
}


// TableName 文件 File自定义表名 file
func (File) TableName() string {
    return "file"
}





