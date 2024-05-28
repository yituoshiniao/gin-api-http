package util

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"io/ioutil"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/yituoshiniao/kit/xlog"
	"gorm.io/gorm"
)

type Gob struct {
}

func NewGob() *Gob {
	return &Gob{}
}

const (
	GobOperationName string = "Gob"
)

func (b *Gob) Trace(ctx context.Context, op, key string) opentracing.Span {
	sp, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("%s-%s", GobOperationName, op))
	ext.DBType.Set(sp, "gob")
	sp.SetTag("operation", op)
	sp.SetTag("key", key)
	return sp
}

// 写入二进制数据到磁盘文件
func (g *Gob) Write(ctx context.Context, data interface{}, filename string) error {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		xlog.S(ctx).Errorw("encoder.Encode-err", "err", err)
		return err
	}

	// 如果写入成功，返回空的 error 信息，如果写入失败，返回 error 信息，使用 ioutil.WriteFile写文件，在写入文件之前，我们不需要判断文件是否存在，如果文件不存在，会自动创建文件，如果文件存在，则会覆盖原来的内容。
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		xlog.S(ctx).Errorw("WriteFile-err", "err", err)
		return err
	}

	return nil
}

// 从磁盘文件加载二进制数据
func (g *Gob) Read(ctx context.Context, data interface{}, filename string) error {
	xlog.S(ctx).Infow("Read", "filename", filename)
	sp := g.Trace(ctx, "Read", filename)
	defer sp.Finish()
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		xlog.S(ctx).Errorw("ReadFile-err", "err", err)
		err = gorm.ErrRecordNotFound // 转换成找不到记录，方便业务下游捕获错误处理异常
		return err
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		xlog.S(ctx).Errorw("WriteFile-err", "err", err)
		return err
	}
	return nil
}
