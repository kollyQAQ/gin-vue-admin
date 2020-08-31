import service from '@/utils/request'

export const getQaList = (data) => {
    return service({
        url: "/zhihu/getQaList",
        method: 'post',
        data
    })
}

export const getQaById = (data) => {
    return service({
        url: "/zhihu/getQaById",
        method: 'post',
        data
    })
}

export const createQa = (data) => {
    return service({
        url: "/zhihu/createQa",
        method: 'post',
        data
    })
}

export const updateQa = (data) => {
    return service({
        url: "/zhihu/updateQa",
        method: 'post',
        data
    })
}

export const deleteQa = (data) => {
    return service({
        url: "/zhihu/deleteQa",
        method: 'post',
        data
    })
}

export const queryQaHistory = (data) => {
    return service({
        url: "/zhihu/queryQaHistory",
        method: 'post',
        data
    })
}