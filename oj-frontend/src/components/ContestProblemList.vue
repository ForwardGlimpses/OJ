<!-- filepath: /C:/Users/乔书祥/Desktop/OJ/oj-frontend/src/components/ContestProblemList.vue -->
<template>
    <div>
        <h1>Contest Problem List</h1>
        <ul>
            <li v-for="problem in problems" :key="problem.id">{{ problem.title }}</li>
        </ul>
        <button @click="prevPage" :disabled="page === 1">Previous</button>
        <button @click="nextPage">Next</button>
    </div>
</template>

<script>
import { getContestProblems } from '../api/ContestProblem';

export default {
    props: {
        contestId: {
            type: Number,
            required: true,
        },
    },
    data() {
        return {
            problems: [],
            page: 1,
            pageSize: 10,
        };
    },
    methods: {
        async fetchProblems() {
            const data = await getContestProblems(this.contestId, this.page, this.pageSize);
            this.problems = data.items;
        },
        prevPage() {
            if (this.page > 1) {
                this.page--;
                this.fetchProblems();
            }
        },
        nextPage() {
            this.page++;
            this.fetchProblems();
        },
    },
    created() {
        this.fetchProblems();
    },
};
</script>

<style scoped>
/* 添加你的样式 */
</style>