<template>
  <component
    :is="to === undefined ? 'button' : httpLink ? 'a' : 'router-link'"
    v-bind="btnAttrs"
    class="relative flex flex-shrink-0 whitespace-nowrap items-center py-1 px-2 rounded-md border shadow-sm cursor-pointer transition-all duration-150 overflow-hidden disabled:opacity-50 disabled:cursor-not-allowed"
    :class="{
      'bg-wp-control-neutral-100 hover:bg-wp-control-neutral-200 border-wp-control-neutral-300 text-wp-text-100':
        color === 'gray',
      'bg-wp-control-ok-100 hover:bg-wp-control-ok-200 border-wp-control-ok-300 text-white': color === 'green',
      'bg-wp-control-info-100 hover:bg-wp-control-info-200 border-wp-control-info-300 text-white': color === 'blue',
      'bg-wp-control-error-100 hover:bg-wp-control-error-200 border-wp-control-error-300 text-white': color === 'red',
      ...passedClasses,
    }"
    :title="title"
    :disabled="disabled"
  >
    <slot>
      <Icon v-if="startIcon" :name="startIcon" class="!w-6 !h-6" :class="{ invisible: isLoading, 'mr-1': text }" />
      <span :class="{ invisible: isLoading }">{{ text }}</span>
      <Icon v-if="endIcon" :name="endIcon" class="ml-2 w-6 h-6" :class="{ invisible: isLoading }" />
      <div
        v-if="isLoading"
        class="absolute left-0 top-0 right-0 bottom-0 flex items-center justify-center"
        :class="{
          'bg-wp-control-neutral-200': color === 'gray',
          'bg-wp-control-ok-200': color === 'green',
          'bg-wp-control-info-200': color === 'blue',
          'bg-wp-control-error-200': color === 'red',
        }"
      >
        <Icon name="spinner" />
      </div>
    </slot>
  </component>
</template>

<script lang="ts" setup>
import { computed, useAttrs } from 'vue';
import type { RouteLocationRaw } from 'vue-router';

import type { IconNames } from '~/components/atomic/Icon.vue';
import Icon from '~/components/atomic/Icon.vue';

const props = withDefaults(
  defineProps<{
    text?: string;
    title?: string;
    disabled?: boolean;
    to?: RouteLocationRaw;
    color?: 'blue' | 'green' | 'red' | 'gray';
    startIcon?: IconNames;
    endIcon?: IconNames;
    isLoading?: boolean;
  }>(),
  {
    text: undefined,
    title: undefined,
    to: undefined,
    color: 'gray',
    startIcon: undefined,
    endIcon: undefined,
  },
);

const httpLink = computed(() => typeof props.to === 'string' && props.to.startsWith('http'));

const btnAttrs = computed(() => {
  if (props.to === null) {
    return { type: 'button' };
  }

  if (httpLink.value) {
    return { href: props.to };
  }

  return { to: props.to };
});

const attrs = useAttrs();
const passedClasses = computed(() => {
  const classes: Record<string, boolean> = {};
  const origClass = (attrs.class as string) || '';
  origClass.split(' ').forEach((c) => {
    classes[c] = true;
  });
  return classes;
});
</script>
