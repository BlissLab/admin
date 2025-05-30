package voice

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileRouter struct {}

// InitFileRouter 初始化 文件 路由信息
func (s *FileRouter) InitFileRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	fileRouter := Router.Group("file").Use(middleware.OperationRecord())
	fileRouterWithoutRecord := Router.Group("file")
	fileRouterWithoutAuth := PublicRouter.Group("file")
	{
		fileRouter.POST("createFile", fileApi.CreateFile)   // 新建文件
		fileRouter.DELETE("deleteFile", fileApi.DeleteFile) // 删除文件
		fileRouter.DELETE("deleteFileByIds", fileApi.DeleteFileByIds) // 批量删除文件
		fileRouter.PUT("updateFile", fileApi.UpdateFile)    // 更新文件
	}
	{
		fileRouterWithoutRecord.GET("findFile", fileApi.FindFile)        // 根据ID获取文件
		fileRouterWithoutRecord.GET("getFileList", fileApi.GetFileList)  // 获取文件列表
	}
	{
	    fileRouterWithoutAuth.GET("getFilePublic", fileApi.GetFilePublic)  // 文件开放接口
	}
}
