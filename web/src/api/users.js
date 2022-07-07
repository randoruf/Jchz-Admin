import request from './request'

export const getUsers = (params) => {
  return request({
    url: 'users',
    params
  })
}

export const changeUserState = (uid, type) => {
  return request({
    url: `users/${uid}/state/:${type}`,
    method: 'put'
  })
}

export const addUser = (data) => {
  return request({
    url: 'users',
    method: 'post',
    data
  })
}

export const editUser = (data) => {
  return request({
    url: `users/${data.id}`,
    method: 'put',
    data
  })
}

export const deleteUser = (id) => {
  return request({
    url: `users/${id}`,
    method: 'delete'
  })
}

export const addComUser = (data) => {
  return request({
    url: 'companyUsers',
    method: 'post',
    data
  })
}

export const editComUser = (data) => {
  return request({
    url: `companyUsers/${data.id}`,
    method: 'put',
    data
  })
}

export const deleteComUser = (id) => {
  return request({
    url: `companyUsers/${id}`,
    method: 'delete'
  })
}

export const addAdmin = (data) => {
  return request({
    url: 'admin',
    method: 'post',
    data
  })
}

export const editAdmin = (data) => {
  return request({
    url: `admin/${data.id}`,
    method: 'put',
    data
  })
}

export const deleteAdmin = (id) => {
  return request({
    url: `admin/${id}`,
    method: 'delete'
  })
}

export const editMyself = (data) => {
  return request({
    url: `admin/editMyself/${data.id}`,
    method: 'put',
    data
  })
}
