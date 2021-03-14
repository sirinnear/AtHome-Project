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
          <v-btn text dense color="#ea292f" v-if="selected.length > 0" @click="deleteSelected">
            <v-icon
                left
            >
              mdi-note-text-outline
            </v-icon>
            Request Selected
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
          certificate: 'Certificate 1',
          certificateID: 1,
        },
        {
          user: 'B',
          certificate: 'Certificate 2',
          certificateID: 2,
        },
        {
          user: 'C',
          certificate: 'Certificate 3',
          certificateID: 3,
        },
        {
          user: 'D',
          certificate: 'Certificate 4',
          certificateID: 4,
        },
        {
          user: 'A',
          certificate: 'Certificate 5',
          certificateID: 5,
        },
        {
          user: 'B',
          certificate: 'Certificate 6',
          certificateID: 6,
        },
        {
          user: 'C',
          certificate: 'Certificate 7',
          certificateID: 7,
        },
        {
          user: 'D',
          certificate: 'Certificate 8',
          certificateID: 8,
        },
      ],
    }
  },
  methods: {
    deleteSelected() {
      this.certificates = this.certificates.filter(r => this.selected.indexOf(r) < 0);
      this.selected = [];
    }
  }
}
</script>

<style scoped>

</style>