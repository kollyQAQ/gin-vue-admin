import service from '@/utils/request'

export const getHighRateGoodsList = (data) => {
    return service({
        url: "/zhihu/getHighRateGoodsList",
        method: 'post',
        data
    })
}

export const getFirstCategory = (data) => {
    return service({
        url: "/zhihu/getGoodsCategory",
        method: 'post',
        data
    })
}

export const getSecondCategory = (data) => {
    return service({
        url: "/zhihu/getGoodsCategory",
        method: 'post',
        data
    })
}

export const getThirdCategory = (data) => {
    return service({
        url: "/zhihu/getGoodsCategory",
        method: 'post',
        data
    })
}

export const getGoodsList = (data) => {
    return service({
        url: "/zhihu/getGoodsList",
        method: 'post',
        data
    })
}

export const getGoodsCategory = (data) => {
    return service({
        url: "/zhihu/getGoodsCategory",
        method: 'post',
        data
    })
}

export const getGoodsById = (data) => {
    return service({
        url: "/zhihu/getGoodsById",
        method: 'post',
        data
    })
}

export const createGoods = (data) => {
    return service({
        url: "/zhihu/createGoods",
        method: 'post',
        data
    })
}

export const updateGoods = (data) => {
    return service({
        url: "/zhihu/updateGoods",
        method: 'post',
        data
    })
}

export const deleteGoods = (data) => {
    return service({
        url: "/zhihu/deleteGoods",
        method: 'post',
        data
    })
}
