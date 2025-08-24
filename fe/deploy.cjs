const fs = require('fs').promises
const path = require('path')
const { execSync } = require('child_process')
const { existsSync } = require('fs')

async function deploy() {
    try {
        // 1）构建前端项目
        console.log('[INFO][1/3] 开始构建前端项目')
        execSync('pnpm run build', { stdio: 'inherit' })
        console.log('[INFO][1/3] 前端构建完成！')

        // 2）移动 dist 目录至 ../static
        if (!existsSync('./dist')) {
            throw new Error('[ERROR][2/3] 用于部署的 dist 目录不存在！')
        }

        const targetPath = '../static'
        console.log(`[INFO][2/3] 移动 dist 到 ${targetPath}`)

        // 确保目标父目录存在
        await fs.mkdir(path.dirname(targetPath), { recursive: true })
        // 检查目标位置是否已有同名文件夹
        if (existsSync(targetPath)) {
            console.log('[INFO][3/3] 删除已存在的目标目录...')
            await fs.rm(targetPath, { recursive: true, force: true })
        }

        await fs.rename('./dist', targetPath)
        console.log('[INFO][3/3] 前端部署完成！')

    } catch (err) {
        console.error('[ERROR] 前端部署失败:', err.message)
        process.exit(1)
    }
}

// 执行部署
deploy()