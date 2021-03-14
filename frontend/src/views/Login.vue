<template>
  <div>
    <AppBar :authenticated=false />
    <v-content>
      <v-card
          outlined
          class="mt-10 mx-auto elevation-1 rounded-sm"
          max-width="50%">
        <v-card-title class="justify-center text-h5">Login</v-card-title>
        <v-container class="px-10 py-5">
          <v-text-field outlined label="Email" type="email"/>
          <v-text-field outlined label="Password" type="password"/>
          <v-btn dark color="#ea292f">Login</v-btn>
        </v-container>
      </v-card>
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/IssuerAppBar';

export default {
  name: "Login",
  components: {
  AppBar,
  },
  props: {
  authenticated: Boolean,
  },
}
</script>

<style scoped>
</style>