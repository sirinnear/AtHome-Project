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
          certificate: 'Certificate 9',
        },
        {
          user: 'B',
          certificate: 'Certificate 10',
        },
        {
          user: 'C',
          certificate: 'Certificate 11',
        },
        {
          user: 'D',
          certificate: 'Certificate 12',
        },
      ],
    }
  }
}
</script>

<style scoped>

</style>