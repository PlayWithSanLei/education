import axios from "axios";

let config = {
    baseURL: 'http://121.196.144.74:8100/api/v1',
    timeout: 8000
}


const _axios = axios.create(config)

_axios.interceptors.request.use(config => {
    // console.log('异步交互之前触发拦截')
    if (localStorage.getItem('token')) {
        config.headers.Authorization = window.localStorage.getItem('token')
    }
    return config
}, e => {

})

export default _axios
