import Home from '@/views/Home.vue'
import Programs from '@/views/Programs.vue'
import Exercises from '@/views/Exercises.vue'
import Results from '@/views/Results.vue'
import About from '@/views/About.vue'
import FullViewExercise from '@/components/exercises/FullViewExercise'
import ProgramDetails from "@/components/programs/ProgramDetails";

export default [
  { path: '/', name: 'home', component: Home },
  { path: '/programs', name: 'programs', component: Programs },
  { path: '/programs/:id', name: 'programDetails', component: ProgramDetails},
  { path: '/exercises', name: 'exercises', component: Exercises },
  { path: '/exercises/:id', name: 'exercise', component: FullViewExercise},
  { path: '/results', name: 'results', component: Results },
  { path: '/about', name: 'about', component: About }
]
