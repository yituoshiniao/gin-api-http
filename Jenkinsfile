#!groovy
pipeline {
    agent {
        label 'jnlp-slave2'
    }
    parameters {
        choice(name: 'env', choices: [ 'test', 'dev', 'beta', 'prod'], description: '环境')
        choice(name: 'action', choices: ['deploy', 'deploy-in-k8s', 'restart', 'rollback'], description: '动作')
        string(name: 'branch', defaultValue: 'app_store', trim: true, description: '版本/tag')
        //choice(name: 'docker_uname', choices: ['quanlong_gan', 'jingmin_lin'], description: 'docker 登录用户')
        string(name: 'docker_uname', description: 'docker login account', defaultValue:'\'robot$$cs_healthy+jenkins\'')
        password(name: 'docker_pwd', description: 'docker login 秘钥', defaultValue:'YToo5noykvHFrjDUnSCoe6onWoqSXI0A')
        string(name: 'host_list',defaultValue: '10.2.8.212', trim: true, description: '部署机器列表 多台用逗号个该')
        //gitParameter name: 'branch', branchFilter: 'origin/.*', tagFilter: 'v*', defaultValue: 'origin/test', selectedValue: 'DEFAULT', listSize: '10',  type: 'PT_BRANCH_TAG', description: '版本'
    }

    stages {
        // 输出
        stage("检出代码") {
            steps {
                script {
                    sh '(date "+%Y-%m-%d %H:%M:%S")'

                    // sh 'printenv'
                    // printf("任务：%s", env.JOB_NAME)

                    // hosts = params.host_list.split(',')
                    // for ( i in hosts){
                    //     printf("机器ip：%s", i)
                    //     echo "机器ip地址: $i"
                    // }

                    // 分支获取
                    // sh "echo version $params.version"
                    // tmpBranch = params.version.split('/') as List
                    // for ( i in tmpBranch){
                    //     printf("分支：%s", i)
                    // }

                    // sh "echo 构建分支: $tmpBranch"
                    // selectBranch = tmpBranch[tmpBranch.size()-1]
                    // sh "echo 构建分支1 $selectBranch"

                    selectBranch = params.branch
                    sh """
                    echo $params.docker_pwd
                    echo $params.docker_uname
                    pwd


                    """
                    sh "/usr/local/bin/check_shell -H 10.2.8.212 -p 5777 -c pull_pay_center_logic -t 180 -a $params.env $selectBranch"
                }
            }
        }
        stage("构建镜像") {
            steps {
                script {
                    if (params.action == 'deploy') {
                        sh "date"
                        // sh "/usr/local/bin/check_shell -h"
                    	sh "/usr/local/bin/check_shell -H 10.2.8.212 -p 5777 -t 180 -c build_pay_center_logic  -a $params.env $selectBranch $params.docker_uname $params.docker_pwd"
                        sh "date"
                    }
                }
            }
        }
        // 镜像推送到harbor
        stage("推送镜像") {
            steps {
                script {
                    if (params.action == 'deploy') {
                    	sh "/usr/local/bin/check_shell -H 10.2.8.212 -p 5777 -c push_pay_center_logic -t 180 -a $params.env $selectBranch"
                    }
                }
            }
        }
        stage("部署镜像") {
            steps {
                script {
                    if (params.action == 'deploy') {
                        hosts = params.host_list.split(',')
                        for ( i in hosts) {
                            printf("机器ip：%s", i)
                    		sh "/usr/local/bin/check_shell -H $i -p 5777 -c deploy_pay_center_logic -t 180 -a $params.env $selectBranch $params.docker_uname $params.docker_pwd"
                            sh '(date "+%Y-%m-%d %H:%M:%S")'
                        }
                    }
                }
            }
        }
    }
}