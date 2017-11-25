Vagrant.configure("2") do |config|

    (2..2).each do |i|
		config.vm.box_check_update = false
        config.vm.define "knode#{i}" do |node|

        # 设置虚拟机的Box
        #node.vm.box = "ubuntu/trusty64"
        node.vm.box = "ubuntu/xenial64"

        # 设置虚拟机的主机名
        node.vm.hostname="knode#{i}"

        # 设置虚拟机的IP
        node.vm.network "private_network", ip: "192.168.59.#{i}"

        # 设置主机与虚拟机的共享目录
        node.vm.synced_folder "/Volumes/256G/share", "/home/vagrant/share"

        # VirtaulBox相关配置
        node.vm.provider "virtualbox" do |v|

            # 设置虚拟机的名称
            v.name = "node#{i}"

            # 设置虚拟机的内存大小  
            v.memory = 4096

            # 设置虚拟机的CPU个数
            v.cpus = 2
        end
  
        # 使用shell脚本进行软件安装和配置
        node.vm.provision "shell", inline: <<-SHELL
 echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCnvlNF7kGBPbwXPL7hEsDjw4geyFtQNo5uKAzVPNa2beStAqdP3HuYYB5D67t+tPr1YpCX6vYWlxMeS5Xog+eBIcd74MPrin1iRGzDUGAU+3soTGpUDG+jDZ2E0zGGCyczW7vAX+cfmLFa+D2uWOFrwxJ6219/QFi5jQgkCBMtE1V6C4vpNLTs+X/q847h6KtV48qwDkfQ+w1S+nqWpioSIF9BaupBnbtqDGFX6spsb4Ujrwu4lZHs/352nTa4OKy3LbEeX8z1y1KstS254ZeaOSSStt2gKruEaBT22e7zln8iikOvSw0cmid/n+vT6zpNUNkfZ/RW6vPs5kO/fXfP dean@DeanMacbook.local" >> ~/.ssh/authorized_keys
SHELL
        end
    end
end
