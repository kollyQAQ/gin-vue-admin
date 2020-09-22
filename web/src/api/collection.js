import service from '@/utils/request'

export const getCollectionList = (data) => {
    return service({
        url: "/zhihu/getCollectionList",
        method: 'post',
        data
    })
}

export const getCollectionListV2 = (data) => {
    return service({
        url: "/zhihu/getCollectionListV2",
        method: 'post',
        data
    })
}

export const getCollectionById = (data) => {
    return service({
        url: "/zhihu/getCollectionById",
        method: 'post',
        data
    })
}

export const updateCollection = (data) => {
    return service({
        url: "/zhihu/updateCollection",
        method: 'post',
        data
    })
}

export const createCollection = (data) => {
    return service({
        url: "/zhihu/createCollection",
        method: 'post',
        data
    })
}

export const deleteCollection = (data) => {
    return service({
        url: "/zhihu/deleteCollection",
        method: 'post',
        data
    })
}
