
package voice

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voice"
    voiceReq "github.com/flipped-aurora/gin-vue-admin/server/model/voice/request"
)

type FileService struct {}
// CreateFile 创建文件记录
// Author [yourname](https://github.com/yourname)
func (fileService *FileService) CreateFile(ctx context.Context, file *voice.File) (err error) {
	err = global.GVA_DB.Create(file).Error
	return err
}

// DeleteFile 删除文件记录
// Author [yourname](https://github.com/yourname)
func (fileService *FileService)DeleteFile(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&voice.File{},"id = ?",ID).Error
	return err
}

// DeleteFileByIds 批量删除文件记录
// Author [yourname](https://github.com/yourname)
func (fileService *FileService)DeleteFileByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]voice.File{},"id in ?",IDs).Error
	return err
}

// UpdateFile 更新文件记录
// Author [yourname](https://github.com/yourname)
func (fileService *FileService)UpdateFile(ctx context.Context, file voice.File) (err error) {
	err = global.GVA_DB.Model(&voice.File{}).Where("id = ?",file.ID).Updates(&file).Error
	return err
}

// GetFile 根据ID获取文件记录
// Author [yourname](https://github.com/yourname)
func (fileService *FileService)GetFile(ctx context.Context, ID string) (file voice.File, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&file).Error
	return
}
// GetFileInfoList 分页获取文件记录
// Author [yourname](https://github.com/yourname)
func (fileService *FileService)GetFileInfoList(ctx context.Context, info voiceReq.FileSearch) (list []voice.File, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&voice.File{})
    var files []voice.File
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	if info.StartDuration != nil && info.EndDuration != nil {
		db = db.Where("duration BETWEEN ? AND ? ", *info.StartDuration, *info.EndDuration)
	}
    if info.UserId != nil {
        db = db.Where("user_id = ?", *info.UserId)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["ID"] = true
           orderMap["CreatedAt"] = true
         	orderMap["duration"] = true
         	orderMap["user_id"] = true
         	orderMap["status"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&files).Error
	return  files, total, err
}
func (fileService *FileService)GetFilePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
