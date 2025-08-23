import '../styles/footer.less'
import config from '../../config'
import externalIcon from '../assets/images/external.svg'
// import miitIcon from '../assets/images/miit.png'

export default () => 
   <footer>
      <p>本项目由 <a href="https://cdnjs.com">
            CDNJS<img src={externalIcon} alt="iconfont-external" /></a> 驱动，
            服务由 <a href={ 'https://github.com/' + config.username }>
         { config.nickname }<img src={externalIcon} alt="iconfont-external" /></a> 提供。</p>
         <p>&copy; 2025 crrashh1542</p>
   </footer>