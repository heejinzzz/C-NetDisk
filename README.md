# C-NetDisk
**通过 CLI 工具来使用的简易云盘。**

**支持用户注册、登录、文件上传、下载、删除、重命名等操作。上传、下载文件时显示进度。**

---
### 1. 服务端部署：
提前配置好 nfs 服务器并修改 deploy-kubernetes/pv-mongo.yaml 、deploy-kubernetes/pv-cnetdisk.yaml 中的 nfs Server IP 后，在 kubernetes 集群中的 master 节点上执行：
```shell
cd C-NetDisk
bash deploy.sh
```

---
### 2. 编译 CLI 工具
执行以下指令编译生成 cnetdisk CLI 工具：
```shell
cd C-NetDisk/CLI
go build -o cnetdisk  # "go build -o cnetdisk.exe" on Windows
```
即可执行：
```shell
cnetdisk --host <C-NetDisk server host> -- port <C-NetDisk server port>
```
开始使用 C-NetDisk 云盘。