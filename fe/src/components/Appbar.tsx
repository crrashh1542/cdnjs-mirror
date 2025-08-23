import githubImg from '../assets/images/github.svg'
import '../styles/appbar.less'

export default () =>
   <header>
      <a className="title" href="#">CDNJS Mirror</a>
      <div className="grow"></div>
      <a href="https://github.com/crrashh1542/cdnjs-mirror-index">
         <img src={ githubImg } alt="本引导页的 GitHub 项目地址" />
      </a>
   </header>