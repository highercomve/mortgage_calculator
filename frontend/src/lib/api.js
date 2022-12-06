import fetch from 'isomorphic-fetch'

export function ResolvePath(obj, path, defaultReturn) {
  return path.split(".").reduce(function (prev, curr) {
    return prev && prev.hasOwnProperty(curr) ? prev[curr] : defaultReturn;
  }, obj);
}

export const GetEnv = (env, def) => {
  return ResolvePath(window, `env.${env}`, def)
}

export const GetApiURL = () =>
  GetEnv('REACT_APP_API_URL','http://localhost:9090/api/v1')

export const _requestContentType = async (
  url,
  contentType = 'application/json',
  method = 'GET',
  body = {}
) => {
  const headers = {
    'Content-Type': contentType
  }

  let options = {
    method,
    headers: headers,
    credentials: "include"
  }

  if (method !== 'GET' && method !== 'HEAD') {
    options['body'] = JSON.stringify(body)
  }

  return fetch(url, options)
}

const _requestJSON = async (url, method = 'GET', body = {}) => {
  const response = await _requestContentType(
    url,
    'application/json',
    method,
    body
  )

  const json = await response.json()
  
  return {
    ok: response.ok,
    redirected: response.redirected,
    headers: response.headers,
    status: response.status,
    json
  }
}

export const Get = async (url) =>
  _requestJSON(url, 'GET')

export const Post = async (url, body) =>
  _requestJSON(url, 'POST', body)

export const Put = async (url, body) =>
  _requestJSON(url, 'PUT', body)

export const Patch = async (url, body) =>
  _requestJSON(url, 'PATCH', body)

export const Delete = async (url) =>
  _requestJSON(url, 'DELETE')
