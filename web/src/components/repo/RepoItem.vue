<template>
  <router-link
    v-if="repo"
    :to="{ name: 'repo', params: { repoId: repo.id } }"
    class="flex flex-col border rounded-md bg-wp-background-100 overflow-hidden p-4 border-wp-background-400 dark:bg-wp-background-200 cursor-pointer hover:shadow-md hover:bg-wp-background-300 dark:hover:bg-wp-background-300"
  >
    <div class="grid grid-cols-[auto,1fr] gap-y-4 items-center">
      <div class="text-wp-text-100 text-lg">{{ `${repo.owner} / ${repo.name}` }}</div>
      <div class="ml-auto text-wp-text-100">
        <div
          v-if="repo.visibility === RepoVisibility.Private"
          :title="`${$t('repo.visibility.visibility')}: ${$t(`repo.visibility.private.private`)}`"
        >
          <Icon name="visibility-private" />
        </div>
        <div
          v-else-if="repo.visibility === RepoVisibility.Internal"
          :title="`${$t('repo.visibility.visibility')}: ${$t(`repo.visibility.internal.internal`)}`"
        >
          <Icon name="visibility-internal" />
        </div>
      </div>

      <div class="col-span-2 text-wp-text-100 flex w-full gap-x-4">
        <template v-if="lastPipeline">
          <div class="flex flex-1 min-w-0 gap-x-1 items-center">
            <PipelineStatusIcon v-if="lastPipeline" :status="lastPipeline.status" />
            <span class="whitespace-nowrap overflow-hidden overflow-ellipsis">{{ shortMessage }}</span>
          </div>

          <div class="flex flex-shrink-0 gap-x-1 items-center ml-auto">
            <Icon name="since" />
            <span>{{ since }}</span>
          </div>
        </template>

        <div v-else class="flex gap-x-2">
          <span>{{ $t('repo.pipeline.no_pipelines') }}</span>
        </div>
      </div>
    </div>
  </router-link>
</template>

<script lang="ts" setup>
import { computed } from 'vue';

import Icon from '~/components/atomic/Icon.vue';
import PipelineStatusIcon from '~/components/repo/pipeline/PipelineStatusIcon.vue';
import usePipeline from '~/compositions/usePipeline';
import type { Repo } from '~/lib/api/types';
import { RepoVisibility } from '~/lib/api/types';

const props = defineProps<{
  repo: Repo;
}>();

const lastPipeline = computed(() => props.repo.last_pipeline_item);
const { since, shortMessage } = usePipeline(lastPipeline);
</script>
