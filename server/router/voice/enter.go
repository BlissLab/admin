package voice

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ FileRouter }

var fileApi = api.ApiGroupApp.VoiceApiGroup.FileApi
