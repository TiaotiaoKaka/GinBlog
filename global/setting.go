package global

import (
	"github.com/TiaotiaoKaka/GinBlog/pkg/logger"
	"github.com/TiaotiaoKaka/GinBlog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
