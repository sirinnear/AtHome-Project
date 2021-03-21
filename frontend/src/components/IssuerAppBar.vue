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
      Transfer Requests
    </v-btn>
    <v-spacer />
    <v-row align="center" justify="center" no-gutters v-if="authenticated">
      <v-spacer />
      <v-col cols="1">
        <v-avatar>
          <img src="../assets/muic-logo.png" alt="MUIC Logo">
        </v-avatar>
      </v-col>
      <v-col>
        <span class="white--text">{{ name }}</span>
      </v-col>
    </v-row>
    <v-btn text dark v-if="authenticated" @click="logout">
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
import api from "@/services";

export default {
  name: "AppBar",
  props: ['authenticated', 'name'],
  methods: {
    logout() {
      api.logout();
      this.$router.push('/login');
    }
  }
}
</script>

<style scoped>
  /*.v-btn {*/
  /*  color: white !important;*/
  /*}*/
</style>