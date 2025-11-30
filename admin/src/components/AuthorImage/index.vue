<template>
  <div
    :style="{
      color: '#fff',
      borderRadius: '50%',
      width: size + 'px',
      height: size + 'px',
      display: 'block',
      textAlign: 'center',
      lineHeight: size + 'px',
      fontSize: size * 0.8 + 'px',
      background: color(name),
    }"
  >
    {{ name.slice(0, 1) }}
  </div>
</template>

<script lang="ts" setup>
defineOptions({
  name: 'AuthorImage',
})
defineProps({
  name: {
    type: String,
    default: '',
  },
  size: {
    type: Number,
    default: 16,
  },
})

const color = (name: string) => {
  if (name) {
    let num = ''
    for (let i = 0; i < name.length; i++) {
      num += name[i].charCodeAt(0).toString()
    }
    const r = Math.min(100 + (parseInt(num.slice(0, 2), 10) % 55), 255)
    const g = Math.min(100 + (parseInt(num.slice(2, 4), 10) % 55), 255)
    const b = Math.min(100 + (parseInt(num.slice(4, 6), 10) % 55), 255)
    return `rgb(${r}, ${g}, ${b})`
  } else {
    return 'rgb(125, 125, 125)'
  }
}
</script>
