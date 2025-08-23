## CDNJS Mirror

### 技术栈
Gin + React + TypeScript + Less + Vite

### 构建及食用方法
1. 克隆项目源码
```shell
git clone https://github.com/crrashh1542/light-ghchart-index --depth=1
```
2. 进入 `fe` 目录，按照注释修改 [`config.ts`](./fe/config.ts) 中相关配置。

3. 构建前端项目
```shell
# 已安装 pnpm 的可以跳过这一步
npm install pnpm -g

pnpm install
pnpm run deploy
``` 

4. 回到上级目录，启动服务程序即可。程序默认运行在 23646 端口。
```shell
go run main.go
```
