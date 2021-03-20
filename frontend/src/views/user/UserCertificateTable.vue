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
              <v-btn text dense color="#ea292f" v-if="selected.length > 0" @click="clearSelected">
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
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/UserAppBar';

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
          text: 'Certificate',
          align: 'start',
          value: 'certificate'},
        {text: 'Issuer', value: 'issuer'},
      ],
      certificates: [
        {
          certificate: 'Math competition',
          issuer: 'MUIC',
          certificateID: 1,
        },
        {
          certificate: 'Honors certificate',
          issuer: 'MUIC',
          certificateID: 2,
        },
        {
          certificate: 'Graduation certificate',
          issuer: 'MUIDS',
          certificateID: 3,
        },
        {
          certificate: 'Honors certificate',
          issuer: 'MUIDS',
          certificateID: 4,
        },
        {
          certificate: 'Sports competition',
          issuer: 'MUIC',
          certificateID: 5,
        },
        {
          certificate: 'Math competition',
          issuer: 'MUIC',
          certificateID: 6,
        },
        {
          certificate: 'Math competition',
          issuer: 'MUIDS',
          certificateID: 7,
        },
        {
          certificate: 'Honors certificate',
          issuer: 'MUIDS',
          certificateID: 8,
        },
      ],
    }
  },
  methods: {
    clearSelected() {
      this.selected = [];
    }
  }
}
</script>

<style scoped>
</style>