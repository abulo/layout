import * as IconSpace from '@icon-space/vue-next/es'
import * as EpIcons from '@element-plus/icons-vue'
import { toKebabCase } from '@/utils'

export const IconJson = () => {
  const IconData = {}
  const iconSpaceList = IconSpace as any
  const ip: string[] = []
  for (const i in iconSpaceList) {
    ip.push(toKebabCase(i))
  }
  const ep: string[] = []
  const icons = EpIcons as any
  for (const i in icons) {
    ep.push(toKebabCase(icons[i].name))
  }
  IconData['ep:'] = ep
  IconData['ip:'] = ip
  return IconData
}
