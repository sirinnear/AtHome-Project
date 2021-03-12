<template>
  <v-container>
    <v-data-table
        v-model="selected"
        :headers="headers"
        :items="certificates"
        :items-per-page="5"
        show-select
        :single-select=false
        item-key="name"
        class="elevation-1 text-left">
<!--      <template v-slot:body="{ items }">-->
<!--        <tbody>-->
<!--          <tr-->
<!--              v-for="item in items"-->
<!--              :key="item.name"-->
<!--          >-->
<!--            <td>{{ item.name }}</td>-->
<!--            <td>{{ item.issuer }}</td>-->
<!--            <td class="text-center">-->
<!--              <v-btn class="ma-0">Export</v-btn>-->
<!--            </td>-->
<!--          </tr>-->
<!--        </tbody>-->
<!--      </template>-->
    </v-data-table>
  </v-container>
</template>

<script>
export default {
  name: "DataTable",
  data() {
    return {
      selected: [],
      headers: [
        {
          text: 'Certificate',
          align: 'start',
          value: 'name',
        },
        { text: 'Issuer', value: 'issuer' },
      ],
      certificates: [
        {
          name: 'A',
          issuer: 'MUIC',
        },
        {
          name: 'B',
          issuer: 'MUIC',
        },
        {
          name: 'C',
          issuer: 'MUIDS',
        },
        {
          name: 'D',
          issuer: 'MUIDS',
        },
      ],
    }
  }
}
</script>

<style scoped>

</style>