package i18n

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// GetI18nCopy 获取真实文案
func GetI18nCopy(i18nKey string) string {
	return ginI18n.MustGetMessage(i18nKey)
}

// GetI18nCopyWithParam 有参数的
func GetI18nCopyWithParam(i18nKey string, paramMap map[string]interface{}) string {
	return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID:    i18nKey,
		TemplateData: paramMap,
	})
}
