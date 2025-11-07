import { useDictStore } from '@/stores/modules/dict'
export interface DictDataType {
  label: string
  value: string | number | boolean
  colorType: 'success' | 'info' | 'warning' | 'danger' | ''
  cssClass: string
}

export interface NumberDictDataType extends DictDataType {
  value: number
}

/**
 * 根据字典类型获取字典选项列表
 * @param dictType - 字典类型标识符
 * @returns 返回指定类型的字典数据列表，如果未找到则返回空数组
 */
export const getDictOptions = (dictType: string) => {
  const dictStore = useDictStore()
  const dictList = (dictStore.getDict(dictType) as DictDataType[]) || []
  return dictList
}
/**
 * 获取指定字典类型的数字类型字典选项列表
 * @param dictType 字典类型标识符
 * @returns 返回转换后的数字类型字典数据数组
 */
export const getIntDictOptions = (dictType: string): NumberDictDataType[] => {
  // 获得通用的 DictDataType 列表
  const dictOptions: DictDataType[] = getDictOptions(dictType)
  // 转换成 number 类型的 NumberDictDataType 类型
  const dictOption: NumberDictDataType[] = []
  dictOptions.forEach((dict: DictDataType) => {
    dictOption.push({
      ...dict,
      value: parseInt(dict.value + ''),
    })
  })
  return dictOption
}
/**
 * 获取字符串类型的字典选项数据
 * @param dictType - 字典类型标识符
 * @returns 返回处理后的字典数据数组，其中每个选项的value值都转换为字符串类型
 */
export const getStrDictOptions = (dictType: string) => {
  const dictOption: DictDataType[] = []
  const dictOptions: DictDataType[] = getDictOptions(dictType)
  dictOptions.forEach((dict: DictDataType) => {
    dictOption.push({
      ...dict,
      value: dict.value + '',
    })
  })
  return dictOption
}
/**
 * 获取布尔类型字典选项
 * @param dictType 字典类型标识符
 * @returns 转换后的字典数据数组，其中value字段被转换为布尔类型
 */
export const getBoolDictOptions = (dictType: string) => {
  const dictOption: DictDataType[] = []
  const dictOptions: DictDataType[] = getDictOptions(dictType)
  dictOptions.forEach((dict: DictDataType) => {
    dictOption.push({
      ...dict,
      value: dict.value + '' === 'true',
    })
  })
  return dictOption
}

/**
 * 获取指定字典类型的指定值对应的字典对象
 * @param dictType 字典类型
 * @param value 字典值
 * @return DictDataType 字典对象
 */
export const getDictObj = (dictType: string, value: any): DictDataType | undefined => {
  const dictOptions: DictDataType[] = getDictOptions(dictType)
  for (const dict of dictOptions) {
    if (dict.value === value + '') {
      return dict
    }
  }
}

/**
 * 获得字典数据的文本展示
 *
 * @param dictType 字典类型
 * @param value 字典数据的值
 * @return 字典名称
 */
export const getDictLabel = (dictType: string, value: any): string => {
  const dictOptions: DictDataType[] = getDictOptions(dictType)
  const dictLabel = ref('')
  dictOptions.forEach((dict: DictDataType) => {
    if (dict.value === value + '') {
      dictLabel.value = dict.label
    }
  })
  return dictLabel.value
}
