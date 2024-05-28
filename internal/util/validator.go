package util

import (
	"context"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/yituoshiniao/kit/xlog"
)

var trans ut.Translator

// 设置验证器中文翻译
func init() {
	if trans != nil {
		return
	}
	zhLoc := zh.New()
	enLoc := en.New()
	uni := ut.New(enLoc, zhLoc)
	trans, _ = uni.GetTranslator("zh") // 设置语言
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := zhtranslations.RegisterDefaultTranslations(v, trans); err != nil {
			xlog.S(context.Background()).Infow("验证信息翻译错误", "err", err)
		}
	}
}

// TransError 翻译验证器错误提示
func TransError(err error) map[string]string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		return errs.Translate(trans)
	}
	xlog.S(context.Background()).Infow("TransError-验证信息翻译错误", "err", err)
	return nil
}

// TransErrorStr 翻译验证器错误提示将map转换成string字符串
func TransErrorStr(err error) string {
	errStr := ""
	if errs, ok := err.(validator.ValidationErrors); ok {
		errTrans := errs.Translate(trans)
		length := len(errTrans)
		tmpStr := ""
		if length > 1 {
			tmpStr = "\n" // 多个错误，换行展示
		}
		var i int = 1
		for _, v := range errTrans {
			if i == length {
				tmpStr = "" // 最后一行错误，去除换行
			}
			errStr += v + tmpStr
			i++
		}
		return errStr
	}
	xlog.S(context.Background()).Infow("TransError-验证信息翻译错误", "err", err)
	return errStr
}
