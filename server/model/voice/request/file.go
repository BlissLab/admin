package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FileSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	StartDuration  *int        `json:"startDuration" form:"startDuration"`
	EndDuration    *int        `json:"endDuration" form:"endDuration"`
	UserId         *int        `json:"userId" form:"userId"`
	Status         *string     `json:"status" form:"status"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
