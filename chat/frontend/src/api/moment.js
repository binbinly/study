import api from './index'

export function momentCreate(data) {
  return api.post(api.Moment.Create, data)
}

export function momentLike(data) {
  return api.post(api.Moment.Like, data)
}

export function momentComment(data) {
  return api.post(api.Moment.Comment, data)
}

export function momentList(key, user_id, p) {
  if (key == 'timeline') {
    return api.get(api.Moment.Timeline, { p })
  }
  return api.get(api.Moment.List, { user_id, p })
}
