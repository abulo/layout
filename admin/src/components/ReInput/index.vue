<template>
  <div class="custom-range-input">
    <div class="range-inputs-container">
      <!-- 开始输入框 -->
      <el-input
        ref="startInputRef"
        v-model="startValue"
        :placeholder="startPlaceholder"
        :clearable="clearable"
        :disabled="disabled"
        :readonly="readonly"
        :maxlength="maxlength"
        class="start-input"
        @input="handleStartInput"
        @change="handleStartChange"
        @focus="handleStartFocus"
        @blur="handleStartBlur"
        @clear="handleStartClear"
        @keyup.enter="handleEnter"
        @keyup="handleKeyup"
      >
      </el-input>
      <!-- 分隔符 -->
      <span class="separator" :class="separatorClass">{{ rangeSeparator }}</span>

      <!-- 结束输入框 -->
      <el-input
        ref="endInputRef"
        v-model="endValue"
        :placeholder="endPlaceholder"
        :clearable="clearable"
        :disabled="disabled"
        :readonly="readonly"
        :maxlength="maxlength"
        class="end-input"
        @input="handleEndInput"
        @change="handleEndChange"
        @focus="handleEndFocus"
        @blur="handleEndBlur"
        @clear="handleEndClear"
        @keyup.enter="handleEnter"
        @keyup="handleKeyup"
      >
      </el-input>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, computed, nextTick } from 'vue'
defineOptions({ name: 'ReInput' })
const props = defineProps({
  // 双向绑定的值，格式为 [startValue, endValue]
  modelValue: {
    type: Array,
    default: () => [null, null],
  },
  // 分隔符
  rangeSeparator: {
    type: String,
    default: '至',
  },
  // 开始输入框占位符
  startPlaceholder: {
    type: String,
    default: '请输入开始值',
  },
  // 结束输入框占位符
  endPlaceholder: {
    type: String,
    default: '请输入结束值',
  },
  // 是否可清空
  clearable: {
    type: Boolean,
    default: true,
  },
  // 是否禁用
  disabled: Boolean,
  // 是否只读
  readonly: Boolean,
  // 最大输入长度
  maxlength: [Number, String],
  // 分隔符样式类
  separatorClass: String,
  // 输入类型限制
  type: {
    type: String,
    default: 'string',
    validator: value => ['string', 'number', 'integer', 'float'].includes(value),
  },
  // 是否严格模式（严格模式下输入非法字符会被阻止）
  strict: {
    type: Boolean,
    default: false,
  },
  // 最小值（仅对数字类型有效）
  min: {
    type: [Number, String],
    default: null,
  },
  // 最大值（仅对数字类型有效）
  max: {
    type: [Number, String],
    default: null,
  },
  // 小数位数（仅对float类型有效）
  decimal: {
    type: Number,
    default: 2,
  },
  // 是否允许负数
  allowNegative: {
    type: Boolean,
    default: true,
  },
})

const emit = defineEmits([
  'update:modelValue',
  'change',
  'start-change',
  'end-change',
  'start-input',
  'end-input',
  'start-focus',
  'end-focus',
  'start-blur',
  'end-blur',
  'start-clear',
  'end-clear',
  'enter',
  'keyup',
  'invalid-input',
])

// 本地值
const startValue = ref('')
const endValue = ref('')
const startInputRef = ref(null)
const endInputRef = ref(null)

// 上一次的有效值
const lastValidStartValue = ref('')
const lastValidEndValue = ref('')

