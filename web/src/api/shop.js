import request from './request'

export const getShops = (params) => {
  return request({
    url: 'shop',
    params
  })
}

export const addShop = (data) => {
  return request({
    url: 'shop',
    method: 'post',
    data
  })
}

export const editShop = (data) => {
  return request({
    url: `shop/${data.id}`,
    method: 'put',
    data
  })
}

export const deleteShop = (id) => {
  return request({
    url: `shop/${id}`,
    method: 'delete'
  })
}

export const CheckComID = (id) => {
  return request({
    url: `checkComID/${id}`,
    method: 'get'
  })
}
