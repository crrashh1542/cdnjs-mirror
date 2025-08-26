// 此 util 用于通过后端的 getStatus 接口获取站点信息

interface StatusResponse {
    code: number
    site: string
    build: string
    version: string
}

const fetchData = async (): Promise<StatusResponse> => {
    const response = await fetch('/getStatus')
    const data = await response.json()
    return data
}

export default fetchData()