// 监听外部值变化
watch(
  () => props.modelValue,
  newVal => {
    if (Array.isArray(newVal) && newVal.length >= 2) {
      const startVal = newVal[0] ?? ''
      const endVal = newVal[1] ?? ''

      // 验证并设置开始值
      if (isValidInput(startVal, 'start')) {
        startValue.value = formatValue(startVal)
        lastValidStartValue.value = startValue.value
      } else {
        startValue.value = lastValidStartValue.value
      }

      // 验证并设置结束值
      if (isValidInput(endVal, 'end')) {
        endValue.value = formatValue(endVal)
        lastValidEndValue.value = endValue.value
      } else {
        endValue.value = lastValidEndValue.value
      }
    } else {
      startValue.value = ''
      endValue.value = ''
      lastValidStartValue.value = ''
      lastValidEndValue.value = ''
    }
  },
  { immediate: true, deep: true }
)

// 验证输入是否合法
const isValidInput = (value, field) => {
  if (value === '' || value === null || value === undefined) {
    return true
  }

  const strValue = String(value)

  if (props.type === 'number' || props.type === 'integer' || props.type === 'float') {
    // 检查是否只包含数字、小数点、负号
    if (props.type === 'integer') {
      const integerRegex = props.allowNegative ? /^-?\d*$/ : /^\d*$/
      if (!integerRegex.test(strValue)) {
        emitInvalidEvent(field, value, '只能输入整数' + (props.allowNegative ? '，可包含负号' : ''))
        return false
      }
    } else {
      const numberRegex = props.allowNegative ? /^-?\d*\.?\d*$/ : /^\d*\.?\d*$/
      if (!numberRegex.test(strValue)) {
        emitInvalidEvent(field, value, '只能输入数字' + (props.allowNegative ? '，可包含负号和小数点' : '和小数点'))
        return false
      }

      // 检查小数位数
      if (props.type === 'float' && strValue.includes('.')) {
        const decimalPart = strValue.split('.')[1]
        if (decimalPart && decimalPart.length > props.decimal) {
          emitInvalidEvent(field, value, `小数位数不能超过 ${props.decimal} 位`)
          return false
        }
      }
    }

    // 检查范围限制
    const numValue = parseFloat(strValue)
    if (!isNaN(numValue)) {
      if (props.min !== null && numValue < parseFloat(props.min)) {
        emitInvalidEvent(field, value, `不能小于 ${props.min}`)
        return false
      }
      if (props.max !== null && numValue > parseFloat(props.max)) {
        emitInvalidEvent(field, value, `不能大于 ${props.max}`)
        return false
      }
    }
  } else if (props.type === 'string') {
    // 字符串类型默认允许所有输入
  }

  return true
}
// 格式化值
const formatValue = value => {
  if (value === '' || value === null || value === undefined) {
    return ''
  }

  const strValue = String(value)

  switch (props.type) {
    case 'integer': {
      // 使用大括号创建块级作用域
      const intMatch = strValue.match(props.allowNegative ? /^-?\d+/ : /^\d+/)
      return intMatch ? intMatch[0] : ''
    }

    case 'float': {
      // 使用大括号创建块级作用域
      const floatMatch = strValue.match(props.allowNegative ? /^-?\d*\.?\d*/ : /^\d*\.?\d*/)
      if (floatMatch && floatMatch[0]) {
        let formatted = floatMatch[0]
        if (formatted.includes('.')) {
          const parts = formatted.split('.')
          if (parts[1].length > props.decimal) {
            formatted = `${parts[0]}.${parts[1].substring(0, props.decimal)}`
          }
        }
        return formatted
      }
      return ''
    }

    case 'number': {
      // 使用大括号创建块级作用域
      const numMatch = strValue.match(props.allowNegative ? /^-?\d*\.?\d*/ : /^\d*\.?\d*/)
      return numMatch ? numMatch[0] : ''
    }

    default:
      return strValue
  }
}

// 发出无效输入事件
const emitInvalidEvent = (field, value, reason) => {
  emit('invalid-input', { field, value, reason, type: props.type })
}

