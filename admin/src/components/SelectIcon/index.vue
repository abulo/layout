<template>
  <el-popover
    :placement="placement"
    trigger="focus"
    :hide-after="0"
    :width="state.selectorWidth"
    :visible="state.popoverVisible"
    :show-arrow="false"
  >
    <div
      class="icon-selector"
      @mouseover.stop="state.iconSelectorMouseover = true"
      @mouseout.stop="state.iconSelectorMouseover = false"
    >
      <transition name="el-zoom-in-center">
        <div v-if="show" class="icon-selector-box">
          <div class="selector-header">
            <!-- <div class="selector-title">{{ title ? title : "" }}</div> -->
            <el-tabs v-model="currentActiveType" type="border-card" @tab-click="handleTab">
              <el-tab-pane v-for="(tab, key) in tabsList" :key="key" :label="tab.label" :name="tab.name">
                <div class="selector-body">
                  <el-scrollbar ref="selectorScrollbarRef" height="20vh">
                    <div v-if="pageList.length > 0">
                      <div
                        v-for="(item, index) in pageList"
                        :key="index"
                        class="icon-selector-item"
                        :title="String(item)"
                        @click="onIcon(String(item))"
                      >
                        <el-button>
                          <Icon :icon="currentActiveType + item" />
                        </el-button>
                      </div>
                    </div>
                  </el-scrollbar>
                </div>
              </el-tab-pane>
            </el-tabs>
          </div>
        </div>
      </transition>
      <el-pagination
        :current-page="currentPage"
        :page-size="pageSize"
        :total="iconCount"
        background
        size="small"
        layout="prev, pager, next"
        @current-change="onCurrentChange"
      />
    </div>
    <template #reference>
      <el-input
        ref="selectorInput"
        v-model="state.inputValue"
        :size="size"
        :disabled="disabled"
        :placeholder="placeholder"
        :class="'size-' + size"
        @focus="onInputFocus"
        @blur="onInputBlur"
      >
        <template #prepend>
          <div class="icon-prepend">
            <el-button :key="'icon' + state.iconKey">
              <Icon :icon="state.prependIcon ? state.prependIcon : state.defaultModelValue" />
            </el-button>
            <div v-if="showIconName" class="name">
              {{ state.prependIcon ? state.prependIcon : state.defaultModelValue }}
            </div>
          </div>
        </template>
        <template #append>
          <el-button :icon="Icons.Refresh" @click="onInputRefresh" />
        </template>
      </el-input>
    </template>
  </el-popover>
</template>

<script setup lang="ts">
defineOptions({ name: 'SelectIcon' })
import { reactive, ref, onMounted, onUnmounted, nextTick, watch, computed } from 'vue'
import { useEventListener } from '@vueuse/core'
import { Placement } from 'element-plus'
import { IconJson } from '@/components/Icon/data'
import { cloneDeep } from 'lodash-es'
import Icon from '@/components/Icon/index.vue'
import * as Icons from '@element-plus/icons-vue'

const tabsList = [
  {
    label: 'Element Plus',
    name: 'ep:',
  },
  {
    label: 'Icon Space',
    name: 'ip:',
  },
]

const iconList = ref(IconJson())

const pageSize = ref(100)
const currentPage = ref(1)

const currentActiveType = ref('ep:')

// 深拷贝图标数据，前端做搜索
const copyIconList = cloneDeep(iconList.value)

interface SelectIconProps {
  size?: 'default' | 'small' | 'large'
  disabled?: boolean
  title?: string
  placement?: Placement
  modelValue?: string
  showIconName?: boolean
  placeholder?: string
  show?: boolean
}

const props = withDefaults(defineProps<SelectIconProps>(), {
  size: 'default',
  disabled: false,
  title: '',
  placement: 'bottom',
  modelValue: '',
  showIconName: false,
  placeholder: '',
  show: true,
})

// const emits = defineEmits<{
//   (_e: 'update:modelValue', _value: string): void
//   (_e: 'change', _value: string): void
// }>()

const emits = defineEmits<{
  'update:modelValue': [value: string]
  change: [value: string]
}>()

const handleTab = ({ props }) => {
  currentActiveType.value = props.name
}

