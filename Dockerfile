FROM smgano/golang-base:v1



RUN mkdir -p /apps/webroot/production/gin-api-http \
&& mkdir  -p /apps/conf/gin-api-http   \
&& mkdir  -p /apps/log/app/gin-api-http

#拷贝文件<源文件  目标>
ADD app /apps/webroot/production/gin-api-http/app
RUN chmod +x /apps/webroot/production/gin-api-http/app \
&& touch /apps/log/app/gin-api-http/test.txt \
&& chmod -R +rwx /apps/log/app/gin-api-http


#工作目录
WORKDIR /apps/webroot/production/gin-api-http

COPY config/* /apps/conf/gin-api-http/

ENV PATH=/apps/webroot/production/gin-api-http:$PATH

#ARG GOTMP=test111

ENV envPath = ""
#ENV input1=""

#测试环境

#不支持动态命令替换
#ENTRYPOINT ["/apps/webroot/production/gin-api-http/app" , "-envPath=$envPath"]

#shell 模式 支持动态命令替换
#CMD /apps/webroot/production/gin-api-http/app -envPath=${envPath}
ENTRYPOINT /apps/webroot/production/gin-api-http/app -envPath=${envPath}

#ENTRYPOINT ["tail", "-f", "/apps/log/app/gin-api-http/test.txt"]


EXPOSE 3013 6013