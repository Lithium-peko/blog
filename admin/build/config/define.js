import dayjs from "dayjs";

/**
 *  * 此处定义的是全局变量, 启动或打包后将添加到 window 中
 */

const _BUILD_TIME_ = JSON.stringify(dayjs().format('YYYY-MM-DD HH:mm:ss'))

export const viteDefine = {
  _BUILD_TIME_
}
