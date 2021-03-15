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
        to="/user">
      Certificates
    </v-btn>
    <RequestTransferDialog v-if="authenticated" />
    <v-spacer />
    <v-row align="center" justify="center" no-gutters>
      <v-spacer />
      <v-col cols="1">
        <v-avatar>
          <v-icon dark x-large>
            mdi-account-circle
          </v-icon>
        </v-avatar>
      </v-col>
      <v-col>
        <span class="white--text">John Lee</span>
      </v-col>
    </v-row>
    <v-btn text dark v-if="authenticated" to="/login">
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