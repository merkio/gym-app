import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        programs: [],
        program: {},
        exercises: [],
        exercise: {},
        results: [],
        result: {},
        snackbar: {
            show: false,
            variant: 'success',
            message: 'Success!'
        }
    },
    mutations: {
        updateSnackbar(state, settings) {
            state.snackbar = {
                ...state.snackbar,
                ...settings
            }
        }
    },
    actions: {}
})
