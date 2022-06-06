# 取消gonivinck容器自启动
sudo docker update --restart=no $(sudo docker ps -a|grep gonivinck|awk '{print $1}')

# 停止gonivinck容器
sudo docker stop $(sudo docker ps -a|grep gonivinck|awk '{print $1}')

# 开启gonivinck容器
sudo docker-compose up -d

# 访问主程序容器
sudo docker exec -it gonivinck_golang_1 bash
# 进入项目目录
cd /usr/src/code