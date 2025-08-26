import { useState, useEffect } from 'react'

import '../styles/footer.less'
import { version } from '../../package.json'
import statusPromise from '../utils/getStatus'
import externalIcon from '../assets/images/external.svg'

interface ResultType {
    build: string
    version: string
}

export default () => {
   const [beVersion, setBeVersion] = useState<string[]>(['', ''])
   useEffect(() => {
      statusPromise.then((result: ResultType) => {
         setBeVersion([result.version, result.build])
      }).catch(error => {
         console.error('Error resolving Promise object! ', error)
      })
   }, [])

   return (
      <footer>
         <p>服务端: v{beVersion[0]} ({beVersion[1]}) / 前端: v{version}</p>
         <p>&copy; 2025 crrashh1542. 本项目由 <a href="https://cdnjs.com">
            CDNJS<img src={externalIcon} alt="iconfont-external" /></a> 驱动</p>
      </footer>
   )
}