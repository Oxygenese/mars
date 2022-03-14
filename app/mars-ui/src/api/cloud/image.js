import request from '@/utils/request'

export function listImage(query) {
  return request({
    url: '/master/image/list',
    method: 'get',
    params: query
  })
}
