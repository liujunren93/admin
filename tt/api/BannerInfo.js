import * as http from '@/utils/http'
const api='bannerInfo'
export function list (parameter) {
   return http.get(api, parameter)
}
export function create (parameter) {
   return http.post(api, parameter)
}
export function info (id,parameter) {
   return http.get(api+'/'+id, parameter)
}
export function update (id,parameter) {
   return http.put(api+'/'+id, parameter)
}
export function del (id,parameter) {
   return http.delete(api+'/'+id, parameter)
}
