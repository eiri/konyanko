import { ApolloClient, InMemoryCache } from '@apollo/client/core'
import { DefaultApolloClient } from '@vue/apollo-composable'
import { createApp, provide, h } from 'vue'

import App from './App.vue'

const cache = new InMemoryCache()

const apolloClient = new ApolloClient({
  uri: 'http://localhost:2315/graphql',
  cache,
})

const app = createApp({
  setup () {
    provide(DefaultApolloClient, apolloClient)
  },
  render: () => h(App),
})

app.mount('#app')
