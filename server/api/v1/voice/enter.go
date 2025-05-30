package voice

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ FileApi }

var fileService = service.ServiceGroupApp.VoiceServiceGroup.FileService
