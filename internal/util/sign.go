package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"sort"

	"github.com/yituoshiniao/kit/xlog"
	"go.uber.org/zap"
)

// SignKeySort 签名信息按照map k排序
func SignKeySort(params map[string]string) []string {
	var arrStr []string
	for k := range params {
		arrStr = append(arrStr, k)
	}
	sort.Strings(arrStr)

	return arrStr
}

// GetSign 生成访问签名
func GetSign(ctx context.Context, signParam map[string]string, key string) string {
	keys := SignKeySort(signParam)
	str := ""
	for _, k := range keys {
		str += k + "=" + signParam[k] + "&"
	}
	str += "partnerKey=" + key
	xlog.L(ctx).Info("签名前信息 ", zap.String("str", str))

	md5Ctx := md5.New()
	_, err := md5Ctx.Write([]byte(str))
	if err != nil {
		xlog.S(ctx).Error("GetSign-md5Ctx--err", err)
	}
	sign := hex.EncodeToString(md5Ctx.Sum(nil))
	xlog.L(ctx).Info("签名参数信息", zap.String("originStr", str), zap.String("生成签名", sign))
	return sign
}
