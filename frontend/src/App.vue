<template>
  <v-app id="inspire">
    <v-app-bar
        app
        color="white"
        flat
    >
      <v-avatar
          :color="$vuetify.breakpoint.smAndDown ? 'grey darken-1' : 'transparent'"
          size="32"
      ></v-avatar>

      <v-tabs
          centered
          class="ml-n9"
          color="grey darken-1"
      >
        <v-tab
            v-for="tab in tabs"
            :key="tab.name"
            :to="tab.to"
        >
          {{ tab.name }}
        </v-tab>
      </v-tabs>

      <v-avatar
          class="hidden-sm-and-down"
          color="grey darken-1 shrink"
          size="32"
      ></v-avatar>
    </v-app-bar>

    <v-main class="grey lighten-3">
      <router-view>
      </router-view>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    tabs: [
      {name: 'Dashboard', to: '/'},
      {name: 'Workouts', to: '/workouts'},
      {name: 'Exercises', to: '/exercises'},
      {name: 'Results', to: '/results'},
      {name: 'Profile', to: '/profile'},
    ]
  })
}
</script>
