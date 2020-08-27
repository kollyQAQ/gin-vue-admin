import service from '@/utils/request'

export const getQaList = (data) => {
    return service({
        url: "/zhihu/getQaList",
        method: 'get',
        data
    })
}