<template>
    <div>
        <h1>Contest List</h1>
        <ul>
            <li v-for="contest in contests" :key="contest.id">{{ contest.title }}</li>
        </ul>
        <button @click="prevPage" :disabled="page === 1">Previous</button>
        <button @click="nextPage">Next</button>
    </div>
</template>

<script>
import { getContests } from '../api/Contest';

export default {
    data() {
        return {
            contests: [],
            page: 1,
            pageSize: 10,
        };
    },
    methods: {
        async fetchContests() {
            const data = await getContests(this.page, this.pageSize);
            this.contests = data.items;
        },
        prevPage() {
            if (this.page > 1) {
                this.page--;
                this.fetchContests();
            }
        },
        nextPage() {
            this.page++;
            this.fetchContests();
        },
    },
    created() {
        this.fetchContests();
    },
};
</script>

<style scoped>
/* 添加你的样式 */
</style>