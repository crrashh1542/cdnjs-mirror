const fs = require('fs').promises
const path = require('path')
const { execSync } = require('child_process')
const { existsSync } = require('fs')

const getTime = (time = new Date()) => {
    const pad = num => num.toString().padStart(2, '0') // 用于加0
    const year = time.getFullYear().toString().slice(-2)
    const month = pad(time.getMonth() + 1)
    const day = pad(time.getDate())
    const hours = pad(time.getHours())
    const minutes = pad(time.getMinutes())
    const seconds = pad(time.getSeconds())
    
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

async function deploy() {
    try {
        // 1）构建前端项目
        console.log('[' + getTime() + '][1/8] 开始构建前端项目')
        execSync('pnpm run build', { stdio: 'inherit' })
        console.log('[' + getTime() + '][1/8] 前端构建完成！')

        // 2）移动 dist 目录至 ../static
        if (!existsSync('./dist')) {
            throw new Error('[' + getTime() + '][2/8] 用于部署的 dist 目录不存在！')
        }

        const targetPath = '../static'
        console.log('[' + getTime() + '][2/8] 移动 dist 到 ' + targetPath)

        // 确保目标父目录存在
        await fs.mkdir(path.dirname(targetPath), { recursive: true })
        // 检查目标位置是否已有同名文件夹
        if (existsSync(targetPath)) {
            console.log('[' + getTime() + '][2/8] 删除已存在的目标目录...')
            await fs.rm(targetPath, { recursive: true, force: true })
        }

        await fs.rename('./dist', targetPath)
        console.log('[' + getTime() + '][2/8] 前端部署完成！')

    } catch (err) {
        console.error('[' + getTime() + '] 前端部署失败:', err.message)
        process.exit(1)
    }
}

// 执行部署
deploy()