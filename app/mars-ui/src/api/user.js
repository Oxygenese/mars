import request from '@/utils/request'

// login 登陆
export function login(data) {
  return request({
    url: '/auth/token',
    method: 'post',
    params: data
  })
}

// logout 退出
export function logout() {
  return request({
    url: '/auth/logout',
    method: 'put'
  })
}

// refreshtoken 刷新token
export function refreshtoken(data) {
  return request({
    url: '/refreshtoken',
    method: 'post',
    data
  })
}

// getInfo 获取用户基本信息
export function getInfo() {
  return request({
    url: '/sys/user/info',
    method: 'get'
  })
}

