import request from './request'

export const getArticles = (params) => {
  return request({
    url: 'article',
    params
  })
}

export const editArticle = (data) => {
  return request({
    url: `article/${data.id}`,
    method: 'put',
    data
  })
}

export const deleteArticle = (id) => {
  return request({
    url: `article/${id}`,
    method: 'delete'
  })
}
