import service from '@/utils/request'

export const queryQaStat = (data) => {
    return service({
        url: "/zhihu/queryQaStat",
        method: 'post',
        data
    })
}

export const queryTodoList = (data) => {
    return service({
        url: "/zhihu/queryTodoList",
        method: 'post',
        data
    })
}