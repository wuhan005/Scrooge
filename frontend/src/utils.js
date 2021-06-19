import axios from 'axios'

let baseURL = '/api'

export default {
    baseURL: baseURL,

    GET: (uri, auth = true, bURL = baseURL) => {
        return new Promise((resolve, reject) => {
            axios.get(bURL + uri, {
                headers: {'Authorization': auth ? localStorage.getItem('token') : ''}
            }).then(res => resolve(res.data.data)).catch(err => reject(err.response.data.msg));
        })
    },

    POST: (uri, payload = {}, auth = true, bURL = baseURL) => {
        return new Promise((resolve, reject) => {
            axios.post(bURL + uri, payload, {
                headers: {'Authorization': auth ? localStorage.getItem('token') : ''}
            }).then(res => resolve(res.data.data)).catch(err => reject(err.response.data.msg));
        })
    },

    DELETE: (uri, auth = true, bURL = baseURL) => {
        return new Promise((resolve, reject) => {
            axios.delete(bURL + uri, {
                headers: {'Authorization': auth ? localStorage.getItem('token') : ''}
            }).then(res => resolve(res.data.data)).catch(err => reject(err.response.data.msg));
        })
    },

    PUT: (uri, payload = {}, auth = true, bURL = baseURL) => {
        return new Promise((resolve, reject) => {
            axios.put(bURL + uri, payload, {
                headers: {'Authorization': auth ? localStorage.getItem('token') : ''}
            }).then(res => resolve(res.data.data)).catch(err => reject(err.response.data.msg));
        })
    },
}