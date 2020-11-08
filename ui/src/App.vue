<template>
  <v-app>
    <TopNav/>
    <v-main>
      <router-view/>
    </v-main>
    <BottomNav/>
    <v-snackbar
        v-model="snackbar.show"
        :multi-line="true"
        :right="true"
        :top="true"
        :timeout="6000"
        :color="snackbar.variant"
    >
      {{ snackbar.message }}
      <v-btn
          text
          dark
          @click="updateSnackbar({show: false})">
        Close
      </v-btn>
    </v-snackbar>
  </v-app>
</template>

<script>
import TopNav from "@/components/TopNav";
import BottomNav from './components/BottomNav';
import {mapState, mapMutations} from 'vuex'

export default {
  name: 'App',

  components: {
    TopNav,
    BottomNav,
  },
  computed: {
    ...mapState(['snackbar'])
  },
  methods: {
    ...mapMutations(['updateSnackbar'])
  }
};
</script>
