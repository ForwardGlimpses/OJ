import { createRouter, createWebHistory } from 'vue-router';
import ProblemList from '../components/ProblemList.vue';
import SolutionList from '../components/SolutionList.vue';
import SourceCodeList from '../components/SourceCodeList.vue';
import ContestList from '../components/ContestList.vue';
import ContestProblemList from '../components/ContestProblemList.vue';
import UserList from '../components/UserList.vue';

const routes = [
    { path: '/problems', component: ProblemList },
    { path: '/solutions', component: SolutionList },
    { path: '/source-codes', component: SourceCodeList },
    { path: '/contests', component: ContestList },
    { path: '/contest/:contestId/problems', component: ContestProblemList, props: true },
    { path: '/users', component: UserList },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;