import {
  GetApiURL,
  Post,
} from '../lib/api'

const CALCULATOR_URL = `${GetApiURL()}/calculate`

export const PostCalculate = async (body = {}) => {
  return Post(`${CALCULATOR_URL}`, body)
}
