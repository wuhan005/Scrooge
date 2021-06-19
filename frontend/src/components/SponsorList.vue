<template>
  <v-card elevation="1">
    <v-data-table
        :headers="sponsorHeaders"
        :items="sponsorList"
        :expanded.sync="expanded"
        item-key="name"
        show-expand
        hide-default-footer
    >
      <template v-slot:item.index="{ item }">
        {{ item.index }}
      </template>
      <template v-slot:item.name="{ item }">
        {{ item.name === '' ? '不愿意透露姓名的好心人' : item.name }}
      </template>

      <template v-slot:item.subtotal="{ item }">
        {{ item.subtotal / 100 }}
      </template>

      <template v-slot:expanded-item="{ item }">
        <td :colspan="sponsorHeaders.length">
          <v-simple-table dense class="mt-3 mb-3">
            <tbody>
            <tr v-for="i in item.invoices">
              <td width="200px">{{ (new Date(i.created_at)).toISOString().slice(0, 10) }}</td>
              <td>{{ i.price_cents / 100 }}</td>
              <td>{{ i.comment }}</td>
            </tr>
            </tbody>
          </v-simple-table>
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
      sponsorHeaders: [
        {text: '#', value: 'index'},
        {text: '老板', value: 'name'},
        {text: '总额（元）', value: 'subtotal'},
        {text: '次数', value: 'count'},
        {text: '', value: 'data-table-expand'},
      ],
      sponsorList: [],
    }
  },

  mounted() {
    this.getSponsorList()
  },

  methods: {
    getSponsorList() {
      this.utils.GET('/sponsor_list').then(res => {
        this.sponsorList = res
      }).catch(err => {
        this.messageBar = true
        this.message = err.response.data.msg
      })
    },
  }
}
</script>

<style scoped>

</style>