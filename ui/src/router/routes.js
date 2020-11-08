import Home from '@/views/Home.vue'
import Programs from '@/views/Programs.vue'
import Exercises from '@/views/Exercises.vue'
import Results from '@/views/Results.vue'
import About from '@/views/About.vue'
import ProgramDetails from "@/components/programs/ProgramDetails";

export default [
  { path: '/', name: 'home', component: Home },
  { path: '/programs', name: 'programs', component: Programs },
  { path: '/programs/:id', name: 'program-details', component: ProgramDetails},
  { path: '/exercises', name: 'exercises', component: Exercises },
  { path: '/results', name: 'results', component: Results },
  { path: '/about', name: 'about', component: About }
]
