<template>
  <v-container>
    <v-data-table
        :headers="headers"
        :items="certificates"
        :items-per-page="5"
        class="elevation-1 text-left">
    </v-data-table>
  </v-container>
</template>

<script>
export default {
  name: "IssuedCertTab",
  data () {
    return {
      headers: [
        {
          text: 'Certificate',
          align: 'start',
          value: 'name',
        },
        {text: 'Issuer', value: 'issuer'},
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