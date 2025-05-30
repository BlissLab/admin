import service from '@/utils/request'
// @Tags File
// @Summary 创建文件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.File true "创建文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /file/createFile [post]
export const createFile = (data) => {
  return service({
    url: '/file/createFile',
    method: 'post',
    data
  })
}

// @Tags File
// @Summary 删除文件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.File true "删除文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /file/deleteFile [delete]
export const deleteFile = (params) => {
  return service({
    url: '/file/deleteFile',
    method: 'delete',
    params
  })
}

// @Tags File
// @Summary 批量删除文件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /file/deleteFile [delete]
export const deleteFileByIds = (params) => {
  return service({
    url: '/file/deleteFileByIds',
    method: 'delete',
    params
  })
}

// @Tags File
// @Summary 更新文件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.File true "更新文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /file/updateFile [put]
export const updateFile = (data) => {
  return service({
    url: '/file/updateFile',
    method: 'put',
    data
  })
}

// @Tags File
// @Summary 用id查询文件
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.File true "用id查询文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /file/findFile [get]
export const findFile = (params) => {
  return service({
    url: '/file/findFile',
    method: 'get',
    params
  })
}

// @Tags File
// @Summary 分页获取文件列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /file/getFileList [get]
export const getFileList = (params) => {
  return service({
    url: '/file/getFileList',
    method: 'get',
    params
  })
}

// @Tags File
// @Summary 不需要鉴权的文件接口
// @Accept application/json
// @Produce application/json
// @Param data query voiceReq.FileSearch true "分页获取文件列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /file/getFilePublic [get]
export const getFilePublic = () => {
  return service({
    url: '/file/getFilePublic',
    method: 'get',
  })
}
