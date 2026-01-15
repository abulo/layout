<!-- 只允许输入字符串 -->

<SimpleRangeInput
  v-model="range1"
  type="string"
  range-separator="到"
  start-placeholder="最小值"
  end-placeholder="最大值"
/>

<!-- 只允许输入数字（可含小数） -->

<SimpleRangeInput
  v-model="range2"
  type="number"
  range-separator="到"
  start-placeholder="最小值"
  end-placeholder="最大值"
/>

<!-- 只允许输入整数 -->

<SimpleRangeInput
  v-model="range3"
  type="integer"
  range-separator="到"
  start-placeholder="最小值"
  end-placeholder="最大值"
  :allow-negative="false"
/>

<!-- 浮点数，限制2位小数 -->

<SimpleRangeInput
  v-model="range4"
  type="float"
  :decimal="2"
  range-separator="到"
  start-placeholder="最小值"
  end-placeholder="最大值"
  :min="0"
  :max="100"
  :strict="true"
/>
