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
            <UnissuedCertTab />
          </v-tab-item>
          <v-tab-item>
            <IssuedCertTab />
          </v-tab-item>
        </v-tabs-items>
      </v-container>
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/UserAppBar';
import UnissuedCertTab from "@/views/issuer/CertificateRequestTabs/UnissuedCertTab";
import IssuedCertTab from "@/views/issuer/CertificateRequestTabs/IssuedCertTab";

export default {
  name: "CertificateTables",
  components: {
    IssuedCertTab,
    AppBar,
    UnissuedCertTab
  },
  props: {
    authenticated: Boolean,
  },
  data () {
    return {
      tab: null,
    }
  },
}
</script>

<style scoped>
</style>