#!/bin/bash
# 健康检查接口地址
url=http://127.0.0.1:3013/goodsCenterLogic/health
#机器人地址
robotUrl="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=779520ea-b377-43b1-a615-a429884457d4"
date=$(date "+%Y-%m-%d %H:%M:%S")


#获取ip地址 --注意变量的 等于符号("=") 之间不能有空格
#local=$(ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:")
local=$(ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -s "\r\n" "--")

#linux中 在crontab 中可能无法执行 ifconfig, 需要设置环境变量； 或者如下使用完整路径
#/usr/sbin/ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -s "\r\n" ","



# 添加服务状态检查方法
checkHealth(){
#	result=$(curl -s ${url} | grep "ok")
	result=$(curl -s ${url})
	echo "$result"
	if [ "$result" == "ok" ]; then
		return "1"
	else
		return "0"
	fi
}

checkHealth
if [ $? -eq "1" ];then
	echo "server is running" "$date"
else
  	tmp=$(curl -s ${robotUrl} -d  '   {
                                           "msgtype": "markdown",
                                           "markdown": {
                                               "content": "<font color=\"warning\">健康检查失败：</font>
                                               >
                                               > IP地址 : <font color=\"info\"> '${local}'  </font>
                                               >
                                               "
                                           }
                                        }'
    )
	echo "server is wrong" "$date"
fi
