package cron

import "context"

type TemplateData interface {
	Handle() error
	Name() string
}

var Template = make(map[string]TemplateData)

// FactoryTemplateSrv ...
type FactoryTemplateSrv struct {
	testSrv *TestSrv
}

func NewFactoryTemplateSrv(
	testSrv *TestSrv,

) *FactoryTemplateSrv {
	return &FactoryTemplateSrv{
		testSrv: testSrv,
	}
}

// RegisterTemplate 注册模板 ...
func (e *FactoryTemplateSrv) RegisterTemplate(ctx context.Context) {
	once.Do(func() {
		Template[e.testSrv.Name()] = e.testSrv
		//Template[e.testSrv.Name()] = NewTestSrv(ctx)
	})
}
