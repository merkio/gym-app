<template>
  <v-card
      outlined
      max-width="700"
      class="mx-auto"
  >
    <v-card-title>
      {{ program.title.substring(0, 30) }}
    </v-card-title>

    <v-card-subtitle>
      ${{ product.text }}
    </v-card-subtitle>
    <v-card-actions>
      <v-btn
          color="success"
          outlined
          :to="programDetails"
      >
        <v-icon>fa-glasses</v-icon>
        Details
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import VueTypes from 'vue-types'

export default {
  name: "Program",
  props: {
    program: VueTypes.shape({
      title: VueTypes.string.isRequired,
      text: VueTypes.number.isRequired,
      date: VueTypes.string.isRequired
    })
  }
}
</script>

<style scoped>

</style>
