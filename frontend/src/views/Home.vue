<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Car Management"/>
    
    <div v-if="!$auth.loading">
      <!-- show login when not authenticated -->
      <button v-if="!$auth.isAuthenticated" @click="login">Log in</button>
      <!-- show logout when authenticated -->
      <button v-if="$auth.isAuthenticated" @click="logout">Log out</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'home',
  components: {
    HelloWorld
  },
  methods: {
    login() {
      this.$auth.loginWithRedirect();
    },

    logout() {
      this.$auth.logout({
        returnTo: window.location.origin
      });
    }
  }
}
</script>
