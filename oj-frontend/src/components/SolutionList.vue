<!-- filepath: /C:/Users/乔书祥/Desktop/OJ/oj-frontend/src/components/SolutionList.vue -->
<template>
    <div>
        <h1>Solution List</h1>
        <ul>
            <li v-for="solution in solutions" :key="solution.id">{{ solution.title }}</li>
        </ul>
        <button @click="prevPage" :disabled="page === 1">Previous</button>
        <button @click="nextPage">Next</button>
    </div>
</template>

<script>
import { getSolutions } from '../api/Solution';

export default {
    data() {
        return {
            solutions: [],
            page: 1,
            pageSize: 10,
        };
    },
    methods: {
        async fetchSolutions() {
            const data = await getSolutions(this.page, this.pageSize);
            this.solutions = data.items;
        },
        prevPage() {
            if (this.page > 1) {
                this.page--;
                this.fetchSolutions();
            }
        },
        nextPage() {
            this.page++;
            this.fetchSolutions();
        },
    },
    created() {
        this.fetchSolutions();
    },
};
</script>

<style scoped>
/* 添加你的样式 */
</style>