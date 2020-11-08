<template>
  <v-footer
      color="grey"
      padless
  >
    <v-row
        justify="center"
    >
      <v-btn
          v-for="link in links"
          :key="link"
          color="white"
          text
          class="my-2"
      >
        {{ link }}
      </v-btn>
      <v-col
          class="lighten-1 grey py-4 text-center white--text"
          cols="12"
      >
        {{ new Date().getFullYear() }} — <strong>GYM TRAININGS</strong>
      </v-col>
    </v-row>
  </v-footer>
</template>

<script>
export default {
  name: "BottomNav",
  data: () => ({
    links: [
      'Home',
      'About',
      'Blog',
      'Contact Us',
    ],
  }),
}
</script>

<style scoped>

</style>
