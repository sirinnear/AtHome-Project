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
          text: 'Student',
          align: 'start',
          value: 'student',
        },
        {text: 'Certificate', value: 'certificate'},
      ],
      certificates: [
        {
          student: 'Bay Leaf',
          certificate: 'Honors certificate',
        },
        {
          student: 'Sand Witch',
          certificate: 'Cooking competition',
        },
        {
          student: 'John Lee',
          certificate: 'Sports competition',
        },
        {
          student: 'Abra Cadabra',
          certificate: 'Talent show contest',
        },
      ],
    }
  }
}
</script>

<style scoped>

</style>