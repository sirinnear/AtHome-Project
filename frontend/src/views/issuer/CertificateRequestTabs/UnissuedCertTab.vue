<template>
  <v-container>
    <v-data-table
        v-model="selected"
        :headers="headers"
        :items="certificates"
        :items-per-page="5"
        show-select
        :single-select=false
        item-key="certificateID"
        class="elevation-1 text-left">
      <template v-slot:top>
        <v-toolbar flat dense>
          <v-btn text dense color="#ea292f" v-if="selected.length > 0">
            <v-icon
                left
            >
              mdi-export
            </v-icon>
            Export Selected
          </v-btn>
        </v-toolbar>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
export default {
  name: "UnissuedCertTab",
  data () {
    return {
      selected: [],
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
          certificateID: 1,
        },
        {
          user: 'B',
          certificate: 'MUIC',
          certificateID: 2,
        },
        {
          user: 'C',
          certificate: 'MUIDS',
          certificateID: 3,
        },
        {
          user: 'D',
          certificate: 'MUIDS',
          certificateID: 4,
        },
        {
          user: 'A',
          certificate: 'MUIC',
          certificateID: 5,
        },
        {
          user: 'B',
          certificate: 'MUIC',
          certificateID: 6,
        },
        {
          user: 'C',
          certificate: 'MUIDS',
          certificateID: 7,
        },
        {
          user: 'D',
          certificate: 'MUIDS',
          certificateID: 8,
        },
      ],
    }
  }
}
</script>

<style scoped>

</style>