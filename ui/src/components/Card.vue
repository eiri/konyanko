<script setup>
import { ref } from 'vue'
import humanizeDuration from 'humanize-duration'

defineProps({
  node: {
    title: String,
    episodes: Object
  },
})

const timeAgo = (past) => {
  const dur = Date.now() - Date.parse(past)
  const opts = { units: ["d", "h", "m"], largest: 2, round: true }
  return humanizeDuration(dur, opts)
}
</script>

<template>
  <div class="bg-white mt-4 p-4 rounded-xl overflow-hidden shadow border border-gray-200 text-gray-500">
    <h4 class="font-semibold text-lg lading-tight truncate">
      {{ node.title }}
    </h4>

    <div class="text-sm font-medium text-center border-b border-gray-200">
      <ul class="flex flex-wrap -mb-px">
        <li v-for="episode in node.episodes.edges" class="me-2">
          <a href="#" class="active text-teal-600 border-teal-600 inline-block p-4 border-b-2 rounded-t-lg">{{ episode.node.animeSeason }} &bull; {{ episode.node.episodeNumber }}</a>
        </li>
      </ul>
    </div>

    <div v-for="episode in node.episodes.edges">
      <ul class="max-w-md space-y-1 list-inside">
        <li class="flex items-center">
          <a :href="episode.node.item.downloadURL" class="mr-2 truncate">
          {{ episode.node.releaseGroup.name }}
            <span class="text-xs text-gray-400 ml-2">
              {{ timeAgo(episode.node.item.publishDate) }}
              <!-- { humanize.Bytes(uint64(i.FileSize)) } -->
            </span>
          </a>
          <a :href="episode.node.item.viewURL" target="_blank">
            <svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" class="w-5 h-5 fill-none stroke-current stroke-gray-400 stroke-2" width="24" height="24" viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round">
              <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
              <path d="M12 6h-6a2 2 0 0 0 -2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-6" />
              <path d="M11 13l9 -9" />
              <path d="M15 4h5v5" />
            </svg>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>
