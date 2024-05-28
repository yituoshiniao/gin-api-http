#版本号
#branch = test

#当前项目目录名字
dirName := $(shell basename $(PWD))
envPath = cs_healthy/$(dirName)/$(env)
dockerName = $(dirName)-$(env)
currentBranch = $(shell git symbolic-ref HEAD | sed -e 's,.*/\(.*\),\1,')
find_str = "github_com_yituoshiniao_gin-api-http_"
replace_str = ""
swag_path = "gen/swag-doc"

swagger_file = "/$(dirName)/$(swag_path)/swagger/swagger.json"
api_html_file = "/$(dirName)/gen/html2/index.html"
api_tmp_html_file = "/$(dirName)/gen/html2/api.html"
SRCS = $(shell git ls-files '*.go')

#docker hub 仓库地址
docker_hub_url = "hub.docker.com"



pwd:
	@echo $(dirName)

#branch =v$(shell (date "+%Y.%m.%d.%H.%M.%S"))
#branch =v$(shell (date "+%Y.%m.%d.%H.%M"))

#运行一下所有命令
all:wire db doc

#一键生成doc快捷命令; [1、生成swag； 2、替换多余字符串;]
doc: doc-swag replace-swag-json

#生成openapi sdk 客户端
gen-openapi:openapi-doc-docker-go openapi-doc-docker-lua openapi-doc-html openapi-documentation openapi-doc-cwiki gen-subsplit




demo:
	echo  $(branch)
	echo  $(envPath)
	echo  $(dockerName)
	echo  $(docker_uname)
	echo  $(docker_pwd)


#测试环境pull + 登录 docker login
docker_uname = 'robot$$cs_healthy+jenkins'
docker_pwd = "YToo5noykvHFrjDUnSCoe6onWoqSXI0A"

pull-image-v2:
	docker login -u '$(docker_uname)' -p $(docker_pwd)  $(docker_hub_url) \
	docker pull $(docker_hub_url)/$(envPath):$(branch)
	@echo "拉取成功版本:"$(branch)

#构建linux环境包
build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o app cmd/server/main.go

#构建mac 环境包
build-mac:
	 go build  -o app cmd/server/main.go

#test 构建项目 创建镜像image
build-image:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o app cmd/server/main.go;\
    docker login -u $(docker_uname) -p $(docker_pwd)  $(docker_hub_url); \
	docker build  -t  $(docker_hub_url)/$(envPath):$(branch)  -f Dockerfile . ;  \
	rm app
	@echo "构建成功版本:"$(branch)



login:
	echo '$(docker_uname)'
	echo "$(docker_uname)"
	echo $(docker_uname)
	docker login -u $(docker_uname) -p $(docker_pwd)  $(hub.docker.com)

#拉取镜像image到本地
pull-image:
	docker login -u $(docker_uname) -p $(docker_pwd)  $(docker_hub_url) \
	docker pull $(docker_hub_url)/$(envPath):$(branch)
	@echo "拉取成功版本:"$(branch)

#推送镜像image到hub
push-image:
	docker push $(docker_hub_url)/$(envPath):$(branch)
	@echo "推送成功版本:"$(branch)


#1、构建镜像; 2 推送镜像到hub
test: build-image push-image


#日志上报配置详见： https://doc.abc.net/pages/viewpage.action?pageId=350650880
#运行
run-local:
	docker stop $(dockerName); \
	docker rm $(dockerName); \
	docker run --name $(dockerName)   \
		-e envPath=$(env) \
		-d -p 3013:3013 -p 6013:6013 \
		--log-driver=syslog \
		--log-opt syslog-address=udp://10.2.8.203:10058 \
	  	--restart always $(docker_hub_url)/$(envPath):$(branch)
	@echo "启动成功 容器名:"$(dockerName)



#启动gateway
deploy:  run-local


##********生成swagger文档 ****##
#merge 有三种导入方式(v1.3.23+支持) normal, good, mergin
#普通模式(normal)：不导入已存在的接口；
#智能合并(good)：已存在的接口，将合并返回数据的 response，适用于导入了 swagger 数据，保留对数据结构的改动；
#完全覆盖(mergin)：不保留旧数据，完全使用新数据，适用于接口定义完全交给后端定义， 默认为 normal

