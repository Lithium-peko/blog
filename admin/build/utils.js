import path from 'path'

/**
 * 项目根路径
 * @description 结尾不带 '/'
 */
export function getRootPath() {
  return path.resolve(process.cwd()) // cwd - current work directory
}

/**
 * 项目src路径
 * @param srcName src目录名称(默认: "src")
 * @description 结尾不带 '/'
 */
export function getSrcPath(srcName = 'src') {
  return path.resolve(getRootPath(), srcName)
}

export function convertEnv(envObj) {
  const result = {}
  if (!envObj)
    return result
  for (const envKey in envObj) {
    let envVal = envObj[envKey]
    if (['true', 'false'].includes(envVal))
      envVal = envVal === 'true'
    if (['VITE_PORT'].includes(envKey))
      envVal = +envVal
    result[envKey] = envVal
  }
  return result
}
