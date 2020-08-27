import service from '@/utils/request'

export const getCollectionList = (data) => {
    return service({
        url: "/zhihu/getCollectionList",
        method: 'get',
        data
    })
}