#--parseDependency --parseInternal外部依赖struct解析 https://blog.csdn.net/lsjweiyi/article/details/122948171
#parseDepth 解析依赖深度, 默认100;
#swag swag-doc 文档参数： https://blog.csdn.net/qq_41630102/article/details/128411210
#--tags="iap_apple" 生成那些tag的文档,多个用逗号[,]分割
#   --generatedTime                        是否在 docs.go 顶部生成时间戳，默认 false

doc-swag:
	#swag init -g cmd/server/main.go -o gen/swag-doc/swagger
	swag fmt
	swag init   -md ./$(swag_path)/md  \
		--parseDependency \
		--parseInternal=true \
		--parseGoList=true \
		--parseDepth=300 \
		-g cmd/server/main.go \
		--generatedTime=false \
		-o $(swag_path)/swagger
	sleep 1
	#yapi-import-tmp.json yapi 服务器配置，其中server配置是服务端地址
	yapi import  || true
	@echo "swagger 成功"

#copy-doc:
#	rm -rf ./gen/swagger
#	cp -rf ./$(swag_path)/swagger ./gen #拷贝swagger到gen目录下
#	@echo "success 成功"





#替换 find_str 字符串为 replace_str 字符串
replace-swag-json:
	$(shell (find ./$(swag_path)/swagger -iname "*.json") | xargs -n1 -IX bash -c 'sed s/$(find_str)/$(replace_str)/ X > X.tmp && mv X{.tmp,}')
	$(shell (find ./$(swag_path)/swagger -iname "*.go") | xargs -n1 -IX bash -c 'sed s/$(find_str)/$(replace_str)/ X > X.tmp && mv X{.tmp,}')
	$(shell (find ./$(swag_path)/swagger -iname "*.yaml") | xargs -n1 -IX bash -c 'sed s/$(find_str)/$(replace_str)/ X > X.tmp && mv X{.tmp,}')
	@echo "swagger 成功"


#支持ssh地址,当不支持ssh地址可以使用http
#git subsplit publish  --no-tags --heads=$(currentBranch) gen:git@gitlab.abc.net:cs-server2/services/openapi-client.git

#subtree 更新管理子分支
#详情：https://github.com/dflydev/git-subsplit
gen-subsplit:
	@echo "当前git分支 $(currentBranch)"
	rm -rf .subsplit
	git subsplit init .
	git subsplit publish  --no-tags --heads=$(currentBranch) gen:https://github.com/yituoshiniao/openapi-client.git #不支持ssh,使用它http同步
	git subsplit publish  --no-tags --heads=$(currentBranch) gen/go:https://github.com/yituoshiniao/openapi-client-go.git
	#git subsplit publish  --no-tags --heads=$(currentBranch) gen/lua:https://github.com/yituoshiniao/openapi-client-lua.git
	rm -rf .subsplit
	@echo "执行成功"



# 生成 openapi-generator 生成多语言sdk服务端(本地命令)
#详情：https://openapi-generator.tech/docs/generators
openapi-doc-server:
	rm -rf ./$(swag_path)/swagger/go
	npx @openapitools/openapi-generator-cli generate -i ./$(swag_path)/swagger/swagger.json \
		-g go-gin-server \
	 	-o ./gen/go-gin-server
		@echo "执行成功"


#生成 Confluence wiki 文档
out_path_cwiki =  "/gen/cwiki"
openapi-doc-cwiki:
	docker run --rm \
		-v ${PWD}:/$(dirName) openapitools/openapi-generator-cli generate \
		-i $(swagger_file) \
		-g cwiki \
	 	-o .$(out_path_cwiki)
		@echo "执行成功"


#生成html文档 #-c "./$(swag_path)/swagger/config.json"
out_path_html =  "/$(dirName)/gen/html2"
openapi-doc-html:
	rm -rf .$(out_path_html)
	docker run --rm \
		-v ${PWD}:/$(dirName) openapitools/openapi-generator-cli generate \
		-i $(swagger_file) \
		-g html2 \
	 	-o $(out_path_html)
		@echo "执行成功"

openapi-documentation-openapi:
	npx @openapitools/openapi-generator-cli generate \
 		-i ./$(swag_path)/swagger/swagger.json \
		-g openapi-yaml	\
	 	-o ./gen/openapi-yaml

