<script setup>
import { computed } from 'vue'
import { gql } from '@apollo/client/core'
import { useQuery } from '@vue/apollo-composable'

import Card from './Card.vue'

const ALL_ANIME = gql`
  query AllAnime {
    animes(orderBy: {direction: ASC, field:TITLE}) {
      edges {
        node {
          id
          title
          episodes(where: {hasReleaseGroup: true}) {
            edges {
              node {
                animeSeason
                episodeNumber
                item {
                  downloadURL
                  viewURL
                  fileSize
                  publishDate
                }
                releaseGroup {
                  name
                }
              }
            }
          }
        }
      }
    }
  }
`

const { result, loading, error } = useQuery(ALL_ANIME)
</script>

<template>
  <article v-if="error">Error</article>
  <article v-if="loading">Loading...</article>
  <article v-else class="columns columns-3 gap-4 mt-4">
    <Card v-for="edge in result.animes.edges" :node="edge.node" :key="edge.node.id" />
  </article>
</template>
