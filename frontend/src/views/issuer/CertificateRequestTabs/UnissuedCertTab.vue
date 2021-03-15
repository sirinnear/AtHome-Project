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
          text: 'Student',
          align: 'start',
          value: 'student',
        },
        {text: 'Certificate', value: 'certificate'},
      ],
      certificates: [
        {
          student: 'John Lee',
          certificate: 'Math competition',
          certificateID: 1,
        },
        {
          student: 'Sand Witch',
          certificate: 'Graduation Certificate',
          certificateID: 2,
        },
        {
          student: 'Camelia Anders',
          certificate: 'Singing competition',
          certificateID: 3,
        },
        {
          student: 'Douglas Stein',
          certificate: 'Graduation Certificate',
          certificateID: 4,
        },
        {
          student: 'Abra Cadabra',
          certificate: 'Sports competition',
          certificateID: 5,
        },
        {
          student: 'Bay Leaf',
          certificate: 'Graduation Certificate',
          certificateID: 6,
        },
        {
          student: 'Fedrick Hams',
          certificate: 'Math competition',
          certificateID: 7,
        },
        {
          student: 'Zaynah Ahmad',
          certificate: 'Honors certificate',
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