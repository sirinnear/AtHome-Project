<template>
  <div>
    <AppBar :authenticated=true />
    <v-content>
      <v-container>
        <v-tabs v-model="tab" centered>
          <v-tab>Unissued Certificates</v-tab>
          <v-tab>Issued Certificates</v-tab>
        </v-tabs>
        <v-tabs-items v-model="tab">
          <v-tab-item>
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
              </v-data-table>
            </v-container>
          </v-tab-item>
          <v-tab-item>
            <v-container>
              <v-data-table
                  :headers="headers"
                  :items="certificates"
                  :items-per-page="5"
                  class="elevation-1 text-left">
              </v-data-table>
            </v-container>
          </v-tab-item>
        </v-tabs-items>
      </v-container>
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/AppBar';
// import CertificateRequestTab from "@/views/issuer/CertificateRequestTab";

export default {
  name: "Dashboard",
  components: {
    AppBar,
    // CertificateRequestTab
  },
  props: {
    authenticated: Boolean,
  },
  data () {
    return {
      tab: null,
      selected: [],
      headers: [
        {
          text: 'Certificate',
          align: 'start',
          value: 'name',
        },
        { text: 'Issuer', value: 'issuer' },
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
  },
}
</script>

<style scoped>
</style>