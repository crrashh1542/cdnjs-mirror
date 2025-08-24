// 此 util 用于通过后端的 getData 接口获取设定的站点

const fetchData = async (): Promise<string> => {
    const response = await fetch('/getSite')
    const data = await response.json()
    return data.site
}

export default fetchData()