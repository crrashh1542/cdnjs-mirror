## CDNJS Mirror

### 技术栈
Gin + React + TypeScript + Vite

### 食用方法
从 Releases 下载对应的二进制包，直接启动即可，程序默认监听 23657 端口。

如果需要服务在特定端口号，需加上 `-p` 参数；如果需要设置项目的域名，需加上 `-s` 参数，且参数内容需包含 https:// 这样的 scheme。

举个栗子，如果我希望服务运行在 9178 端口上，且 Nginx 反代后能通过 https://cdn.example.com 访问，则这样运行：
```
./cdnjs-mirror -h 9178 -s https://cdn.example.com
```

### 构建方法
1. 克隆项目源码
```shell
git clone https://github.com/crrashh1542/cdnjs-mirror --depth=1
```

2. 构建前端项目
```shell
# 已安装 pnpm 的可以跳过这一步
npm install pnpm -g

pnpm install
pnpm run deploy
``` 

3. 回到上级目录，启动服务程序即可，参数与上述一致。
```shell
go run main.go
```
