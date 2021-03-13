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
            item-key="userID"
            class="elevation-1 text-left">
          <template v-slot:top>
            <v-toolbar flat dense>
              <v-btn text dense color="#ea292f" v-if="selected.length > 0">
                <v-icon
                    left
                >
                  mdi-check
                </v-icon>
                Approve Selected
              </v-btn>
            </v-toolbar>
          </template>
        </v-data-table>
      </v-container>
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/AppBar';

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
          certificate: 'MUIC',
          issuer: 'MUIC',
          certificateID: 1,
        },
        {
          user: 'B',
          certificate: 'MUIC',
          issuer: 'MUIC',
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
  },
}
</script>

<style scoped>
</style>