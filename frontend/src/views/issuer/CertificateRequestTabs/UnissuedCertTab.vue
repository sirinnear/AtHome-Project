<template>
  <v-container>
    <v-data-table
        v-model="selected"
        :headers="headers"
        :items="certificates"
        :items-per-page="5"
        show-select
        :single-select=false
        item-key="user"
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