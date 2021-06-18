<template>
  <v-card elevation="1">
    <v-data-table
        :headers="dessertHeaders"
        :items="desserts"
        :expanded.sync="expanded"
        item-key="id"
        show-expand
        hide-default-footer
    >
      <template v-slot:expanded-item="{ headers, item }">
        <td :colspan="headers.length">
          More info about {{ item.sponsor }}
        </td>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  name: "SponsorList",
  data() {
    return {
      expanded: [],
      dessertHeaders: [
        {text: '老板', value: 'sponsor'},
        {text: '总额（元）', value: 'subtotal'},
        {text: '次数', value: 'count'},
        {text: '', value: 'data-table-expand'},
      ],
      desserts: [
        {
          id: 1,
          sponsor: 'Frozen Yogurt',
          subtotal: 159,
          count: 6,
        },
        {
          id: 2,
          sponsor: 'Ice cream sandwich',
          subtotal: 237,
          count: 9,
        },
        {
          id: 3,
          sponsor: 'Eclair',
          subtotal: 262,
          count: 16,
        },
      ],
    }
  },
}
</script>

<style scoped>

</style>