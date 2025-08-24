import { useState, useEffect } from 'react'

import Catalog from './widgets/Catalog'
import CodeBlock from "./widgets/CodeBlock"

import sitePromise from '../utils/getSite'
import '../styles/content.less'

export default () => {
    // 将 site 的 Promise 对象解析并保存到 site 变量里
    const [site, setSite] = useState<string>('')
    useEffect(() => {
        sitePromise.then(result => {
            setSite(result)
        }).catch(error => {
            console.error('Error resolving Promise object! ', error)
        })
    }, [])

    return ( 
        <main>
            <Catalog>默认配置</Catalog>
            <p>将你原来引用 CDNJS 静态资源的地址：</p>
            <CodeBlock code={'<script src="https://cdnjs.cloudflare.com/ajax/libs/react/19.1.1/cjs/react.production.min.js"></script>'} />
            <p>改为本站点域名即可！</p>
            <CodeBlock code={'<script src="' + site + '/react/19.1.1/cjs/react.production.min.js"></script>'} />
            <br /><br />

            <Catalog>校验内容</Catalog>
            <p>如果担心被劫持，可以在标签中加入 <code>{'integrity'}</code> 属性用于验证。</p>
            <CodeBlock code={'integrity="sha512-{对应的 SHA512 哈希}"'} />
            <p>举个栗子，当引入 CSS 文件时，食用方法如下：</p>
            <CodeBlock code={'<link rel="stylesheet" href="' + site + '/normalize/8.0.1/normalize.min.css" integrity="sha512-NhSC1YmyruXifcj/KFRWoC561YpHpc5Jtzgvbuzx5VozKpWvQ+4nXhPdFgmx8xqexRcpAglTj9sIBWINXa8x5w=="/>'} />

        </main>
    )
}