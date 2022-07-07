import request from './request'

export const getTags = (params) => {
  return request({
    url: 'tags',
    params
  })
}

export const addTag = (data) => {
  return request({
    url: 'tags',
    method: 'post',
    data
  })
}

export const editTag = (data) => {
  return request({
    url: `tags/${data.id}`,
    method: 'put',
    data
  })
}

export const deleteTag = (id) => {
  return request({
    url: `tags/${id}`,
    method: 'delete'
  })
}

export const tagNameExists = (params) => {
  return request({
    url: 'tags/tagNameExists',
    method: 'get',
    params
  })
}
