<template>
  <v-card elevation="1" class="mb-6">
    <v-card-title>
      <v-avatar size="80"><img :src="profile.avatar_url"></v-avatar>
      <div class="ml-3 text-left">
        <div>{{ profile.name }}</div>
        <div class="text--secondary subtitle-1">{{ profile.slogan }}</div>
      </div>
    </v-card-title>
    <v-card-text>
      <div class="text-left body-1 mt-5">{{ profile.description }}</div>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "Profile",

  data: () => ({
    profile: {}
  }),

  mounted() {
    this.getProfile()
  },

  methods: {
    getProfile() {
      this.utils.GET('/profile').then(res => {
        this.profile = res
      }).catch(err => {
        this.profile = {
          avatar_url: '',
          name: '',
          slogan: '',
          description: '',
        }
      })
    }
  }
}
</script>

<style scoped>

</style>