#同步markdown文档到sdk的client 中
out_path_md =  "/$(dirName)/gen/markdown"
openapi-documentation:
	docker run --rm \
		-v ${PWD}:/$(dirName) openapitools/openapi-generator-cli generate \
		-i $(swagger_file) \
		-g markdown	\
	 	-o $(out_path_md)
	cp -r ./gen/markdown/* ..$(out_lua_path) && cp -rf ..$(swagger_file) ..$(out_lua_path)  && rm -rf ./gen/markdown
	#复制 swagger.json 接口文档到 sdk client中
	cp -rf ..$(swagger_file) ..$(out_go_path) && cp -rf ..$(swagger_file) ..$(out_go_path)
	#复制html接口文档到 sdk client中
	cp ..$(api_html_file) ..$(api_tmp_html_file)
	cp -rf ..$(api_tmp_html_file) ..$(out_go_path) && cp -rf ..$(api_tmp_html_file) ..$(out_lua_path) && rm -rf ..$(api_tmp_html_file)
	@echo "执行success"


# 生成 openapi-generator 生成多语言sdk客户端(本地命令)
openapi-doc:
	rm -rf ./$(swag_path)/swagger/go
	npx @openapitools/openapi-generator-cli generate \
		-i ./$(swag_path)/swagger/swagger.json \
		-g go \
	 	-o ./api/swagger/go_bak


#git 仓库host
git-host = "github.com"
#git仓库host后的路径
git-user-id = "yituoshiniao"

##指定运行容器版本  -v ${PWD}:$(dirName) openapitools/openapi-generator-cli:v7.5.0 generate \
#通过docker生成 openapi-generator 生成多语言sdk客户端
# -i petstore.yaml 表示要根据哪个文档生成，支持yaml和json
# -g ruby          表示生成ruby client
# -o /tmp/test/    指定存放的地址

#go 生成 openapi客户端路径
out_go_path  = "/$(dirName)/gen/go"
openapi-doc-docker-go:
	rm -rf ..$(out_go_path)
	docker run --rm \
		-v ${PWD}:/$(dirName) openapitools/openapi-generator-cli generate \
		-i $(swagger_file) \
		-g go \
		--git-host $(git-host) \
		--git-user-id $(git-user-id)  \
		--git-repo-id "openapi-client-go" \
		--release-note "openapi生成go客户端demo" \
		-o $(out_go_path)
	#更新go mod
	cd ..$(out_go_path) && go mod tidy
	@echo "执行成功"




out_lua_path  = "/$(dirName)/gen/lua"
openapi-doc-docker-lua:
	rm -rf ..$(out_lua_path)
	docker run --rm \
		-v ${PWD}:/$(dirName) openapitools/openapi-generator-cli generate \
		-i $(swagger_file) \
		--git-host $(git-host) \
		--git-user-id $(git-user-id) \
		--git-repo-id "openapi-client-lua" \
		--release-note "openapi生成-lua客户端demo" \
		-g lua \
		-o $(out_lua_path)
	@echo "执行成功"




out_java_path = "/$(dirName)/gen/java"
openapi-doc-docker-java:
	rm -rf ..$(out_java_path)
	docker run --rm \
	  -v ${PWD}:/$(dirName) openapitools/openapi-generator-cli generate \
	  -i $(swagger_file) \
		--git-host $(git-host) \
		--git-user-id $(git-user-id) \
		--git-repo-id "openapi-client-java" \
		--release-note "openapi生成-java客户端demo" \
		-g java \
		-o $(out_java_path)
	@echo "执行成功"

#拉取代码
#make git-pull  env=test branch=demo
#构建镜像
#make build-image  env=test branch=demo
#推送镜像
#make push-image  env=test branch=demo
#部署镜像
#make deploy  env=test branch=demo


#gentool使用doc:https://github.com/go-gorm/gen/tree/master/tools/gentool
#安装 gentool: go install gorm.io/gen/tools/gentool@latest
#tables 指定生成的表,空则全部表生成
#onlyModel 只生成 model模型，不生成查询query
# -outPath "dao/a"  指定生产文件目录,默认 更项目的dao目录
#-tables 指定生成的表明,生成全部 则表留空即可
# -withUnitTest 为查询代码生成单元测试

#生成gorm运行文件
gentool:
	gentool -dsn "root:passw0rd@tcp(127.0.0.1:3306)/db_goods_center?charset=utf8mb4&parseTime=True&loc=Local" \
	-tables "apple_product_price,apple_config,user_score" \
   	-fieldWithTypeTag \
   	-fieldWithIndexTag \
   	-modelPkgName "model"

#自定义 gorm Generator 生成db文件
db:
	go run cmd/gen/main.go

#wire执行
wire:
	cd ./inject/ && wire

## Format the Code
gofmt:
	gofmt -s -l -w $(SRCS)
