






	
ContainerOPS及Componet AutoTestFlow
使用安装文档

目录
1.	背景简介	3
1.1	文档目的	3
1.2	参考资料	4
1.3	表示方法	4
2.	集群安装环境及配置	5
2.1	安装环境描述	5
2.2	创建集群环境	5
2.3	Go语言环境配置	8
2.4	配置Singular部署Kubernetes集群环境	8
3.	Containerops	11
3.1	Assembling 配置部署	11
3.1.1	Postman执行测试	13
3.2	在Pilotage安装配置	13
3.2.1	配置	13
3.2.2	测试	15
3.3	在dockeryard 安装配置	15
4.	组件自动化测试Componet AutoTestFlow	16
4.1	ctest执行流程示意图文档目的	16
4.2	ctest执行流程示意图文档目的	16
4.3	ctest手动测试步骤	16
5.	总结备注	18


1.	背景简介
1.1	文档目的
本文档描述了Containerops人员对问题的理解，问题的当前状态，处理及分析的工作过程，以及后续的工作计划、问题结论及建议等方面的信息。
本文档既用来使最终用户对问题的处理状况有全面的了解，也用来使各种协作人员对问题状态有统一的认识。
1.2	参考资料
1.3	表示方法

2.	集群安装环境及配置
源代码地址：
https://github.com/Huawei/containerops
介质存放路径：
http://hub.hubops.sh
推荐环境：
Node0 Singulaer 服务
Node1 Assimbling 服务
Node2 Pilotage 服务
Node3 K8S Master
Node4 K8S Slaver
Node5 K8S Slaver

演示环境：
Node0 
1）Singulaer Service  
2）Assimbling Service 
3）Pilotage Service  
Node1 K8S Master
Node2 K8S Slaver
Node3 K8S Slaver


2.1	安装环境描述
安装环境：
安装用户：root用户
Linux环境：Ubuntu 16.04 64位

2.2	创建集群环境
在你开始之前，到您要安装的集群的系统上安装VirtualBOX及Vagrant，支持Linux，Win及MAC OS系统

1. 安装VirtualBox
2. 安装Vagrant
3. 下载Box
vagrant box add ubuntu/trusty64
Box相当于虚拟机所依赖的镜像文件。
4. 编辑Vagrantfile
mkdir vagrant-cluster
cd vagrant-cluster
vim Vagrantfile

Vagrantfile如下，可以通过注释理解每个自定义配置的含义：
Vagrant.configure("2") do |config|
    (0..3).each do |i|
        config.vm.define "node#{i}" do |node|
        # 设置虚拟机的Box
        node.vm.box = "ubuntu/trusty64"
        # 设置虚拟机的主机名
        node.vm.hostname="node#{i}"
        # 设置虚拟机的IP
        node.vm.network "private_network", ip: "192.168.59.#{i}"
        # 设置主机与虚拟机的共享目录
        node.vm.synced_folder "~/Desktop/share", "/home/vagrant/share"
        # VirtaulBox相关配置
        node.vm.provider "virtualbox" do |v|
        # 设置虚拟机的名称
            v.name = "node#{i}"
        # 设置虚拟机的内存大小  
            v.memory = 2048
        # 设置虚拟机的CPU个数
            v.cpus = 1
        end
        # 使用shell脚本进行软件安装和配置
        node.vm.provision "shell", inline: <<-SHELL
            # 安装docker 1.11.0
            wget -qO- https://get.docker.com/ | sed 's/docker-engine/docker-engine=1.11.0-0~trusty/' | sh
            usermod -aG docker vagrant   
        SHELL
        end
    end
end
与创建单个虚拟机相比，创建多个虚拟机时多了一层循环，而变量i可以用于设置节点的名称与IP，使用#{i}取值：
(1..3).each do |i|
end
可知，一共创建了3个虚拟机。
5. 在桌面上创建share目录
桌面上的share目录将与虚拟机内的/home/vagrant/share目录内容实时同步
mkdir ~/Desktop/share
6. 创建虚拟机
vagrant up
创建3个虚拟机大概需要5分钟，当然这和机器性能还有网速相关。安装Docker可能会比较慢，不需要的话删除下面几行就可以了：
# 使用shell脚本进行软件安装和配置
node.vm.provision "shell", inline: <<-SHELL
# 安装docker 1.11.0
    wget -qO- https://get.docker.com/ | sed 's/docker-engine/docker-engine=1.11.0-0~trusty/' | sh
    usermod -aG docker vagrant
SHELL

下面是Vagrant虚拟机的配置，可以根据需要进行更改:
•	用户/密码: vagrant/vagrant
•	共享目录: 桌面上的share目录将与虚拟机内的/home/vagrant/share目录内容实时同步
•	内存：4GB
•	CPU: 2
现在已经有四台配置相同的虚机，从Node0~Node3。
下面在node0上面安装go语言及其它服务。
2.3	Go语言环境配置
一、解压缩Go云tar文件
https://golang.org/dl/ go1.9.2.linux-amd64.tar.gz
二、使用如下命令进行安装
tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz
成功后查看/usr/local/go目录

三、在$HOME/.profile或/etc/profile文件中中添加环境变量
export PATH=$PATH:/usr/local/go/bin
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
最后检查GO是否安装成功
go version




2.4	配置Singular部署Kubernetes集群环境

