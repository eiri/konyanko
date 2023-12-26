POST http://localhost:2315/graphql
```graphql
{
  items(
    where: {and: [
        {publishDateGTE: "2023-12-09T00:00:00Z"},
        {publishDateLT: "2023-12-10T00:00:00Z"},
        {hasEpisode: true}
    ]}
    orderBy: {direction: DESC, field: PUBLISH_DATE}
  ) {
    totalCount
    edges {
      node {
        fileName
        fileSize
        publishDate
        episode {
          episodeNumber
          anime {
            title
          }
        }
      }
    }
  }
}
```
HTTP 200