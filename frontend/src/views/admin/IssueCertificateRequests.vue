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
          text: 'User',
          align: 'start',
          value: 'user',
        },
        {text: 'Certificate', value: 'certificate'},
        {text: 'Issuer', value: 'issuer'},
      ],
      certificates: [
        {
          user: 'A',
          certificate: 'Certificate 1',
          issuer: 'MUIC',
          certificateID: 1,
        },
        {
          user: 'B',
          certificate: 'Certificate 2',
          issuer: 'MUIC',
          certificateID: 2,
        },
        {
          user: 'C',
          certificate: 'Certificate 3',
          issuer: 'MUIDS',
          certificateID: 3,
        },
        {
          user: 'D',
          certificate: 'Certificate 4',
          issuer: 'MUIDS',
          certificateID: 4,
        },
        {
          user: 'E',
          certificate: 'Certificate 5',
          issuer: 'MUIC',
          certificateID: 5,
        },
        {
          user: 'F',
          certificate: 'Certificate 6',
          issuer: 'MUIC',
          certificateID: 6,
        },
        {
          user: 'G',
          certificate: 'Certificate 7',
          issuer: 'MUIDS',
          certificateID: 7,
        },
        {
          user: 'H',
          certificate: 'Certificate 8',
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