// 处理开始输入
const handleStartInput = value => {
  if (props.strict && !isValidInput(value, 'start')) {
    // 严格模式下，恢复上一次的有效值
    nextTick(() => {
      startValue.value = lastValidStartValue.value
      if (startInputRef.value) {
        startInputRef.value.setSelectionRange(
          startInputRef.value.selectionStart - 1,
          startInputRef.value.selectionStart - 1
        )
      }
    })
  } else {
    const formattedValue = formatValue(value)
    startValue.value = formattedValue
    lastValidStartValue.value = formattedValue
    emit('start-input', formattedValue)
  }
}

// 处理结束输入
const handleEndInput = value => {
  if (props.strict && !isValidInput(value, 'end')) {
    // 严格模式下，恢复上一次的有效值
    nextTick(() => {
      endValue.value = lastValidEndValue.value
      if (endInputRef.value) {
        endInputRef.value.setSelectionRange(endInputRef.value.selectionStart - 1, endInputRef.value.selectionStart - 1)
      }
    })
  } else {
    const formattedValue = formatValue(value)
    endValue.value = formattedValue
    lastValidEndValue.value = formattedValue
    emit('end-input', formattedValue)
  }
}

// 内部值变化时更新 modelValue
watch([startValue, endValue], () => {
  const newValue = [startValue.value, endValue.value]
  emit('update:modelValue', newValue)
})

// 其他事件处理函数保持不变
const handleStartChange = value => {
  emit('start-change', value)
  emit('change', [value, endValue.value])
}

const handleEndChange = value => {
  emit('end-change', value)
  emit('change', [startValue.value, value])
}

const handleStartFocus = event => {
  emit('start-focus', event)
}

const handleEndFocus = event => {
  emit('end-focus', event)
}

const handleStartBlur = event => {
  // 失焦时进行最终格式化
  if (props.type === 'float' && startValue.value && startValue.value.includes('.')) {
    const parts = startValue.value.split('.')
    if (parts[1] && parts[1].length < props.decimal) {
      startValue.value = startValue.value.padEnd(startValue.value.length + (props.decimal - parts[1].length), '0')
    }
  }
  emit('start-blur', event)
}

const handleEndBlur = event => {
  // 失焦时进行最终格式化
  if (props.type === 'float' && endValue.value && endValue.value.includes('.')) {
    const parts = endValue.value.split('.')
    if (parts[1] && parts[1].length < props.decimal) {
      endValue.value = endValue.value.padEnd(endValue.value.length + (props.decimal - parts[1].length), '0')
    }
  }
  emit('end-blur', event)
}

const handleStartClear = () => {
  lastValidStartValue.value = ''
  emit('start-clear')
}

const handleEndClear = () => {
  lastValidEndValue.value = ''
  emit('end-clear')
}

const handleEnter = event => {
  emit('enter', event, [startValue.value, endValue.value])
}

const handleKeyup = event => {
  emit('keyup', event)
}

// 公开方法
defineExpose({
  clear: () => {
    startValue.value = ''
    endValue.value = ''
    lastValidStartValue.value = ''
    lastValidEndValue.value = ''
  },
  getValues: () => [startValue.value, endValue.value],
  validate: () => {
    return {
      start: isValidInput(startValue.value, 'start'),
      end: isValidInput(endValue.value, 'end'),
      both: isValidInput(startValue.value, 'start') && isValidInput(endValue.value, 'end'),
    }
  },
})
</script>

<style scoped>
/* 样式保持不变 */
.custom-range-input {
  display: inline-block;
  vertical-align: middle;
}
.range-inputs-container {
  display: flex;
  gap: 8px;
  align-items: center;
}
.start-input,
.end-input {
  flex: 1;
}
.separator {
  flex-shrink: 0;
  padding: 0 4px;
  color: var(--el-text-color-secondary);
  white-space: nowrap;
}

/* 禁用状态 */
.range-inputs-container .start-input.is-disabled,
.range-inputs-container .end-input.is-disabled {
  cursor: not-allowed;
  opacity: 0.6;
}
.range-inputs-container .separator.disabled {
  opacity: 0.6;
}
</style>
