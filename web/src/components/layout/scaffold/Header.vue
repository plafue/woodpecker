<template>
  <header
    class="bg-wp-background-100 border-b-1 border-wp-background-400 dark:border-wp-background-100 dark:bg-wp-background-300 text-wp-text-100"
    :class="{ 'md:px-4': fullWidth }"
  >
    <Container :full-width="fullWidth" class="!py-0">
      <div class="flex w-full md:items-center flex-col py-3 gap-2 md:gap-10 md:flex-row md:justify-between">
        <div
          class="flex items-center content-start min-h-10"
          :class="{
            'md:flex-1': searchBoxPresent,
          }"
        >
          <IconButton
            v-if="goBack"
            icon="back"
            :title="$t('back')"
            class="flex-shrink-0 mr-2 <md:hidden md:justify-between w-8 h-8"
            @click="goBack"
          />
          <h1 class="flex text-xl min-w-0 text-wp-text-100 items-center gap-x-2">
            <slot name="title" />
          </h1>
        </div>
        <TextField
          v-if="searchBoxPresent"
          class="w-auto <md:w-full flex-grow <md:order-3"
          :aria-label="$t('search')"
          :placeholder="$t('search')"
          :model-value="search"
          @update:model-value="(value: string) => $emit('update:search', value)"
        />
        <div
          v-if="$slots.headerActions"
          class="flex items-center md:justify-end gap-x-2 min-w-0"
          :class="{
            'md:flex-1': searchBoxPresent,
          }"
        >
          <slot name="headerActions" />
        </div>
      </div>

      <div v-if="enableTabs" class="flex md:items-center flex-col py-2 md:flex-row md:justify-between md:py-0">
        <Tabs class="<md:order-2" />
        <div v-if="$slots.headerActions" class="flex content-start md:justify-end">
          <slot name="tabActions" />
        </div>
      </div>
    </Container>
  </header>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import IconButton from '~/components/atomic/IconButton.vue';
import TextField from '~/components/form/TextField.vue';
import Container from '~/components/layout/Container.vue';

import Tabs from './Tabs.vue';

const props = defineProps<{
  goBack?: () => void;
  enableTabs?: boolean;
  search?: string;
  fullWidth?: boolean;
}>();

defineEmits<{
  (event: 'update:search', query: string): void;
}>();

const searchBoxPresent = computed(() => props.search !== undefined);
</script>
