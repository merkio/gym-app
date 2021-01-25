import Home from '@/views/Home.vue'
import Programs from '@/views/gym/Programs.vue'
import Exercises from '@/views/gym/Exercises.vue'
import Results from '@/views/gym/Results.vue'
import About from '@/views/About.vue'
import FullViewExercise from '@/components/exercises/FullViewExercise'
import ProgramDetails from "@/components/programs/ProgramDetails";
import Register from "@/views/user/Register";

export default [
  { path: '/', name: 'home', component: Home },
  { path: '/register', name: 'register', component: Register },
  { path: '/programs', name: 'programs', component: Programs },
  { path: '/programs/:id', name: 'programDetails', component: ProgramDetails},
  { path: '/exercises', name: 'exercises', component: Exercises },
  { path: '/exercises/:id', name: 'exercise', component: FullViewExercise},
  { path: '/results', name: 'results', component: Results },
  { path: '/about', name: 'about', component: About }
]
