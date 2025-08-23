const fs = require('fs').promises
const path = require('path')
const process = require('child_process')
const { existsSync } = require('fs')
const { error } = require('console')

async function deployFolder(src, dest) {
  // 01 - 构建前端项目
  process.execSync('pnpm run build', (error, stdout, stderr) => {})

  // 02 - 移动 dist 目录至上级 static
  try {
    // 确保源目录存在
    if (!existsSync(src)) {
      throw new Error(`源目录不存在: ${src}`)
    }

    // 获取源目录的父目录和文件夹名
    const srcParent = path.dirname(src)
    const srcName = path.basename(src)

    // 解析目标路径（可以包含新名字）
    const destParent = path.dirname(dest)
    const destName = path.basename(dest)
    const finalDest = path.join(destParent, destName)

    // 确保目标父目录存在
    await fs.mkdir(destParent, { recursive: true })

    // 检查目标位置是否已有同名文件夹
    if (existsSync(finalDest)) {
      throw new Error(`目标目录已存在: ${finalDest}`)
    }
    // 使用 rename 进行移动（原子操作，高效）
    await fs.rename(src, finalDest);
  } catch (err) {
    throw err
  }
}

// 使用示例
(async () => {
  const sourceFolder = './dist';
  const targetPath = '../static';

  await deployFolder(sourceFolder, targetPath);
})()