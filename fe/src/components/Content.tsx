import Catalog from './widgets/Catalog'
import CodeBlock from "./widgets/CodeBlock"

import config from '../../config'
import '../styles/content.less'

const codeDefault = '<script src="https://cdnjs.cloudflare.com/ajax/libs/react/19.1.1/cjs/react.production.min.js"></script>'
const codeCustom = '<script src="' + config.site + '/react/19.1.1/cjs/react.production.min.js"></script>'
const codeCustomIntegrity = 'integrity="sha512-{对应的 SHA512 哈希}"'
const codeCustomIntegrityExample = '<link rel="stylesheet" href="' + config.site + '/normalize/8.0.1/normalize.min.css" integrity="sha512-NhSC1YmyruXifcj/KFRWoC561YpHpc5Jtzgvbuzx5VozKpWvQ+4nXhPdFgmx8xqexRcpAglTj9sIBWINXa8x5w=="/>'

export default () =>
   <main>

      <Catalog>默认配置</Catalog>
      <p>将你原来引用 CDNJS 静态资源的地址：</p>
      <CodeBlock code={codeDefault} />
      <p>改为本站点域名即可！</p>
      <CodeBlock code={codeCustom} />

      <br /><br />
      
      <Catalog>校验内容</Catalog>
      <p>如果担心被劫持，可以在标签中加入 <code>{'integrity'}</code> 属性用于验证。</p>
      <CodeBlock code={codeCustomIntegrity} />
      <p>举个栗子，当引入 CSS 文件时，食用方法如下：</p>
      <CodeBlock code={codeCustomIntegrityExample}/>

   </main>