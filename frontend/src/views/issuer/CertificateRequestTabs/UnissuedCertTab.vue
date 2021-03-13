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