const selectorInput = ref()
const selectorScrollbarRef = ref()
const state: {
  selectorWidth: number
  popoverVisible: boolean
  inputFocus: boolean
  iconSelectorMouseover: boolean
  // fontIconNames: string[];
  inputValue: string
  prependIcon: string
  defaultModelValue: string
  iconKey: number
} = reactive({
  selectorWidth: 0,
  popoverVisible: false,
  inputFocus: false,
  iconSelectorMouseover: false,
  // fontIconNames: [],
  inputValue: '',
  prependIcon: props.modelValue,
  defaultModelValue: props.modelValue || 'ep:search',
  iconKey: 0, // 给icon标签准备个key，以随时使用 h 函数重新生成元素
})

const pageList = computed(() => {
  let list = Array.from(new Set(copyIconList[currentActiveType.value])) as string[]
  if (currentPage.value === 1) {
    return list?.filter(v => v.includes(state.inputValue)).slice(currentPage.value - 1, pageSize.value)
  } else {
    return list
      ?.filter(v => v.includes(state.inputValue))
      .slice(pageSize.value * (currentPage.value - 1), pageSize.value * (currentPage.value - 1) + pageSize.value)
  }
})
const iconCount = computed(() => {
  // 去除重复值的数组
  let list = Array.from(new Set(copyIconList[currentActiveType.value])) as string[]
  list = list?.filter(v => v.includes(state.inputValue)) as string[]
  let total = list == undefined ? 0 : list.length
  return total
})

const onInputFocus = () => {
  state.inputFocus = state.popoverVisible = true
}

const onInputBlur = () => {
  state.inputFocus = false
  state.popoverVisible = state.iconSelectorMouseover
}

const onInputRefresh = () => {
  state.iconKey++
  state.prependIcon = state.defaultModelValue
  state.inputValue = ''
  emits('update:modelValue', state.defaultModelValue)
  emits('change', state.defaultModelValue)
}

// const renderFontIconNames = computed(() => {
//   if (!state.inputValue) return state.fontIconNames;
//   let inputValue = state.inputValue.trim().toLowerCase();
//   return state.fontIconNames.filter((icon: string) => {
//     if (icon.toLowerCase().indexOf(inputValue) !== -1) {
//       return icon;
//     }
//   });
// });

// const customIcons: { [key: string]: any } = Icons

// 获取 input 的宽度
const getInputWidth = () => {
  nextTick(() => {
    state.selectorWidth = selectorInput.value.$el.offsetWidth < 260 ? 260 : selectorInput.value.$el.offsetWidth
  })
}

const popoverVisible = () => {
  state.popoverVisible = state.inputFocus || state.iconSelectorMouseover ? true : false
}

// const IconNames = () => {
//   return new Promise<string[]>((resolve, reject) => {
//     nextTick(() => {
//       // console.log(currentActiveType.value);
//       // 定义 list 为一个空数组
//       const list: string[] = [];
//       // const icons = Icons as any;
//       // for (const i in icons) {
//       // list.push(`${icons[i].name}`);
//       // }
//       // list = iconList.value[currentActiveType.value];
//       list.push(...copyIconList[currentActiveType.value]);
//       if (list.length > 0) {
//         resolve(list);
//       } else {
//         reject("No ElementPlus Icons");
//       }
//     });
//   });
// };

const onIcon = (icon: string) => {
  state.iconSelectorMouseover = state.popoverVisible = false
  state.iconKey++
  state.prependIcon = currentActiveType.value + icon
  state.inputValue = ''
  emits('update:modelValue', currentActiveType.value + icon)
  emits('change', currentActiveType.value + icon)
  nextTick(() => {
    selectorInput.value.blur()
  })
}
watch(
  () => props.modelValue,
  () => {
    state.iconKey++
    if (props.modelValue != state.prependIcon) state.defaultModelValue = props.modelValue
    if (props.modelValue == '') state.defaultModelValue = ''
    state.prependIcon = props.modelValue
  }
)

// 监听 currentActiveType
watch(
  () => currentActiveType.value,
  () => {
    if (currentActiveType.value != '') {
      // IconNames().then(res => {
      //   state.fontIconNames = res;
      // });
      currentPage.value = 1
    }
  }
)

watch(
  () => {
    return state.inputValue
  },
  () => {
    currentPage.value = 1
  }
)

function onCurrentChange(page) {
  currentPage.value = page
}

onMounted(() => {
  getInputWidth()
  useEventListener(document, 'click', popoverVisible)
  // IconNames().then(res => {
  //   state.fontIconNames = res;
  // });
  window.onresize = () => {
    getInputWidth()
  }
})

onUnmounted(() => {
  window.onresize = null
})
</script>

<style scoped lang="scss">
@use './index';
</style>
