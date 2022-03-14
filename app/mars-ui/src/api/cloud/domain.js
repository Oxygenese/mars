import request from '@/utils/request'

export function listDomain(query) {
  return request({
    url: '/master/domain/page',
    method: 'get',
    params: query
  })
}

export function createDomain(data) {
  return request({
    url: '/master/domain/create',
    method: 'post',
    data: data
  })
}

export function deleteDomain(query) {
  return request({
    url: '/master/domain',
    method: 'delete',
    params: query
  })
}

export function stopDomain(domainId) {
  return request({
    url: '/master/domain/delete',
    method: 'delete'
  })
}

