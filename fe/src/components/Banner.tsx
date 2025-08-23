import '../styles/banner.less'
import externalIcon from '../assets/images/external.svg'

export default () =>
   <div className="banner">
      <div className="title">CDNJS Mirror</div>
      <div className="subtitle">
         一个 <a href="https://cdnjs.com">CDNJS
            <img src={externalIcon} alt="iconfont-external" />
         </a> 的镜像站点
      </div>
   </div>