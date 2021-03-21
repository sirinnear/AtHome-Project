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
          <v-text-field v-model="user.email" outlined label="Email" type="email"/>
          <v-text-field v-model="user.password" outlined label="Password" type="password"/>
          <v-btn dark color="#ea292f" @click="login">Login</v-btn>
        </v-container>
      </v-card>
    </v-content>
  </div>
</template>

<script>
import AppBar from '@/components/IssuerAppBar';
import api from '@/services';

export default {
  name: "Login",
  components: {
  AppBar,
  },
  props: {
  authenticated: Boolean,
  },
  data () {
    return {
      user: {
        email: '',
        password: '',
      }
    }
  },
  methods: {
    login () {
      const data = {
        email: this.user.email,
        password: this.user.password,
      };
      api.login(data)
        .then((response) => {
          const role = response.data.role;
          if (role === 'student') {
            this.$router.push({ name: 'StudentCert', params: {name: response.data.name}});
            console.log(response.data.name);
          }
          else if (role === 'issuer') {
            this.$router.push({ name: 'IssuerCert', params: {name: response.data.name}})
          }
          else if (role === 'admin') {
            this.$router.push('/admin')
          }
          else {
            alert('There exists no such user role!')
          }
        })
        .catch((e) => {
          alert(e.response.data.error);
        });
    }
  }
}
</script>

<style scoped>
</style>