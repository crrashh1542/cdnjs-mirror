#!/bin/bash
set -e

echo "[$(date "+%y-%m-%d %H:%M:%S")] CDNJS Mirror 构建开始"

VERSION="1.1.2"
BUILD_HASH=$(git rev-parse --short HEAD)
PLATFORMS=("linux" "darwin" "windows")
ARCHITECTURES=("amd64" "arm64")
BUILD_COUNT=1

echo "[$(date "+%y-%m-%d %H:%M:%S")] 当前版本：v${VERSION}"

# 判断是否有未提交的更改，如果有就在 hash 后加上 -dirty
if git diff-index --quiet HEAD -- && git diff-files --quiet; then
    echo "[$(date "+%y-%m-%d %H:%M:%S")] 当前提交：${BUILD_HASH}"
else
    echo "[$(date "+%y-%m-%d %H:%M:%S")] 当前提交：${BUILD_HASH}，存在未提交的更改"
    BUILD_HASH="${BUILD_HASH}-dirty"
fi
LDFLAGS="-w -s -X main.buildId=$BUILD_HASH"

# 构建前端
cd fe
pnpm install
pnpm run deploy
cd ..
BUILD_COUNT=$((BUILD_COUNT + 1))

# 构建后端
for platform in "${PLATFORMS[@]}"; do
    for arch in "${ARCHITECTURES[@]}"; do
        BUILD_COUNT=$((BUILD_COUNT + 1))
        echo "[$(date "+%y-%m-%d %H:%M:%S")][${BUILD_COUNT}/8] 当前构建平台：${platform}；架构：${arch}"
        if [ "${platform}" == "windows" ]; then
            OUTPUT_NAME="dist/cdnjs-mirror-v${VERSION}-${platform}-${arch}.exe"
        else
            OUTPUT_NAME="dist/cdnjs-mirror-v${VERSION}-${platform}-${arch}"
        fi
        
        CGO_ENABLED=0 GOOS="${platform}" GOARCH="${arch}" go build -ldflags "${LDFLAGS}" -o "${OUTPUT_NAME}" .
        
        if [ $? -eq 0 ]; then
            echo "[$(date "+%y-%m-%d %H:%M:%S")][${BUILD_COUNT}/8] ${platform}-${arch} 构建成功！"
        else
            echo "[$(date "+%y-%m-%d %H:%M:%S")][${BUILD_COUNT}/8] ${platform}-${arch} 构建失败！"
            exit 1
        fi
    done
done

echo "构建结束！"