go version
mkdir –p gopath/src/github.com/Huawei/
#在gopath/src/github.com/Huawei/目录下
git clone https://github.com/Huawei/containerops.git
#切换到singular目录执行
go get
#执行go build 构建singular二进制或执行
cd gopath/src/github.com/Huawei/containerops/singular/
go install
#在gopath/src/bin目录下执行k8s集群安装
./singular deploy template singular_template.yml  --verbose –timestamp
#开始执行k8s部署工作，3个节点部署过程约5分钟。




三、配置singular_template.yml

https://github.com/imlidian/installer/singular/ singular_template.yml

#集群配置 包括预制用户root及IP地址
nodes:
  -
    ip: 192.168.59.15
    user: root
    distro: centos
  -
    ip: 192.168.59.16
    user: root
    distro: centos
  -
    ip: 192.168.59.17
    user: root
distro: centos
#SSH连接每台机器的公钥及私钥地址
tools:
  ssh:
    private: /root/.ssh/id_rsa
    public: /root/.ssh/id_rsa.pub


部署中遇到问题可以在以下目录查看日志
~/.containerops/singular/containerops/singular/
3.	Containerops
3.1	Assembling 配置部署
1安装准备

#安装nigix
apt-get install nginx
#修改软件源
vim sources.list
apt-get update
#软连接
ln -s /etc/nginx/sites-available/assembling /etc/nginx/sites-enabled/assembling
etc/nginx/sites-available/assembling
#重新启动nginx
systemctl restart nginx
journalctl -fu nginx
systemctl restart nginx
lsof -i:9000
tail -100f /var/log/nginx/error.log
#配置文件路径
~/.containerops/config
/etc/nginx/nginx.conf
/etc/nginx/sites-enabled
/etc/nginx/sites-available/
access_log /var/log/nginx/access.log;
error_log /var/log/nginx/error.log;

配置文件#assembling.runtime.toml
    #unix mode
    [assembling]
kubeconfig = "/root/.kube/config"
#kubeconfig = "/var/containerops/confs/kubeconfig"
docker_daemon_image = "hub.opshub.sh/containerops/docker:dind-1.11"
mode = "unix"
address = '/root/.containerops/run/assembling.socket'
service_type = 'NodePort'
#https mode
[assembling]
kubeconfig = "/root/gopath/src/github.com/Huawei/containerops/assembling/kubeconfig"
docker_daemon_image = "docker:1.11.2-dind"
mode = "https"
address = '0.0.0.0'
port = 8083
cert = "/etc/containerops/config/containerops.crt"


#执行 assembling daemon star 
./assembling  daemon start –config \ ~/.containerops/config/assembling.runtime.toml
#打印log
tail -f /var/containerops/logs/assembling.log
3.1.1	Postman执行测试
下载地址 https://www.getpostman.com/

3.2	在Pilotage安装配置

3.2.1	配置
	
#运行
pilotage daemon daemon --config ~/.containerops/config/pilotage.toml
#执行后监控  
tail -f /var/containerops/logs/pilotage.log
screen -list
screen -r 24763.pts-4.ubuntu-sfo1
pilotage daemon daemon --config ~/.containerops/config/pilotage.toml
数据库配置

[database]
driver = "mysql"
host = "127.0.0.1"
port = 3306
user = "root"
password = "root"
db = "pilotagedb"

[web]
domain = "opshub.sh"
mode = "http"
address = "127.0.0.1"
port = 8080
[storage]
dockerv2 = "/tmp/dockerv2" # path for image files of Docker Distribution V2 Protocol
binaryv1 = "/tmp/binaryv1" # path for binary files of Dockyard Binary V1 Protocol
[warship]
domain = "hub.opshub.sh"
[singular]
[hook]
host='https://flow.opshub.sh'
flowBaseDir='/var/containerops/flows'

邮件发送配置
[mail]
smtp_address = "smtp.gmail.com"
smtp_port = "587"
user = “xxx@containerops.sh"
password = "xxx"
3.2.2	测试
curl -i -X POST -H 'Content-type':'application/yaml' --data-binary @output.yml \ https://flow.opshub.sh/flow/v1/containerops/python_analysis_coala/flow/latest/yaml
3.3	在dockeryard 安装配置
待完善

4.	组件自动化测试Componet AutoTestFlow

4.1	Ctest自动化执行流程示意

4.2	ctest手动测试步骤

#生成镜像
docker build -t containerops/component-python-coala-workflow:latest .
返回tag 或者image name
#上传hub
docker push [tag]
#验证上传成功
docker pull hub.opshub.sh/containerops/component-python-coala-workflow111:latest

4.2.1	手动测试
curl -XPUT --data-binary @action.yml https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/action1.yml -i

5.	总结备注

在测试过程中 Kubernetes Cluster 会有很多 Pod 产生，隔段时间要删除一下。删除的命令是 
kubectl delete pod `kubectl get pods -a -o=custom-columns=NAME:.metadata.name`


环境清理
docker delete
$ docker rm $(docker ps -a -q)
$ docker ps -a
$ docker rmi $(docker images -q)
$ docker rmi -f $(docker images -q)  解决问题
tar -cvf ./镜像名称.tar -C  路径 .   （最后有个.）

tar -cvf ./component-python-coala-workflow.tar -C  /usr/local/gopath/src/github.com/Huawei/containerops/component/python/analysis/coala .


 
