import { useState, useEffect } from 'react'
import { Light as SyntaxHighlighter } from 'react-syntax-highlighter'
import javascript from 'react-syntax-highlighter/dist/esm/languages/hljs/javascript'
import { atomOneLight } from 'react-syntax-highlighter/dist/esm/styles/hljs'

import Catalog from './widgets/Catalog'

import sitePromise from '../utils/getSite'
import '../styles/content.less'

// Light 引入的高亮语法需要手动注册
SyntaxHighlighter.registerLanguage('javascript', javascript)

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
            <SyntaxHighlighter language="xml" style={ atomOneLight } className="code-block">{'<script src="https://cdnjs.cloudflare.com/ajax/libs/react/19.1.1/cjs/react.production.min.js"></script>'}</SyntaxHighlighter>
            <p>改为本站点域名即可！</p>
            <SyntaxHighlighter language="xml" style={ atomOneLight } className="code-block">{'<script src="' + site + '/react/19.1.1/cjs/react.production.min.js"></script>'}</SyntaxHighlighter>
            <br /><br />

            <Catalog>校验内容</Catalog>
            <p>如果担心被劫持，可以在标签中加入 <code>{'integrity'}</code> 属性用于验证。</p>
            <SyntaxHighlighter language="xml" style={ atomOneLight } className="code-block">{'integrity="sha512-{对应的 SHA512 哈希}"'}</SyntaxHighlighter>
            <p>举个栗子，当引入 CSS 文件时，食用方法如下：</p>
            <SyntaxHighlighter language="xml" style={ atomOneLight } className="code-block">{'<link rel="stylesheet" href="' + site + '/normalize/8.0.1/normalize.min.css" integrity="sha512-NhSC1YmyruXifcj/KFRWoC561YpHpc5Jtzgvbuzx5VozKpWvQ+4nXhPdFgmx8xqexRcpAglTj9sIBWINXa8x5w=="/>'}</SyntaxHighlighter>

        </main>
    )
}