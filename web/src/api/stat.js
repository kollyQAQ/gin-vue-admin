import service from '@/utils/request'

export const queryStat = (data) => {
    return service({
        url: "/zhihu/queryStat",
        method: 'post',
        data
    })
}

export const queryMusicList = (data) => {
    return service({
        url: "/zhihu/queryMusicList",
        method: 'post',
        data
    })
}

export const queryPlanList = (data) => {
    return service({
        url: "/zhihu/queryPlanList",
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

export const addTodo = (data) => {
    return service({
        url: "/zhihu/addTodo",
        method: 'post',
        data
    })
}

export const updateTodo = (data) => {
    return service({
        url: "/zhihu/updateTodo",
        method: 'post',
        data
    })
}

export const deleteTodo = (data) => {
    return service({
        url: "/zhihu/deleteTodo",
        method: 'post',
        data
    })
}

export const queryTypeList = (data) => {
    return service({
        url: "/zhihu/queryTypeList",
        method: 'post',
        data
    })
}