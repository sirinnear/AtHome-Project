<template>
  <AppBar :authenticated=true />
</template>

<script>
import AppBar from '@/components/AppBar';

export default {
name: "Login",
  components: {
  AppBar,
  },
  props: {
  authenticated: Boolean,
  }
}
</script>

<style scoped>

</style>