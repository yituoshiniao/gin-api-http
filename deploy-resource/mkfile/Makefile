
#branch =v$(shell (date "+%Y.%m.%d.%H.%M.%S"))
#branch =v$(shell (date "+%Y.%m.%d.%H.%M"))

#当前项目目录名字
dirName := $(shell basename $(PWD))
#docker hub 仓库地址
docker_hub_url = "hub.docker.com"


#版本号
#branch = test
envPath = cs_healthy/$(dirName)/$(env)
dockerName = $(dirName)-$(env)
gateWayHttpPort = 3013
gateWayAdminPort = 6013
branchGateWayPath=/apps/conf/gitlab/$(dirName)

demo:
	@echo  "branch:="  $(branch)
	@echo  "envPath:"  $(envPath)
	@echo  "dockerName:" $(dockerName)


#构建linux环境包
git-pull:
ifeq ($(env),test)
	#test环境使用branch 所以需要git pull 更新代码
	cd $(branchGateWayPath) && git add . && git reset --hard && git fetch -t && git checkout $(branch) && git pull
else
	cd $(branchGateWayPath) && git add . && git reset --hard && git fetch -t && git checkout $(branch)
endif




##构建gateway
build-image:
	cd $(branchGateWayPath) && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o app $(branchGateWayPath)/cmd/server/main.go;
	cd $(branchGateWayPath) && docker build -t $(docker_hub_url)/$(envPath):$(branch) -f Dockerfile $(branchGateWayPath);
	@echo "构建成功版本:"$(branch) " 环境: " $(env)


#推送镜像image到hub
push-image:
	docker push $(docker_hub_url)/$(envPath):$(branch)
	@echo "推送成功版本:"$(branch) " 环境: " $(env)


pull-image:
	docker pull $(docker_hub_url)$(envPath):$(branch)
	@echo "拉取成功版本:"$(branch) " 环境: " $(env)

#测试环境pull + 登录 docker login
docker_uname = "'robot$$cs_healthy+jenkins'"
docker_pwd = "YToo5noykvHFrjDUnSCoe6onWoqSXI0A"
pull-image-v2:
	docker login -u '$(docker_uname)' -p $(docker_pwd)  $(docker_hub_url) \
	docker pull $(docker_hub_url)/$(envPath):$(branch)
	@echo "拉取成功版本:"$(branch)


run-dir:
	mkdir -p /apps/log
	chmod -R 0777 /apps/log
	mkdir -p /apps/conf/healthy_gateway/config

run-image:
	docker stop $(dockerName); \
	docker rm $(dockerName); \
	docker run --name $(dockerName)  -d -p $(gateWayHttpPort):$(gateWayHttpPort) -p $(gateWayAdminPort):$(gateWayAdminPort) \
		--restart always \
		-e envPath=$(env) \
		--log-driver=syslog \  #启用syslog
		#将日志同步到此 syslog-address 路径
		--log-opt syslog-address=udp://10.2.8.203:10035 \
		-v /apps/log:/apps/log \
		#-v /apps/conf/goods-center-logic/config:/apps/conf/goods-center-logic \ #共享本地配置挂载
		$(docker_hub_url)/$(envPath):$(branch) \
	@echo "启动成功 容器名:"$(dockerName) " 环境: " $(env)


#启动 api
deploy: pull-image run-image
