const dayjs = require('dayjs')

const filterTimes = (val, format = 'YYYY-MM-DD') => {
  if (!isNull(val)) {
    val = parseInt(val) * 1000
    return dayjs(val).format(format)
  } else {
    return '--'
  }
}

const filterNull = (val) => {
  if (!isNull(val)) {
    return val
  } else {
    return '--'
  }
}

export const TrimString = (val, len) => {
  if (val.length > len) {
    return val.substring(0, len) + '...'
  }
  return val
}

export const isNull = (date) => {
  if (!date) return true
  if (JSON.stringify(date) === '{}') return true
  if (JSON.stringify(date) === '[]') return true
  if (JSON.stringify(date) === 'undefined') return true
  return false
}

export default (app) => {
  app.config.globalProperties.$filters = {
    filterTimes,
    filterNull
  }
}
