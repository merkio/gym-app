<template>
  <div>
    <h1>Filters</h1>
    <v-spacer />
    <v-form
      outline
    >

    </v-form>
  </div>
</template>

<script>
import groups from '@/store/groups'

export default {
  name: "Sidebar",
  data() {
    return {
      groups: []
    }
  },
  created() {
    this.initialize()
  },
  methods: {
    initialize() {
      this.groups = groups
    }
  }
}


</script>

<style scoped>

</style>
