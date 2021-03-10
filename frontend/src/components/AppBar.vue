<template>
  <v-app-bar flat color="#ea292f" app>
    <v-img
        contain
        max-width="100px"
        alt="AtHome logo"
        src="../assets/athome-logo.png" />
    <v-spacer />
    <v-btn plain dark v-if="authenticated">
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
export default {
  name: "AppBar",
  props: ['authenticated']
}
</script>

<style scoped>
  /*.v-btn {*/
  /*  color: white !important;*/
  /*}*/
</style>