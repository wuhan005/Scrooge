<template>
  <v-app style="background: #F0F8FF">
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  data: () => ({}),
};
</script>


<style scoped>
body {
  background-color: #F0F8FF;
}
</style>
