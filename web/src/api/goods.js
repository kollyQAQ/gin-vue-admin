import service from '@/utils/request'

export const getGoodsList = (data) => {
    return service({
        url: "/zhihu/getGoodsList",
        method: 'get',
        data
    })
}

export const getGoodsById = (data) => {
    return service({
        url: "/zhihu/getGoodsById",
        method: 'get',
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