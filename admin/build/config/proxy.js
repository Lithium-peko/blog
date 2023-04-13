import {getProxyConfig} from '../../setting'

/**
 * http://cn.vitejs.dev/config/server-options.html#server-proxy
 * @param {*} isUserProxy 是否使用代理
 * @param {*} proxyType 当前模式
 */
export function createViteProxy(isUserProxy = true, proxyType) {
  if (!isUserProxy)
    return undefined

  const proxyConfig = getProxyConfig(proxyType)
  return {
    [proxyConfig.prefix]: {
      target: proxyConfig.target,
      changeOrigin: true,
      rewrite: path => path.replace(new RegExp(`^${proxyConfig.prefix}`), ''),
    },
  }
}
