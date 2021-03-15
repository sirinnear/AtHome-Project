<template>
  <div>
    <AppBar :authenticated=true />
    <v-content>
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
              <div v-if="selected.length > 0">
                <v-btn text dense color="#ea292f" @click="deleteSelected">
                  <v-icon
                      left
                  >
                    mdi-check
                  </v-icon>
                  Issue Selected
                </v-btn>
                <v-btn text dense color="#ea292f" @click="deleteSelected">
                  <v-icon
                      left
                  >
                    mdi-cancel
                  </v-icon>
                  Reject Selected
                </v-btn>
              </div>
            </v-toolbar>
          </template>
        </v-data-table>
      </v-container>
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/AdminAppBar';

export default {
  name: "CertificateTables",
  components: {
    AppBar,
  },
  props: {
    authenticated: Boolean,
  },
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
        {text: 'Issuer', value: 'issuer'},
      ],
      certificates: [
        {
          student: 'John Lee',
          certificate: 'Math competition',
          issuer: 'MUIC',
          certificateID: 1,
        },
        {
          student: 'Sand Witch',
          certificate: 'Graduation Certificate',
          issuer: 'MUIC',
          certificateID: 2,
        },
        {
          student: 'Ben Trillion',
          certificate: 'Honors certificate',
          issuer: 'MUIDS',
          certificateID: 3,
        },
        {
          student: 'Dylan Villain',
          certificate: 'Graduation Certificate',
          issuer: 'MUIDS',
          certificateID: 4,
        },
        {
          student: 'Abra Cadabra',
          certificate: 'Sports competition',
          issuer: 'MUIC',
          certificateID: 5,
        },
        {
          student: 'Douglas Stein',
          certificate: 'Graduation Certificate',
          issuer: 'MUIC',
          certificateID: 6,
        },
        {
          student: 'Greg Wales',
          certificate: 'Graduation Certificate',
          issuer: 'MUIDS',
          certificateID: 7,
        },
        {
          student: 'Haru Oyama',
          certificate: 'Sports competition',
          issuer: 'MUIDS',
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
  },
}
</script>

<style scoped>
</style>