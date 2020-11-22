<template>
  <v-row sm="6">
    <v-col sm="2" offset-lg="1">
      <v-container>
        <h1>Filters</h1>
      </v-container>
    </v-col>
    <v-col>
      <v-flex sm4 >
      <v-container align-self-start py-5 ma-5 sm3 v-for="exercise in exercises" :key="exercise.id">
        <ShortViewExercise :exercise="exercise"/>
      </v-container>
      </v-flex>
    </v-col>
  </v-row>
</template>

<script>
import exercises from "@/store/exercises";
import ShortViewExercise from "@/components/exercises/ShortViewExercise";

export default {
  name: "Exercises",
  components: {
    ShortViewExercise,
  },
  data() {
    return {
      exercises: [],
    };
  },
  created() {
    this.initialize();
  },
  methods: {
    initialize() {
      this.exercises = exercises;
    },
  },
};
</script>

<style scoped>
</style>
