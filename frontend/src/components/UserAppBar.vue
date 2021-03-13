<template>
  <v-app-bar flat color="#ea292f" app>
    <v-img
        contain
        max-width="100px"
        class="mr-5"
        alt="AtHome logo"
        src="../assets/athome-logo.png" />
    <v-btn
        text
        dark
        class="white--text"
        v-if="authenticated"
        to="/issuer/certificates">
      Certificates
    </v-btn>
    <v-btn
        text
        dark
        class="white--text"
        v-if="authenticated"
        to="/issuer/transfers">
      Request Transfer
    </v-btn>
    <RequestTransferDialog v-if="authenticated" />
    <v-spacer />
    <v-btn text dark v-if="authenticated">
      Logout
      <v-icon
          dark
          right
      >
        mdi-logout
      </v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
import RequestTransferDialog from "@/views/user/RequestTransferDialog";
export default {
  name: "AppBar",
  components: {RequestTransferDialog},
  props: ['authenticated', 'current']
}
</script>

<style scoped>
  /*.v-btn {*/
  /*  color: white !important;*/
  /*}*/
</style>