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
          text: 'User',
          align: 'start',
          value: 'user',
        },
        {text: 'Certificate', value: 'certificate'},
      ],
      certificates: [
        {
          user: 'A',
          certificate: 'MUIC',
        },
        {
          user: 'B',
          certificate: 'MUIC',
        },
        {
          user: 'C',
          certificate: 'MUIDS',
        },
        {
          user: 'D',
          certificate: 'MUIDS',
        },
      ],
    }
  }
}
</script>

<style scoped>

</style>