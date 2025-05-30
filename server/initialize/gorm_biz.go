package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voice"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(voice.File{})
	if err != nil {
		return err
	}
	return nil
}
